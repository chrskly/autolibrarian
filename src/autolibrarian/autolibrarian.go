

package autolibrarian

import (
    "fmt"
    "os"
    "puppetfile"
    "librarian"
)

func CLI(module string, ref string) {

    var c conf

    // Read in the configuration
    c.LoadConfig()
    if c == (conf{}) {
        fmt.Fprintf(os.Stderr, "No config file found\n")
        return
    }
    fmt.Fprintf(os.Stdout, "conf : %v\n", c)

    // What environments do we have?
    enumerateEnvironments(c.PuppetEnvironmentPath)

    modules, err := puppetfile.Read("Puppetfile.short")
    //contents, err := puppetfile.Read("Puppetfile")
    if err != nil {
        fmt.Fprintf(os.Stderr, "ERROR command failed")
        return
    }

    // Is this module in this Puppetfile
    for _, m := range modules {
        module_name := m.GetName()
        fmt.Printf(">> %v\n", module_name)
        m.ToString()
        if module_name == module {
            fmt.Sprintf("Updating module %s\n", module_name)
            librarian.Update(module_name)
        }
    }
    //fmt.Fprintf(os.Stdout, contents)
}
