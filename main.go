package main

import (
	"bufio"
	"github.com/murex/gamekit-coffeemachine/settings"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"
)

var (
	infoLog   = log.New(os.Stderr, "🟩 ", 0)
	errorLog  = log.New(os.Stderr, "🟥 ", 0)
	stdinLog  = log.New(os.Stderr, "⏩ ", 0)
	stdoutLog = log.New(os.Stderr, "⏪ ", 0)
)

// main is the entry point of the coffee machine command line runner
func main() {
	infoLog.Println("starting coffee machine command line interface", settings.BuildVersion)
	runCli(parseArgs())
	infoLog.Println("closing coffee machine process runner")
}

// runCli starts the coffee machine command line interface.
// runDir is the directory where the coffee machine implementation is located.
// It must contain a run.sh script that starts the coffee machine implementation.
func runCli(runDir string) {
	cmd := exec.Command("bash", "./run.sh")
	cmd.Dir = runDir

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

	go parseAndForward(stdin)

	if errStart := cmd.Start(); errStart != nil {
		errorLog.Fatalln(errStart)
	}

	go scanAndLog(stdout, stdoutLog)
	go scanAndLog(stderr, errorLog)

	if errWait := cmd.Wait(); errWait != nil {
		errorLog.Fatalln(errWait)
	}
}

func parseArgs() string {
	if len(os.Args) < 2 {
		errorLog.Fatalf("syntax: %s <language-implementation-path>", path.Base(os.Args[0]))
	}
	langImplPath, errAbs := filepath.Abs(os.Args[1])
	if errAbs != nil {
		errorLog.Fatal(errAbs)
	}
	language := filepath.Base(langImplPath)

	infoLog.Println("implementation path is", langImplPath)
	infoLog.Println("implementation language is", language)
	return langImplPath
}

func parseAndForward(w io.WriteCloser) {
	defer func(stdin io.WriteCloser) {
		_ = stdin.Close()
	}(w)
	const invite = "enter command to send to the coffee machine"
	infoLog.Println(invite)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		msg := strings.TrimSpace(strings.ToLower(scanner.Text()))
		stdinLog.Println(msg)
		_, err := io.WriteString(w, msg+"\n")
		if err != nil {
			errorLog.Fatalln(err)
		}
		time.Sleep(100 * time.Millisecond)
		infoLog.Println(invite)
	}

	if errScanner := scanner.Err(); errScanner != nil {
		errorLog.Fatalln(errScanner)
	}
}

func scanAndLog(r io.ReadCloser, logger *log.Logger) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		logger.Println(scanner.Text())
	}
	if errScanner := scanner.Err(); errScanner != nil {
		errorLog.Fatalln(errScanner)
	}
}
