package nhaccuatui

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
	"time"
)

type NCTClient struct {
	UserAgent string
	InputURL  string
	MetaURL   string
	MP3URL    string
	FileName  string
	Location  string
}

type NCTMeta struct {
	XMLName xml.Name `xml:"tracklist"`
	Type    string   `xml:"type"`
	Track   Track    `xml:"track"`
}

type Track struct {
	XMLName    xml.Name `xml:"track"`
	Title      string   `xml:"title"`
	Time       string   `xml:"time"`
	Creator    string   `xml:"creator"`
	Location   string   `xml:"location"`
	LocationHQ string   `xml:"locationHQ"`
	HashHQ     string   `xml:"hashHQ"`
	Info       string   `xml:"info"`
	Lyrics     string   `xml:"lyrics"`
	BgImage    string   `xml:"bgimage"`
	Avatar     string   `xml:"avatar"`
	NewTab     string   `xml:"newtab"`
	Kbit       string   `xml:"kbit"`
	Key        string   `xml:"key"`
}

var httpClient = &http.Client{
	Timeout: time.Second * 10,
}

func (c *NCTClient) Do(url string) string {
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("User-Agent", c.UserAgent)
	resp, _ := httpClient.Do(req)
	defer resp.Body.Close()
	payload, _ := ioutil.ReadAll(resp.Body)
	return string(payload)
}

func (c *NCTClient) getMetaURL() {
	page := c.Do(c.InputURL)
	c.MetaURL = extractMetaURL(page)
}

func extractMetaURL(payload string) string {
	pattern := `player.peConfig.xmlURL = "(.*?)";`
	r := regexp.MustCompile(pattern)
	match := r.FindStringSubmatch(payload)
	if match != nil {
		return match[1]
	}
	return ""
}

func (c *NCTClient) getMP3URL() {
	payload := c.Do(c.MetaURL)
	c.MP3URL = extractMP3URL(payload)
	c.FileName = createFileName(c.MP3URL)
}

func extractMP3URL(payload string) string {
	var nctmeta NCTMeta
	xml.Unmarshal([]byte(payload), &nctmeta)
	url := strings.TrimSpace(nctmeta.Track.Location)
	return url
}

func createFileName(url string) string {
	paths := strings.Split(url, "/")
	return paths[len(paths)-1]
}

func (c *NCTClient) DownloadMP3File() string {
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

func NewClient(url string, ua string, location string) (*NCTClient, error) {
	c := &NCTClient{
		InputURL:  url,
		UserAgent: ua,
		Location:  location,
	}
	c.getMetaURL()
	c.getMP3URL()
	return c, nil
}
