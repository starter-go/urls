package main

import (
	"os"

	"github.com/starter-go/starter"
	"github.com/starter-go/urls/modules/urls"
)

func main() {
	m := urls.Module()
	i := starter.Init(os.Args)
	i.MainModule(m)
	i.WithPanic(true).Run()
}
