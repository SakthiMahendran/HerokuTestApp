package main

import (
	"fmt"
	"os"
	"time"

	"github.com/SakthiMahendran/HerokuTestApp/testserver"
)

func main() {
	file, err := os.OpenFile(getLogFileName(), os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("err: %v\n", err)
	}
	ts := testserver.TestServer{LogFile: file}

	ts.StartServing()
}

func getLogFileName() string {
	t := time.Now()

	day := t.Day()
	month := int(t.Month())
	year := t.Year()

	fileName := fmt.Sprint(day, "_", month, "_", year, "_log.txt")

	return fileName
}
