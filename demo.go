package main

import "github.com/traefik/whoami/snippets"

func main() {
	//snippets.Ping()
	snippets.TestEnv()
	go snippets.HandlerSignal()
	snippets.Ping()
}
