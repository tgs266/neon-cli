package neon

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/tgs266/neon-cli/api"
)

func CreateProduct(host string, manifest api.ProductManifest) {
	manifestJSON, err := json.Marshal(manifest)
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	_, err = http.Post(fmt.Sprintf("%s/api/v1/products", host), "application/json", bytes.NewBuffer(manifestJSON))
	if err != nil {
		log.Println("ERROR: ", err)
		return
	}
	fmt.Printf("Successfully created product %s\n", manifest.Name)
}
