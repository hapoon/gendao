package log

import "log"

var DebugMode bool

func Init() {
	log.SetFlags(0)
	log.SetPrefix("gendao: ")
}

func Fatalln(v ...interface{}) {
	log.Fatalln(v...)
}

func Infof(format string, args ...interface{}) {
	log.Printf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	if DebugMode {
		log.Printf(format, args...)
	}
}

func Debugln(v ...interface{}) {
	if DebugMode {
		log.Println(v...)
	}
}

func Errorf(format string, args ...interface{}) {
	log.Fatalf(format, args...)
}
