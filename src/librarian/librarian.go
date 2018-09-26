
// Wrapper around librarian-puppet

package librarian

import (
    "errors"
    "os/exec"
)

func Update(module string) error {
    cmd := exec.Command("librarian-puppet", "update", module)
    err := cmd.Run()
    if err != nil {
        return errors.New("Failed to update module " + module)
    }
    return nil
}
