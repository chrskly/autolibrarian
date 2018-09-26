
package main

import (
    "os"
    "fmt"
    "flag"
    "autolibrarian"
)


func main() {
    // Parse command line args
    modulePtr := flag.String("module", "", "Name of module to check")
    refPtr := flag.String("ref", "", "Branch/tag of module to check")
    flag.Parse()
    if *modulePtr == "" || *refPtr == ""  {
        fmt.Fprintf(os.Stdout, "ERROR : You must specify both <module> and <ref>\n")
        return
    }

    autolibrarian.CLI(*modulePtr, *refPtr)
}
