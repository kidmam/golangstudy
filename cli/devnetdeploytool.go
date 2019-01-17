package main

import (
	"bytes"
	"github.com/spf13/cobra"
	"io"
	"log"
	"os"
	"os/exec"
	"regexp"
	"strconv"
)

func getRatio(text string) float32 {
	re1, _ := regexp.Compile(`step[\s]+(\d+)`)
	result := re1.FindStringSubmatch(text)
	val, _ := strconv.Atoi(result[1])
	return float32(val) / 10
}

func main() {

	var cmdInit = &cobra.Command{
		Use:   "init",
		Short: "Terraform init",
		Long:  `Terraform init`,
		Args:  cobra.MinimumNArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Print: " + strings.Join(args, " "))

			cmdStr := "init.sh"
			command := exec.Command("sh", "-e", cmdStr)

			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			command.Stdout = mw
			command.Stderr = mw

			// Execute the command
			if err := command.Run(); err != nil {
				log.Panic(err)
			}

			log.Println(stdBuffer.String())

		},
	}

	var cmdPlan = &cobra.Command{
		Use:   "plan [devnetname number_of_peers dockertag rsregion]",
		Short: "Terraform plan",
		Long: `Terraform plan
 - var devnetname: unique devnet name
 - var number_of_peers: peer count
 - var dockertag: docker image tag
 - var rsregion: radiostation install region`,
		Args: cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Print: " + strings.Join(args, " "))

			cmdStr := "plan.sh"
			command := exec.Command("sh", "-e", cmdStr, args[0], args[1], args[2], args[3])

			/*cmdOutput := &bytes.Buffer{}
			command.Stdout = cmdOutput
			err := command.Run()
			if err != nil {
				os.Stderr.WriteString(err.Error())
			}
			fmt.Print(string(cmdOutput.Bytes())) */

			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			command.Stdout = mw
			command.Stderr = mw

			// Execute the command
			if err := command.Run(); err != nil {
				log.Panic(err)
			}

			log.Println(stdBuffer.String())

		},
	}

	var cmdApply = &cobra.Command{
		Use:   "apply [devnetname number_of_peers dockertag rsregion]",
		Short: "Terraform apply",
		Long: `Terraform apply
 - var devnetname: unique devnet name
 - var number_of_peers: peer count
 - var dockertag: docker image tag
 - var rsregion: radiostation install region`,
		Args: cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {

			cmdStr := "apply.sh"
			command := exec.Command("sh", "-e", cmdStr, args[0], args[1], args[2], args[3])

			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			command.Stdout = mw
			command.Stderr = mw

			// Execute the command
			if err := command.Run(); err != nil {
				log.Panic(err)
			}

			log.Println(stdBuffer.String())

		},
	}

	var cmdDestroy = &cobra.Command{
		Use:   "destroy [devnetname number_of_peers dockertag rsregion]",
		Short: "Terraform destroy",
		Long: `Terraform destroy
 - var devnetname: unique devnet name
 - var number_of_peers: peer count
 - var dockertag: docker image tag
 - var rsregion: radiostation install region`,
		Args: cobra.MinimumNArgs(4),
		Run: func(cmd *cobra.Command, args []string) {

			cmdStr := "destroy.sh"
			command := exec.Command("sh", "-e", cmdStr, args[0], args[1], args[2], args[3])

			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			command.Stdout = mw
			command.Stderr = mw

			// Execute the command
			if err := command.Run(); err != nil {
				log.Panic(err)
			}

			log.Println(stdBuffer.String())

		},
	}

	var cmdOutput = &cobra.Command{
		Use:   "output",
		Short: "Terraform output [modulename]",
		Long: `Terraform output
 -module: main.tf에 정의된 module 이름`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			//fmt.Println("Print: " + strings.Join(args, " "))

			cmdStr := "output.sh"
			command := exec.Command("sh", "-e", cmdStr, args[0])

			var stdBuffer bytes.Buffer
			mw := io.MultiWriter(os.Stdout, &stdBuffer)

			command.Stdout = mw
			command.Stderr = mw

			// Execute the command
			if err := command.Run(); err != nil {
				log.Panic(err)
			}

			log.Println(stdBuffer.String())

		},
	}

	var rootCmd = &cobra.Command{Use: "devnetdeploytool"}
	rootCmd.AddCommand(cmdInit, cmdPlan, cmdApply, cmdDestroy, cmdOutput)
	rootCmd.Execute()
}
