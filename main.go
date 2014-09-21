// This program provides a sample web service that uses concurrency
// and channels to perform a coordinated set of asynchronous searches.
package main

import (
	"log"

	"github.com/goinggo/concurrentwebservice/service"
)

// init is called before main. We are using init to
// set the logging package.
func init() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
}

// main is the entry point for the application.
func main() {
	service.Start()
}
