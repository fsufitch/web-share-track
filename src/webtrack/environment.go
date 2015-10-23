package webtrack

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"
)

var resource_dir string

func SetResourceDir(newdir string) {
	resource_dir = newdir
}

func GetResource(path ...string) string {
	full_path := append([]string{resource_dir}, path...)
	return filepath.Join(full_path...)
}


/////


type ConfigurationHolder struct {
	ListenHostname string
	ListenPort int
	CallbackHost string
}
var Configuration ConfigurationHolder

func BuildConfiguration(confpath string) {
	data, err := ioutil.ReadFile(confpath)
	CheckErrorFatal(err)
	
	CheckErrorFatal(json.Unmarshal(data, &Configuration))
}
