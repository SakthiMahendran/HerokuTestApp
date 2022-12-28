package main

import (
	"fmt"
	"os"

	"github.com/SakthiMahendran/HerokuTestApp/testserver"
)

func main() {
	file, err := os.OpenFile("logFile.txt", os.O_CREATE, os.ModeAppend)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	ts := testserver.TestServer{LogFile: file}

	ts.StartServing()
}
