package log

import "log"
import "fmt"

var DebugMode bool

func Init() {
	log.SetFlags(0)
	log.SetPrefix("gendao: ")
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func Infof(format string, args ...interface{}) {
	log.Printf(fmt.Sprintf("(info) %v", format), args...)
}

func Debugf(format string, args ...interface{}) {
	if DebugMode {
		log.Printf(fmt.Sprintf("(debug) %v", format), args...)
	}
}

func Debugln(v ...interface{}) {
	if DebugMode {
		log.Println(v...)
	}
}

func Errorf(format string, args ...interface{}) {
	log.Fatalf(fmt.Sprintf("(error) %v", format), args...)
}
