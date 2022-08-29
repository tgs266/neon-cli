package api

type ProductManifest struct {
	Kind string `mapstructure:"kind"`
	Name string `mapstructure:"name"`
}
