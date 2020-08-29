package github

import (
	"bufio"
	"context"
	"encoding/json"
	"github.com/despreston/vimlytics/internal/api"
	"github.com/despreston/vimlytics/internal/cache"
	"github.com/despreston/vimlytics/internal/vimparser"
	db "github.com/despreston/vimlytics/mongo"
	"go.mongodb.org/mongo-driver/bson"
	mongoOpts "go.mongodb.org/mongo-driver/mongo/options"
	"io/ioutil"
	"log"
	"net/http"
)

// prefix for the etag in redis
const cachePrefix = "etag-"

var httpClient = &http.Client{}

type GithubUser struct {
	Login            string `json:"login"`
	Id               int    `json:"id"`
	Avatar_url       string `json:"avatar_url"`
	Html_url         string `json:"html_url"`
	Name             string `json:"name"`
	Company          string `json:"company"`
	Location         string `json:"location"`
	Twitter_username string `json:"twitter_username"`
}

// Get the GH user by "login" name.
func userFromDb(login string) (GithubUser, error) {
	filter := bson.D{{Key: "login", Value: login}}
	coll := db.Db().Collection("users")
	var result GithubUser

	err := coll.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		return GithubUser{}, err
	}

	return result, nil
}

// Save the user into the DB.
func saveUser(u GithubUser) (GithubUser, error) {
	var result GithubUser
	filter := bson.M{"login": u.Login}
	coll := db.Db().Collection("users")
	upsert := true
	opts := mongoOpts.FindOneAndReplaceOptions{Upsert: &upsert}
	ctx := context.Background()
	err := coll.FindOneAndReplace(ctx, filter, u, &opts).Decode(&result)

	return result, err
}

func userFromGh(login string) (*http.Response, error) {
	url := "https://api.github.com/users/" + login

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return &http.Response{}, err
	}

	req.Header.Add("Accept", "application/vnd.github.v3+json")

	// Get etag from redis, add "If-None-Match" header if etag exists.
	etag, found := cache.Get(cachePrefix + login)
	if found {
		log.Printf("Found etag for user %s: %s", login, etag)
		req.Header.Add("If-None-Match", etag)
	}

	resp, err := httpClient.Do(req)
	if err != nil {
		log.Printf("Error fetching from github: %+v", err)
		return &http.Response{}, err
	}

	return resp, nil
}

func User(w http.ResponseWriter, r *http.Request) *api.Error {
	var u GithubUser
	q := r.URL.Query()
	login := q.Get("login")

	ghResp, err := userFromGh(login)

	// User data hasn't changed. Send saved copy.
	if ghResp.StatusCode == 304 {
		log.Printf("GH returned 304 for user %s", login)
		u, err = userFromDb(login)

		if err != nil {
			return &api.Error{
				Error:   err,
				Message: "",
				Code:    500,
			}
		}
	} else if err != nil || ghResp.StatusCode != 200 {
		return &api.Error{
			Error:   err,
			Message: "Error getting user from Github.",
			Code:    500,
		}
	} else {
		defer ghResp.Body.Close()

		data, err := ioutil.ReadAll(ghResp.Body)
		err = json.Unmarshal(data, &u)

		if err != nil {
			return &api.Error{Error: err, Message: "Unknown error", Code: 500}
		}

		// Save etag
		cache.Set(cachePrefix+login, ghResp.Header.Get("ETag"))

		u, err = saveUser(u)
	}

	response, err := json.Marshal(u)
	if err != nil {
		return &api.Error{Error: err, Message: "Unknown error", Code: 500}
	}

	w.Write(response)
	return nil
}

func Vimrc(w http.ResponseWriter, r *http.Request) *api.Error {
	q := r.URL.Query()
	url := "https://raw.githubusercontent.com/"
	url += q.Get("login") + "/" + q.Get("repo") + "/master/.vimrc"

	resp, err := http.Get(url)
	if err != nil {
		return &api.Error{Error: err, Message: "Bad request", Code: 401}
	}

	defer resp.Body.Close()

	buf := bufio.NewReader(resp.Body)
	settings := vimparser.Parse(buf)

	json, err := json.Marshal(settings)
	if err != nil {
		return &api.Error{
			Error:   err,
			Message: "Error fetching vimrc from Github.",
			Code:    500,
		}
	}

	w.Write([]byte(json))
	return nil
}
