package main

import (
	"encoding/json"
	"io"
	"log"
	"log/slog"
	"net/http"
)

func main() {
	err := http.ListenAndServe(":8080", http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		input := struct {
			URL string `json:"url,omitempty"`
		}{}
		if err := json.NewDecoder(req.Body).Decode(&input); err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}

		slog.Info("Downloading", "url", input.URL)
		resp, err := http.Get(input.URL)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		respData, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(rw, err.Error(), http.StatusInternalServerError)
			return
		}
		slog.Info("Response", "code", input.URL, "length", len(respData), "data", string(respData))
		_, _ = rw.Write(respData)
	}))
	log.Fatal(err)
}
