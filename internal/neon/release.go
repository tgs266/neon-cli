package neon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tgs266/neon-cli/api"
)

func CreateRelease(host string, manifest api.ReleaseManifest) {
	manifestJSON, err := json.Marshal(manifest)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	_, err = http.Post(fmt.Sprintf("%s/api/v1/releases", host), "application/json", bytes.NewBuffer(manifestJSON))
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Printf("Successfully created release for product %s\n", manifest.ProductName)
}
