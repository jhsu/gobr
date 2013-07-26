package gobr

import(
  "bytes"
  "fmt"
  "os/exec"
)

// Get a range of local branches.
func Branches() []string {
  fmt.Println("HELLO")
  cmd := exec.Command("git branch")
  var out bytes.Buffer
  err := cmd.Run()
  if err != nil {
  }
  fmt.Printf("in all caps: %q\n", out.String())
  // return []string{"master", "production"}
  return []string{}
}
