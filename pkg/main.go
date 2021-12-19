package main

import "github.com/kenesparta/paackgo/app"

func main() {
	a := &app.App{}
	a.Initialize()
	a.Run()
}
