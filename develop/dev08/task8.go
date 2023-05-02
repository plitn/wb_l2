package dev08

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"syscall"
)

func Shell() {
	r := bufio.NewReader(os.Stdin)
	for {
		pwd, _ := os.Getwd()
		fmt.Println("%s $ ", pwd)
		in, err := r.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				fmt.Println("exiting")
				os.Exit(0)
			}
			fmt.Fprintf(os.Stderr, "input reading err: ", err)
			continue
		}

		in = strings.TrimSpace(in)
		if in == "" {
			continue
		}

		args := strings.Split(in, " ")
		switch args[0] {
		case "cd":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "cd usage err")
				continue
			}
			if err := os.Chdir(args[1]); err != nil {
				fmt.Fprintf(os.Stderr, "change directory err: ", err)
			}
		case "pwd":
			pwd, err := os.Getwd()
			if err != nil {
				fmt.Fprintf(os.Stderr, "getwd err: ", err)
			}
			fmt.Println(pwd)
		case "echo":
			fmt.Println(strings.Join(args[1:], " "))
		case "kill":
			if len(args) < 2 {
				fmt.Fprintf(os.Stderr, "usage err: ")
			}
			pid, err := strconv.Atoi(args[1])
			if err != nil {
				fmt.Fprintf(os.Stderr, "pid err: ", err)
				continue
			}
			if err := syscall.Kill(pid, syscall.SIGTERM); err != nil {
				fmt.Fprintf(os.Stderr, "killing err: ", err)
			}
		case "ps":
			cmd := exec.Command("ps", "aux")
			cmd.Stdout = os.Stdout
			if err := cmd.Run(); err != nil {
				fmt.Fprintf(os.Stderr, "ps err: ", err)
			}
		default:
			cmd := exec.Command(args[0], args[1:]...)
			cmd.Stdout = os.Stdout
			cmd.Stdin = os.Stdin
			if err := cmd.Start(); err != nil {
				fmt.Fprintf(os.Stderr, "cmd start err: ", err)
				continue
			}
			if err := cmd.Wait(); err != nil {
				fmt.Fprintf(os.Stderr, "cmd waiting err: ", err)
			}
		}
	}
}
