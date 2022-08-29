package neon

import (
	"github.com/mitchellh/mapstructure"
	"github.com/tgs266/neon-cli/api"
)

func Apply(host string, manifests []map[string]any) {
	for _, manifest := range manifests {
		if val, exists := manifest["kind"]; exists {
			switch val {
			case "product":
				man := api.ProductManifest{}
				mapstructure.Decode(manifest, &man)
				CreateProduct(host, man)
			case "release":
				man := api.ReleaseManifest{}
				mapstructure.Decode(manifest, &man)
				CreateRelease(host, man)
			}
		}
	}

}
