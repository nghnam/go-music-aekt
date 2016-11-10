package zing

import (
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type ZingClient struct {
	UserAgent string
	InputURL  string
	MetaURL   string
	MP3URL    string
	FileName  string
	Location  string
}

var httpClient = &http.Client{
	Timeout: time.Second * 60,
}

func (c *ZingClient) Do(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", c.UserAgent)
	resp, _ := httpClient.Do(req)
	defer resp.Body.Close()
	payload, _ := ioutil.ReadAll(resp.Body)
	return string(payload)
}

func (c *ZingClient) getMetaURL() {
	page := c.Do(c.InputURL)
	c.MetaURL = extractMetaURL(page)
}

func extractMetaURL(payload string) string {
	pattern := `var xml_link = "(.*?)";`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(payload)
	if match != nil {
		return match[1]
	}
	return ""
}

func (c *ZingClient) getMP3URL() {
	payload := c.Do(c.MetaURL)
	c.MP3URL = extractMP3URL(payload)
	c.FileName = createFileName(payload)
}

func extractMP3URL(payload string) string {
	pattern := `"(s1\.mp3\.zdn.*?)"`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(payload)
	if match != nil {
		return "http://" + match[1]
	}
	return ""
}

func createFileName(payload string) string {
	pattern := `"link":"(.*?)"`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(payload)
	if match != nil {
		s := strings.Split(match[1], "/")
		return s[2] + ".mp3"
	}
	return ""
}

func (c *ZingClient) DownloadMP3File() string {
	file := c.Location + "/" + c.FileName
	out, _ := os.Create(file)
	defer out.Close()
	req, _ := http.NewRequest("GET", c.MP3URL, nil)
	req.Header.Add("User-Agent", c.UserAgent)
	resp, _ := httpClient.Do(req)
	defer resp.Body.Close()
	io.Copy(out, resp.Body)
	return file
}

func NewClient(url string, ua string, location string) (*ZingClient, error) {
	c := &ZingClient{
		InputURL:  url,
		UserAgent: ua,
		Location:  location,
	}
	c.getMetaURL()
	c.getMP3URL()
	return c, nil
}
