
package autolibrarian

import (
  "os"
  "errors"
  "io/ioutil"
  "gopkg.in/yaml.v2"
)


type conf struct {
    PuppetEnvironmentPath string `yaml:"puppet_environment_path"`
}

// Slice containing paths at which to look for the config file
var config_paths = []string{"/etc/autolibrarian.conf",
                            "/usr/local/etc/autolibrarian.conf"}

func findConfigFile() (string, error) {
    // Locate the config file
    for _, p := range config_paths {
        if _, err := os.Stat(p); !os.IsNotExist(err) {
            return p, nil
        }
    }
    return "", errors.New("No config file found")
}

func (c *conf) LoadConfig() *conf {
    // Where is there a config file?
    config_path, err := findConfigFile()
    if err != nil {
        return nil
    }
    // Read in the config file
    yamlFile, err := ioutil.ReadFile(config_path)
    if err != nil {
        return nil
    }
    // Parse the yaml
    err = yaml.Unmarshal(yamlFile, c)
    if err != nil {
        return  nil
    }
    return c
}
