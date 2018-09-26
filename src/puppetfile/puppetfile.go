
package puppetfile

import (
    "fmt"
    "os"
    "io/ioutil"
    "regexp"
    //"errors"
)

// A single puppet module found in a Puppetfile

type PuppetModule struct {
    Name string
    Ref  string
    Git  string
}

func (p *PuppetModule) SetName(name string) {
    p.Name = name
}

func (p *PuppetModule) GetName() string {
    return p.Name
}

func (p *PuppetModule) SetRef(ref string) {
    p.Ref = ref
}

func (p *PuppetModule) SetGit(git string) {
    p.Git = git
}

func (p *PuppetModule) ToString() {
    fmt.Fprintf(os.Stdout, "name:%s, ref:%s, git:%s\n", p.Name, p.Ref, p.Git)
}


func Read(path string) ([]PuppetModule, error) {
    // Read in the contents of the Puppetfile given it's path
    contents, err := ioutil.ReadFile(path)
    if err != nil {
        return []PuppetModule{}, err
    }

    // Our regex
    re := regexp.MustCompile(
        "(?m)^mod [\"'](?P<name>[a-zA-Z0-9\\/\\-\\_]*)[\"']" +
        "(," +
        "(\n)?(\\s+)?(\n)?" +
        "(:git => \"(?P<git>[a-zA-Z0-9:\\/\\.\\-@_]*)?\")?" +
        ")?" +
        "(," +
        "(\n)?(\\s+)?(\n)?" +
        "(:ref => [\"'](?P<ref>[a-zA-Z0-9\\.\\-]*)[\"'])" +
        ")?" +
        "(.*?)")
    // names of named matches
    match_names := re.SubexpNames()

    // results
    matches := re.FindAllSubmatch(contents,-1)
    // Map of results
    module_matches := []PuppetModule{}

    for _, match := range matches {

        // New module
        mod := PuppetModule{}

        for i, n := range match {
            n_string := string(n)
            if match_names[i] == "name" {
                mod.SetName(n_string)
            }
            if match_names[i] == "git" {
                mod.SetGit(n_string)
            }
            if match_names[i] == "ref" {
                mod.SetRef(n_string)
            }
        }

        //
        //mod.ToString()
        module_matches = append(module_matches, mod)

    }

    //return module_matches, errors.New("failed to parse")
    return module_matches, nil

}
