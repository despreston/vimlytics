package plugged

import (
	"encoding/json"
	"fmt"
	"github.com/despreston/vimlytics/pkg/cache"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

type Plugin struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func fetch(name string, wg *sync.WaitGroup, ch chan Plugin) {
	defer wg.Done()

	p := Plugin{Name: name}

	// Check the cache
	if description, has := cache.Get(name); has {
		p.Description = description
		ch <- p
		return
	}

	url := fmt.Sprintf("https://api.github.com/repos/%s", name)
	resp, err := http.Get(url)

	log.Printf("Fetched %s\n", url)

	if err != nil {
		log.Println(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.Printf("Error receiving response from GH for %v: %v\n", name, err)
		return
	}

	var parsed struct {
		Description string `json:"description"`
	}

	if err = json.Unmarshal(body, &parsed); err != nil {
		log.Printf("Error parsing resp from GH for %v: %v\n", name, err)
		return
	}

	log.Println(parsed)

	p.Description = parsed.Description

	// save to cache
	go cache.Set(p.Name, p.Description)

	ch <- p
}

func FetchAll(names []string) []Plugin {
	var plugins []Plugin
	var wg sync.WaitGroup

	pluginCh := make(chan Plugin)

	for _, p := range names {
		wg.Add(1)
		go fetch(p, &wg, pluginCh)
		plugins = append(plugins, <-pluginCh)
	}

	wg.Wait()
	return plugins
}
