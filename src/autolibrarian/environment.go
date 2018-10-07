
package autolibrarian

import (
    "io/ioutil"
    "log"
)

func enumerateEnvironments(environment_path string) []string {
    // Get all of the puppet environments under environment_path

    files, err := ioutil.ReadDir(environment_path)
    if err != nil {
        log.Fatal(err)
    }

    // We only care about directories
    dirs := []string{}
    for _, f := range files {
        if f.IsDir() {
            dirs = append(dirs, f.Name())
        }
    }

    return dirs

}
