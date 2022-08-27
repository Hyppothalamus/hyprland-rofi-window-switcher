package commands

import (
	"log"
	"os/exec"
)

func Command (c string) string {
    out, err := exec.Command("sh", "-c", c).Output()
    
    if err != nil {
        log.Fatal(err)
    }

    return string(out)
}
