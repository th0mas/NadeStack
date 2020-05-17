package csgo

import (
	"encoding/json"
	"errors"
	"time"

	"github.com/th0mas/NadeStack/models"
)

const loadCfgCmd = "get5_loadmatch config.json"

// BuildCSGOServer builds a new server on dathost
func BuildCSGOServer(g *models.Game, GSLT string, statusCallback func(*models.Game)) error {
	// Create the CSGO Server
	id, _ := CreateCSGOServer(g.Match.PlayersPerTeam*2, g.Match.MapList[0], g.Match.ID, GSLT)
	g.IncrementGameStatus()
	statusCallback(g)

	// Upload plugin zip to server
	UploadPluginsToServer(id)
	g.IncrementGameStatus()
	statusCallback(g)

	// Generate and upload JSON
	g.Match.GenerateTeamIDS()
	conf, _ := json.Marshal(g.Match)
	UploadJSONToServer(id, "config.json", conf, "file")
	g.IncrementGameStatus()
	statusCallback(g)

	// Unzip plugin files on server
	UnzipPlugins(id)
	g.IncrementGameStatus()
	statusCallback(g)

	// Start Server
	StartServer(id)
	g.IncrementGameStatus()
	statusCallback(g)

	*g.ServerIP, _ = waitForServerIP(id)
	g.IncrementGameStatus()
	statusCallback(g)

	SendCommandToServer(id, loadCfgCmd)
	g.IncrementGameStatus()
	statusCallback(g)

	return nil
}

func waitForServerIP(id string) (string, error) {
	ch := make(chan string, 1)
	defer close(ch)

	go func() {
		for {
			if status, ip := GetServerStatus(id); !status {
				ch <- ip
				return
			}
			time.Sleep(2 * time.Second)
		}
	}()

	timer := time.NewTimer(2 * time.Minute)
	defer timer.Stop()

	select {
	case ip := <-ch:
		return ip, nil
	case <-timer.C:
		return "", errors.New("timed out waiting for serve to load")
	}

}
