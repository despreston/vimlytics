package vimrc

import (
	"bufio"
	"encoding/json"
	"github.com/despreston/vimlytics/pkg/vimparser"
	"io"
	"log"
	"net/http"
)

func Post(w http.ResponseWriter, r *http.Request) {
	// 5mb
	var maxFileSize int64 = 5 << 20

	r.Body = http.MaxBytesReader(w, r.Body, maxFileSize)

	// Parse the form as a stream
	reader, err := r.MultipartReader()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	part, err := reader.NextPart()
	if err != nil && err != io.EOF {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// parse file field
	if part.FormName() != "vimrc" {
		http.Error(w, "vimrc is expected", http.StatusBadRequest)
		return
	}

	buf := bufio.NewReader(part)
	sniff, _ := buf.Peek(512)
	contentType := http.DetectContentType(sniff)

	if contentType != "text/plain; charset=utf-8" {
		http.Error(w, "file type not allowed", http.StatusBadRequest)
		return
	}

	settings := vimparser.Parse(buf)

	json, err := json.Marshal(settings)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Printf("Error marshalling json: %v", err)
	}

	w.Write([]byte(json))
}
