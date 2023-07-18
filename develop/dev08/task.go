package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("$ ")

		// Read from stdIn
		command, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("Exiting")
				os.Exit(0)
			}
			fmt.Println("Reading error :", err)
			continue
		}

		// Удаление символа новой строки из команды
		command = strings.TrimSuffix(command, "\n")

		// get args
		args := strings.Fields(command)
		if len(args) == 0 {
			continue
		}


		// Catch command
		switch args[0] {
		case ":q":
			fmt.Println("Exiting")
			return
		case "cd":
			if len(args) < 2 {
				fmt.Println("No dir")
				continue
			}

			err := os.Chdir(args[1])
			if err != nil {
				fmt.Println("Change dir error:", err)
			}
		case "pwd":
			dir, err := os.Getwd()
			if err != nil {
				fmt.Println("Pwd error:", err)
			}
			fmt.Println(dir)
		case "echo":
			if len(args) < 2 {
				fmt.Println("No args for Echo")
				continue
			}

			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Println("No proc for kill")
				continue
			}

			pid := args[1]
			// Kill for Windows
			cmd := exec.Command("taskkill", "/F", "/PID", pid)
			err := cmd.Run()
			if err != nil {
				fmt.Println("Kill error:", err)
			}
		case "ps":
			// tasklist in windows
			cmd := exec.Command("tasklist")
			output, err := cmd.Output()
			if err != nil {
				fmt.Println("Ps error:", err)
			}
			fmt.Println(string(output))
		default:
			// If not a command
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stderr = os.Stderr
			err := cmd.Run()
			if err != nil {
				fmt.Println("Run command error:", err)
			}
		}
	}
}