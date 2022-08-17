package main

import (
	"fmt"

	"github.com/sirupsen/logrus"
)

func main() {
	name := "Go Developers"
	fmt.Println("TEST", name)
	l := logrus.New()

	l.Info("This is an info")
	l.Warn("This is a warning")
	l.Error("This is an error")
}
