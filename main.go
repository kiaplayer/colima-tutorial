package main

import (
	"fmt"
	"log/slog"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request) {
		slog.Info("Received request to /hello endpoint")

		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Hello World!")
	})

	slog.Info("Starting server on port 8890")

	err := http.ListenAndServe(":8890", nil)
	if err != nil {
		slog.Error("Application finished with an error", "error", err)
	}
}
