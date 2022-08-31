package neon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/tgs266/neon-cli/api"
)

func CreateApp(host string, manifest api.AppManifest) {
	manifestJSON, err := json.Marshal(manifest)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	resp, err := http.Post(fmt.Sprintf("%s/api/v1/apps", host), "application/json", bytes.NewBuffer(manifestJSON))
	if err != nil || resp.StatusCode >= 400 {
		log.Printf("ERROR: Failed to create app %s\n", manifest.Name)
		return
	}
	fmt.Printf("Successfully created app %s\n", manifest.Name)
}

func ListApps(host string) {
	resp, err := http.Get(fmt.Sprintf("%s/api/v1/apps", host))
	if err != nil || resp.StatusCode >= 400 {
		log.Println("ERROR: could not get apps")
		return
	}
	str, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(string(str))
}
