package fields

// Required packages
import (
        "log"
        "os/exec"
)

// Getoutput wrapper function returns executed commands output
func GetOutput(resource string, pattern string) string {
        out, err := exec.Command("echo", "awesome!!", resource).Output()
        if err != nil {
                log.Fatal(err)
        }
        return string(out)
}
