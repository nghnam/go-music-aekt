package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/nghnam/go-music-aekt/nhaccuatui"
	"github.com/nghnam/go-music-aekt/player"
	"github.com/nghnam/go-music-aekt/zing"
)

func index(writer http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "GET":
		generateHTML(writer, "base", "index", "script", "title")
	case "POST":
		request.ParseForm()
		url := request.Form["music_link"][0]
		var file string
		if strings.Contains(url, "mp3.zing.vn") {
			dl, _ := zing.NewClient(url, config.UserAgent, config.SaveLocation)
			file = dl.DownloadMP3File()
		} else if strings.Contains(url, "nhaccuatui") {
			dl, _ := nhaccuatui.NewClient(url, config.UserAgent, config.SaveLocation)
			file = dl.DownloadMP3File()
		}
		player.Append(file)
	}
}

func admin(writer http.ResponseWriter, request *http.Request) {
	generateHTML(writer, "base", "admin", "script", "title")
}

func info(writer http.ResponseWriter, request *http.Request) {
	out, err := player.Info()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
	}
	m := generateInfoMap(out)
	payload, _ := json.Marshal(m)
	writer.Header().Set("Content-Type", "application/json")
	writer.Write(payload)
}

func play(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Play()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func stop(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Stop()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func togglePause(writer http.ResponseWriter, request *http.Request) {
	_, err := player.TogglePause()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func pause(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Pause()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func unpause(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Unpause()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func next(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Next()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func prev(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Prev()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func clear(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Clear()
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func playlist(writer http.ResponseWriter, request *http.Request) {
	playlist := os.Getenv("HOME") + "/" + config.MocDirectory + "/" + config.MocPlaylist
	pl, err := player.ShowPlaylist(playlist)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		m := make(map[string]interface{})
		m["playlist"] = pl
		payload, _ := json.Marshal(m)
		writer.Header().Set("Content-Type", "application/json")
		writer.Write(payload)
	}
}

func volumeDown(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Volume(config.VolumeCommand, config.VolumeDown)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}

func volumeUp(writer http.ResponseWriter, request *http.Request) {
	_, err := player.Volume(config.VolumeCommand, config.VolumeUp)
	if err != nil {
		writer.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(writer, "Failed")
	} else {
		fmt.Fprintf(writer, "OK")
	}
}
