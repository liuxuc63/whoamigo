package main

import "github.com/traefik/whoami/snippets"

func main() {
	snippets.ReplicaCDC()
}

func demo2() {
	snippets.RuneWaitX('x')
	snippets.RuneWaitX('b')
	snippets.RuneWaitX('X')
	snippets.Whereis()
	//snippets.Ping()
	//snippets.TestEnv()
	//go snippets.TestHandler()
	snippets.Ping()
}
