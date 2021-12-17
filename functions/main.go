package main

import (
	"errors"
	"fmt"
)

func main() {
	printHelloWorld()
	add(27, 27)
	port := 3000
	_, err := startWebServer(port, 3)
	logServerStartResult(err)

	_, failure := startWebServerFailure(port, 3)
	logServerStartResult(failure)
	logError(failure)
}

func printHelloWorld() {
	fmt.Println("Hello, World!")
}

func add(x, y int) {
	fmt.Println(x, "+", y, "=", x+y)
}

func startWebServer(port, numberOfRetries int) (int, error) {
	fmt.Printf("Trying to start Web Server on port %d...\n", port)
	fmt.Printf("Setting up number of retries with value %d\n", numberOfRetries)
	fmt.Println("Succesfully started web server!")
	return port, nil
}

func startWebServerFailure(port, numberOfRetries int) (int, error) {
	fmt.Printf("Trying to start Web Server on port %d...\n", port)
	return port, errors.New(fmt.Sprintf("Port %d already on use", port))
}

func logServerStartResult(err error) {
	fmt.Printf("Server started without errors: %t\n", isNull(err))
}

func logError(err error) {
	fmt.Printf("Error: %s\n", err)
}

func isNull(err error) bool {
	return err == nil
}
