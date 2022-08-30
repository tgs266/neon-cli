package api

type ReleaseManifest struct {
	Kind           string       `mapstructure:"kind"`
	ProductName    string       `mapstructure:"product-name"`
	ProductVersion string       `mapstructure:"product-version"`
	ReleaseChannel string       `mapstructure:"release-channel"`
	HelmChart      string       `mapstructure:"helm-chart"`
	Dependencies   []Dependency `mapstructure:"dependencies"`
}

type Dependency struct {
	ProductName string `mapstructure:"product-name"`
	MinVersion  string `mapstructure:"min-version"`
	MaxVersion  string `mapstructure:"max-version"`
}
