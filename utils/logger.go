package utils

import (
	"log"
	"os"
)

var (
	InfoLog  *log.Logger
	ErrorLog *log.Logger
)

func InitLoggers() {
	green := "\033[32m"
	red := "\033[31m"
	reset := "\033[m"

	InfoLog = log.New(os.Stdout, green+"INFO\t"+reset, log.Ldate|log.Ltime)
	ErrorLog = log.New(os.Stdout, red+"ERROR\t"+reset, log.Ldate|log.Ltime|log.Lshortfile)

}
