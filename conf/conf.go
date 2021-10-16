package conf

import (
	"github.com/BurntSushi/toml"
	"log"
)

var (
	GlocalConfig ConfigStruct
)

func ParseConfig(pathConfFile string){
	if _, err := toml.DecodeFile(pathConfFile, &GlocalConfig); err != nil {
		log.Fatal(err)
	}
}
type ConfigStruct struct {
	Default struct {
		Path string `toml:"path"`
	} `toml:"default"`
	Special struct {
		Monitors []string `toml:"monitors"`
		Path    string   `toml:"path"`
	} `toml:"special"`
}
