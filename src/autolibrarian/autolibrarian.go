

package autolibrarian

import (
    "fmt"
    "os"
    "puppetfile"
    "librarian"
)

func CLI(module string, ref string) {

    // Read in the configuration
    var c conf
    c.LoadConfig()
    if c == (conf{}) {
        fmt.Fprintf(os.Stderr, "No config file found\n")
        return
    }
    fmt.Fprintf(os.Stdout, "conf : %v\n", c)

    // What environments do we have?
    envs := enumerateEnvironments(c.PuppetEnvironmentPath)

    // Step through environments one-by-one
    for _, env := range envs {
        fmt.Fprintf(os.Stdout, "Checking env %v\n", env)
        puppet_file := c.PuppetEnvironmentPath + "/" + env + "/Puppetfile"
        modules, err := puppetfile.Read(puppet_file)
        if err != nil {
            fmt.Fprintf(os.Stderr, "WARNING could not open Puppetfile in %v environment\n", env)
        }

        // Is this module in this Puppetfile
        for _, m := range modules {
            module_name := m.GetName()
            if module_name == module {
                fmt.Fprintf(os.Stdout, "Updating module %s\n", module_name)
                librarian.Update(module_name)
            }
        }

    }
    //fmt.Fprintf(os.Stdout, contents)
}
