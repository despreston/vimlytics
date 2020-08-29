package vimrc

import (
	"bufio"
	"encoding/json"
	"github.com/despreston/vimlytics/internal/api"
	"github.com/despreston/vimlytics/internal/vimparser"
	"io"
	"net/http"
)

func Upload(w http.ResponseWriter, r *http.Request) *api.Error {
	// 5mb
	var maxFileSize int64 = 5 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)

	// Parse the form as a stream
	reader, err := r.MultipartReader()
	if err != nil {
		return &api.Error{Error: err, Message: "", Code: 400}
	}

	part, err := reader.NextPart()
	if err != nil && err != io.EOF {
		return &api.Error{Error: err, Message: "", Code: 400}
	}

	// parse file field
	if part.FormName() != "vimrc" {
		return &api.Error{
			Error:   err,
			Message: "No vimrc field provided",
			Code:    400,
		}
	}

	buf := bufio.NewReader(part)
	sniff, _ := buf.Peek(512)
	contentType := http.DetectContentType(sniff)

	if contentType != "text/plain; charset=utf-8" {
		return &api.Error{
			Error:   err,
			Message: "File type not allowed",
			Code:    400,
		}
	}

	settings := vimparser.Parse(buf)

	json, err := json.Marshal(settings)
	if err != nil {
		return &api.Error{Error: err, Message: "Unknown error", Code: 500}
	}

	w.Write([]byte(json))
	return nil
}
