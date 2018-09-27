

package autolibrarian

import (
    "fmt"
    "os"
    "puppetfile"
    "librarian"
)

func CLI(targetModule string, targetRef string) {

    // Read in the configuration
    var c conf
    c.LoadConfig()
    if c == (conf{}) {
        fmt.Fprintf(os.Stderr, "No config file found\n")
        return
    }
    fmt.Fprintf(os.Stdout, "conf : %v\n", c)

    // What environments do we have?
    environments := enumerateEnvironments(c.PuppetEnvironmentPath)

    // Step through environments one-by-one
    for _, environment := range environments {
        fmt.Fprintf(os.Stdout, "Checking env %v\n", environment)
        puppetFilePath := c.PuppetEnvironmentPath + "/" + environment + "/Puppetfile"
        environmentModules, err := puppetfile.Read(puppetFilePath)
        if err != nil {
            fmt.Fprintf(os.Stderr, "WARNING could not open Puppetfile in %v environment\n", environment)
        }

        // Is this module in this Puppetfile
        for _, environmentModule := range environmentModules {
            environmentModuleName := environmentModule.GetName()
            if environmentModuleName == targetModule {
                fmt.Fprintf(os.Stdout, "Updating module %s\n", targetModule)
                librarian.Update(targetModule, c.PuppetEnvironmentPath + "/" + env)
            }
        }

    }
    //fmt.Fprintf(os.Stdout, contents)
}
