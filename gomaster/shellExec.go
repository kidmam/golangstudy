package main

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
)

func main() {
	/* cmd := exec.Command("./devnettool.sh")
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	fmt.Printf("combined out:\n%s\n", string(out)) */

	/*cmdStr := "devnettool.sh"
	cmd := exec.Command("sh", "-c", cmdStr)
	_, err := cmd.Output()

	if err != nil {
		println(err.Error())
		return
	}*/

	cmdStr := "devnettool.sh"
	cmd := exec.Command("sh", "-c", cmdStr)
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))

	/*cmd := exec.Command("terraform", "version")
	cmdOutput := &bytes.Buffer{}
	cmd.Stdout = cmdOutput
	err := cmd.Run()
	if err != nil {
		os.Stderr.WriteString(err.Error())
	}
	fmt.Print(string(cmdOutput.Bytes()))*/
}
