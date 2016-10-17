package main

import (
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	files := http.FileServer(http.Dir("static"))
	mux.Handle("/static/", http.StripPrefix("/static/", files))
	mux.HandleFunc("/", logHandler(index))
	mux.HandleFunc("/admin", logHandler(admin))
	mux.HandleFunc("/info", logHandler(info))
	mux.HandleFunc("/playlist", logHandler(playlist))
	mux.HandleFunc("/play", logHandler(play))
	mux.HandleFunc("/stop", logHandler(stop))
	mux.HandleFunc("/togglePause", logHandler(togglePause))
	mux.HandleFunc("/pause", logHandler(pause))
	mux.HandleFunc("/unpause", logHandler(unpause))
	mux.HandleFunc("/next", logHandler(next))
	mux.HandleFunc("/prev", logHandler(prev))
	mux.HandleFunc("/clear", logHandler(clear))
	mux.HandleFunc("/volume_up", logHandler(volumeUp))
	mux.HandleFunc("/volume_down", logHandler(volumeDown))
	server := &http.Server{
		Addr:    config.Address,
		Handler: mux,
	}
	server.ListenAndServe()
}
