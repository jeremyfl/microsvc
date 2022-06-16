package main

import (
	"gitlab.com/jeremylo/microsvc/ordersvc/cmd"
	"os"
)

func main() {
	var argsRaw = os.Args
	flag := argsRaw[1]

	switch flag {
	case "listen":
		cmd.Listen()

		return
	default:
		cmd.Serve()

		return
	}
}
