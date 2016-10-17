package player

import (
	"bufio"
	"os"
	"os/exec"
	"strings"
)

func mocp(param ...string) (string, error) {
	out, err := exec.Command("mocp", param...).Output()
	return string(out), err
}

func Play() (string, error) {
	return mocp("-p")
}

func Stop() (string, error) {
	return mocp("-s")
}

func TogglePause() (string, error) {
	return mocp("-G")
}

func Pause() (string, error) {
	return mocp("-P")
}

func Unpause() (string, error) {
	return mocp("-P")
}

func Next() (string, error) {
	return mocp("-f")
}

func Prev() (string, error) {
	return mocp("-r")
}

func Clear() (string, error) {
	return mocp("-c")
}

func Append(file string) (string, error) {
	return mocp("-a", file)
}

func Info() (string, error) {
	return mocp("-i")
}

func ShowPlaylist(path string) ([]map[string]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	pl := []map[string]string{}
	for scanner.Scan() {
		if strings.Contains(scanner.Text(), "#EXTINF") {
			s := strings.Split(scanner.Text(), "#EXTINF:")
			info := strings.Split(s[1], ",")
			scanner.Scan()
			mp3File := scanner.Text()
			m := make(map[string]string)
			m["length"] = info[0]
			m["title"] = info[1]
			m["path"] = mp3File
			pl = append(pl, m)
		}
	}
	return pl, nil

}

func Volume(command string, param ...string) (string, error) {
	out, err := exec.Command(command, param...).Output()
	return string(out), err
}
