package cli

import (
	"bufio"
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"time"

	"github.com/murex/gamekit-coffeemachine/settings"
)

var (
	infoLog   = log.New(os.Stderr, "🟩 ", 0)
	errorLog  = log.New(os.Stderr, "🟥 ", 0)
	stdinLog  = log.New(os.Stderr, "⏩ ", 0)
	stdoutLog = log.New(os.Stderr, "⏪ ", 0)
)

// Run runs the coffee machine command line interface.
func Run(args []string) {
	infoLog.Println("starting coffee machine command line interface", settings.BuildVersion)
	start(parseArgs(args))
	infoLog.Println("closing coffee machine process runner")
}

// start starts the coffee machine command line interface.
// runDir is the directory where the coffee machine implementation is located.
// It must contain a run.sh script that starts the coffee machine implementation.
func start(runDir string) {
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

	go scanAndSend(stdin)

	if errStart := cmd.Start(); errStart != nil {
		errorLog.Fatalln(errStart)
	}

	go scanAndLog(stdout, stdoutLog)
	go scanAndLog(stderr, errorLog)

	if errWait := cmd.Wait(); errWait != nil {
		errorLog.Fatalln(errWait)
	}
}

func parseArgs(args []string) string {
	if len(args) < 2 {
		errorLog.Fatalf("syntax: %s <language-implementation-path>", path.Base(args[0]))
	}
	langImplPath, errAbs := filepath.Abs(args[1])
	if errAbs != nil {
		errorLog.Fatal(errAbs)
	}
	language := filepath.Base(langImplPath)

	infoLog.Println("implementation path is", langImplPath)
	infoLog.Println("implementation language is", language)
	return langImplPath
}

func scanAndSend(w io.WriteCloser) {
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
