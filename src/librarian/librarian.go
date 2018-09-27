
// Wrapper around librarian-puppet

package librarian

import (
    "errors"
    "os/exec"
)

func Update(module string, environmentPath string) error {
    cmd := exec.Command("librarian-puppet", "update", module)
    // Set the correct working dir
    cmd.Dir = environmentPath
    err := cmd.Run()
    if err != nil {
        return errors.New("Failed to update module " + module)
    }
    return nil
}
