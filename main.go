package main

import (
	"io"
	"log"
	"os"
	"path/filepath"
	"time"

	"github.com/oxks/myauth/actions"
)

// main is the starting point for your Buffalo application.
// You can feel free and add to this `main` method, change
// what it does, etc...
// All we ask is that, at some point, you make sure to
// call `app.Serve()`, unless you don't want to start your
// application that is. :)
func main() {

	// log to file
	cwd, err := os.Getwd()
	if err != nil {
		log.Fatalf("Failed to determine working directory: %s", err)
	}
	runID := time.Now().Format("run-2006-01-02")
	logLocation := filepath.Join(cwd, runID+".log")
	logFile, err := os.OpenFile(logLocation, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Failed to open log file %s for output: %s", logLocation, err)
	}
	log.SetOutput(io.MultiWriter(os.Stderr, logFile))
	log.Println("testing logn 1")
	defer logFile.Close()

	// default main function content
	app := actions.App()
	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}

}

/*
# Notes about `main.go`

## SSL Support

We recommend placing your application behind a proxy, such as
Apache or Nginx and letting them do the SSL heavy lifting
for you. https://gobuffalo.io/en/docs/proxy

## Buffalo Build

When `buffalo build` is run to compile your binary, this `main`
function will be at the heart of that binary. It is expected
that your `main` function will start your application using
the `app.Serve()` method.

*/
