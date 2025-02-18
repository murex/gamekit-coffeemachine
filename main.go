package main

import (
	"bufio"
	"github.com/murex/coffee-machine/progress-runner/settings"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"
)

var (
	infoLog   = log.New(os.Stderr, "🟩 ", 0)
	errorLog  = log.New(os.Stderr, "🟥 ", 0)
	stdinLog  = log.New(os.Stderr, "⏩ ", 0)
	stdoutLog = log.New(os.Stderr, "⏪ ", 0)
)

func main() {
	if len(os.Args) < 2 {
		errorLog.Fatalf("syntax: %s <language>", path.Base(os.Args[0]))
	}
	language := os.Args[1]

	infoLog.Println("starting coffee machine process runner", settings.BuildVersion)
	infoLog.Println("language implementation is", language)

	cmd := exec.Command("bash", "./run.sh")
	cmd.Dir = path.Join("..", language)

	stdin, errStdinPipe := cmd.StdinPipe()
	if errStdinPipe != nil {
		errorLog.Fatal(errStdinPipe)
	}

	stdout, errStdoutPipe := cmd.StdoutPipe()
	if errStdoutPipe != nil {
		errorLog.Fatalln(errStdoutPipe)
	}

	stderr, errStderrPipe := cmd.StderrPipe()
	if errStderrPipe != nil {
		errorLog.Fatalln(errStderrPipe)
	}

	go func() {
		defer func(stdin io.WriteCloser) {
			_ = stdin.Close()
		}(stdin)
		infoLog.Println("enter command to send to the coffee machine")
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			msg := strings.TrimSpace(strings.ToLower(scanner.Text()))
			stdinLog.Println(msg)
			_, err := io.WriteString(stdin, msg+"\n")
			if err != nil {
				errorLog.Fatalln(err)
			}
			time.Sleep(100 * time.Millisecond)
			infoLog.Println("enter command to send to the coffee machine")
		}

		if errScanner := scanner.Err(); errScanner != nil {
			errorLog.Fatalln(errScanner)
		}
	}()

	if errStart := cmd.Start(); errStart != nil {
		errorLog.Fatalln(errStart)
	}

	go func() {
		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			stdoutLog.Println(scanner.Text())
		}
		if errScanner := scanner.Err(); errScanner != nil {
			errorLog.Fatalln(errScanner)
		}
	}()

	go func() {
		scanner := bufio.NewScanner(stderr)
		for scanner.Scan() {
			errorLog.Println(scanner.Text())
		}
		if errScanner := scanner.Err(); errScanner != nil {
			errorLog.Fatalln(errScanner)
		}
	}()

	if errWait := cmd.Wait(); errWait != nil {
		errorLog.Fatalln(errWait)
	}

	infoLog.Println("closing coffee machine process runner")
}
