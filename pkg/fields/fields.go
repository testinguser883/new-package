package fields

import (
        "log"
        "os/exec"
)

func GetOutput(resource string, pattern string) string {
        out, err := exec.Command("echo", "awesome!!", resource).Output()
        if err != nil {
                log.Fatal(err)
        }
        return string(out)
}
