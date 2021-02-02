package handler

import (
	"fmt"
	"net/http"
)

func IndexHandler() http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, `
			<h1>MySQL Exporter</h1>
			<a href="/metrics/">Metrics</a>
		`)
	})
}
