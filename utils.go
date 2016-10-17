package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"
)

type Configuration struct {
	Address       string `json:"Address"`
	UserAgent     string `json:"UserAgent"`
	SaveLocation  string `json:"SaveLocation"`
	MocDirectory  string `json:"MocDirectory"`
	MocPlaylist   string `json:"MocPlaylist"`
	VolumeCommand string `json:"VolumeCommand"`
	VolumeUp      string `json:"VolumeUp"`
	VolumeDown    string `json:"VolumeDown"`
}

var config Configuration
var logger *log.Logger

func loadConfig() {
	file, err := os.Open("config.json")
	defer file.Close()
	if err != nil {
		log.Fatalln("Can not open config file")
	}
	decoder := json.NewDecoder(file)
	config = Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		log.Fatalln("Can not get configuration from file")
	}
}

func init() {
	loadConfig()
	logger = log.New(os.Stdout, "", log.Ldate|log.Lmicroseconds)
}

func generateHTML(writer http.ResponseWriter, filenames ...string) {
	var files []string
	for _, file := range filenames {
		files = append(files, fmt.Sprintf("templates/%s.html", file))
	}

	templates := template.Must(template.ParseFiles(files...))
	templates.ExecuteTemplate(writer, "layout", nil)
}

func generateInfoMap(data string) map[string]string {
	m := make(map[string]string)
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		s := strings.Split(scanner.Text(), ": ")
		key, val := s[0], s[1]
		m[key] = val
	}
	return m
}

func logHandler(handler http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		logger.Println(request.RemoteAddr, request.Method, request.URL)
		handler(writer, request)
	}
}
