package main

import (
	"io"
	"log"
	"os"
)

func checkError(err error) {
	if err != nil {
		logError()
	}
}

func logError() {
	f, err := os.OpenFile("orders.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	checkError(err)
	defer f.Close()
	wrt := io.MultiWriter(os.Stdout, f)
	log.SetOutput(wrt)
}
