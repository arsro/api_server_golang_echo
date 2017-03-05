package main

import (
	"github.com/Sirupsen/logrus"
	
	"api/route"
)

func init() {
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetFormatter(&logrus.JSONFormatter{})
}

func main() {
	e := route.Init()
	// Start server
	e.Logger.Fatal(e.Start(":3000"))
}
