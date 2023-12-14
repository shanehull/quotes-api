package main

import (
	"bytes"
	"log"
	"math/rand"
	"net/http"
	"time"

	"github.com/syumai/workers"
)

var globalRand = rand.New(rand.NewSource(time.Now().UnixNano()))

func cors(w http.ResponseWriter, req *http.Request) {
	w.Header().Add("Access-Control-Allow-Origin", "*")
	w.Header().Add("Access-Control-Allow-Methods", "GET")
	w.Header().Add("Access-Control-Allow-Headers", "Content-Type, hx-current-url, hx-request, hx-trigger, hx-target")
	w.Header().Add("Content-Type", "text/html; charset=utf-8")
}

func methods(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		// we allow - how good!
	case http.MethodOptions:
		w.WriteHeader(http.StatusNoContent)
		return
	default:
		http.Error(w, "method not allowed", http.StatusMethodNotAllowed)
		return
	}
}

func genRandIndex() int {
	return globalRand.Intn(len(quotes))
}

func quoteHandler(w http.ResponseWriter, req *http.Request) {
	cors(w, req)
	methods(w, req)

	idx := genRandIndex()

	component := quote(quotes[idx].Text, quotes[idx].Author)

	buf := new(bytes.Buffer)
	defer buf.Reset()

	if err := component.Render(req.Context(), buf); err != nil {
		log.Printf("failed to render component:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	_, err := w.Write(buf.Bytes())
	if err != nil {
		log.Printf("failed to write response:", err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}
}

func main() {
	http.HandleFunc("/quote", quoteHandler)

	workers.Serve(nil)
}
