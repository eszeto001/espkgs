package sysos
// System and OS utilities.

import (
  "fmt"
  "os/exec"
  "log"
)

func System(cmd string) {
   fmt.Printf("+ %s\n", cmd)
   out, err := exec.Command("/bin/sh", "-c", cmd).Output()
   if err != nil {
      log.Fatal(err)
   }
   fmt.Printf("%s\n", out)
}

