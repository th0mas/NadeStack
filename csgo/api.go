package csgo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/th0mas/NadeStack/config"
	"github.com/th0mas/NadeStack/models"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

const (
	baseUrl     = "https://dathost.net/api/0.1/"
	pluginsFile = "plugins.zip"
)

var (
	user, pass string
)

type server struct {
	Id      string
	Booting bool
	Ip      string
	Ports   ports
}

type ports struct {
	Game string
}

type API struct{}

func (a *API) Run(c *config.Config, _ *models.Models) {
	user = c.DatHostUser
	pass = c.DatHostPass
	log.Println("user: " + c.DatHostUser)
	log.Println("pass: " + c.DatHostPass)
}

func (a *API) Close() {}

func CreateCSGOServer(players int, gameMap string, name string, gslt string) (id string, ip string) {
	v := url.Values{}
	v.Set("game", "csgo")
	v.Set("name", name)
	v.Set("location", "bristol")
	v.Set("csgo_settings.disable_bots", "true")
	v.Set("csgo_settings.enable_sourcemod", "true")
	v.Set("csgo_settings.game_mode", "classic_competitive")
	v.Set("csgo_settings.mapgroup_start_map", gameMap)
	v.Set("csgo_settings.maps_source", "mapgroup")
	v.Set("csgo_settings.password", "nadestack")
	v.Set("csgo_settings.rcon", "pls_never_guess_haha_oh_no")
	v.Set("csgo_settings.slots", strconv.Itoa(players))
	v.Set("csgo_settings.steam_game_server_login_token", gslt)
	v.Set("csgo_settings.tickrate", "128")
	resp, err := api("POST", "game-servers", v)

	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	s := &server{}
	err = json.Unmarshal(body, s)

	if err != nil {
		log.Println(err)
	}

	return s.Id, s.Ip
}

func UploadPluginsToServer(id string) {
	plugins, err := os.Open("./plugins.zip")

	if err != nil {
		panic(err)
	}

	uploadFileToServer(id, "plugins.zip", plugins, "file")
}

func UnzipPlugins(id string) error {
	method := fmt.Sprintf("game-servers/%s/unzip/%s", id, pluginsFile)
	resp, err := api("POST", method, url.Values{})
	if err != nil {
		log.Println(err)
	}
	resp.Body.Close()
	return err
}

func GetServerStatus(id string) (booting bool, ip string) {
	resp, err := api("GET", "game-servers/"+id, url.Values{})
	if err != nil {
		log.Print(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	s := &server{}
	err = json.Unmarshal(body, s)

	return s.Booting, s.Ip + s.Ports.Game

}

func StartServer(id string) error {
	r, err := api("POST", fmt.Sprintf("game-servers/%s/start", id), url.Values{})

	if err != nil {
		log.Println(err)
	}

	r.Body.Close()
	return err
}

func SendCommandToServer(id, command string) error {
	resp, err := api("POST", fmt.Sprintf("game-servers/%s/console", id), url.Values{
		"line": []string{command},
	})

	if err != nil {
		log.Println(err)
	}

	resp.Body.Close()
	return err
}

func api(method, path string, form url.Values) (*http.Response, error) {
	AbsURL := baseUrl + path
	req, err := http.NewRequest(method, AbsURL, strings.NewReader(form.Encode()))

	if err != nil {
		log.Println(err)
	}

	req.SetBasicAuth(user, pass)

	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Add("Accept", "application/json")

	return http.DefaultClient.Do(req)

}

func uploadFileToServer(id, path string, file *os.File, field string) {
	client := http.DefaultClient
	absUrl := fmt.Sprintf("%sgame-servers/%s/files/%s", baseUrl, id, path)

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(field, filepath.Base(file.Name()))

	if err != nil {
		panic(err)
	}

	_, err = io.Copy(part, file)

	if err != nil {
		panic(err)
	}
	writer.Close()
	file.Close()

	req, _ := http.NewRequest("POST", absUrl, body)
	req.Header.Set("Content-Type", writer.FormDataContentType())
	req.SetBasicAuth(user, pass)

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

}
