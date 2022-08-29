package parser

import (
	"io/ioutil"
	"log"
	"strings"

	"gopkg.in/yaml.v2"
)

func ReadManifests(path string) []map[string]any {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}
	strFile := string(file)
	splits := strings.Split(strFile, "---")
	out := []map[string]any{}
	for _, str := range splits {
		b := []byte(str)
		data := map[string]any{}
		err2 := yaml.Unmarshal(b, &data)
		if err2 != nil {
			log.Fatal(err2)
		}
		out = append(out, data)
	}
	return out
}
