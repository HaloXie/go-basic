package main

import (
	"encoding/json"
	"io"
	"net/http"

	"../core"
)

var blockChain *core.Blockchain

func run() {
	http.HandleFunc("/blockchian/get", blockchianGetHandler)
	http.HandleFunc("/blockchian/post", blockchianPostHandler)

	// start
	http.ListenAndServe("localhost:8888", nil)
}

func blockchianGetHandler(w http.ResponseWriter, r *http.Request) {
	bytes, error := json.Marshal(blockChain)
	if error != nil {
		http.Error(w, error.Error(), http.StatusInternalServerError)
		return
	}

	io.WriteString(w, string(bytes))
}

func blockchianPostHandler(w http.ResponseWriter, r *http.Request) {
	blockData := r.URL.Query().Get("data") // demo 写死的
	blockChain.SendData(blockData)

	blockchianGetHandler(w, r)
}

func main() {
	blockChain = core.NewBlockchain()
	run()
}
