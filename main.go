package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
	"time"
)

var help = fmt.Sprintf("Usage: %s <duration> <program path> [program args...]\n\tduration is a string accepted by Go's time.ParseDuration (https://golang.org/pkg/time/#ParseDuration)", os.Args[0])

func printHelp() {
	fmt.Println(help)
}

func main() {
	if len(os.Args) < 3 {
		printHelp()
		os.Exit(1)
	}

	dur, err := time.ParseDuration(os.Args[1])
	if err != nil {
		printHelp()
		fmt.Printf("\nInvalid duration \"%s\": %v\n", os.Args[1], err)
		os.Exit(1)
	}
	var args []string
	if len(os.Args) == 3 {
		args = make([]string, 0)
	} else {
		args = os.Args[3:]
	}

	for {
		c := exec.Command(os.Args[2], args...)
		c.Stdin = os.Stdin
		c.Stdout = os.Stdout
		c.Stderr = os.Stderr

		log.Println("INFO: Running:", os.Args[2], strings.Join(args, " "))
		err = c.Run()
		if err != nil {
			log.Println("WARNING: Non-zero exit status:", err)
		}

		time.Sleep(dur)
	}
}
