package api

type AppManifest struct {
	Kind     string   `yaml:"kind"`
	Name     string   `yaml:"name"`
	Products []string `yaml:"products"`
}
