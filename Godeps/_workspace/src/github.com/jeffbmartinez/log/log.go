package log

import (
	"fmt"
	"log"
	"os"
)

const (
	INFO_PREFIX  = "[INFO]"
	WARN_PREFIX  = "[WARN]"
	ERROR_PREFIX = "[ERROR]"
	FATAL_PREFIX = "[FATAL]"
	PANIC_PREFIX = "[PANIC]"
)

func Print(v ...interface{}) {
	log.Print(v)
}

func Printf(format string, v ...interface{}) {
	log.Printf(format, v)
}

func Info(v ...interface{}) {
	log.Printf("%v %v\n", INFO_PREFIX, fmt.Sprint(v...))
}

func Infof(format string, v ...interface{}) {
	log.Printf("%v %v\n", INFO_PREFIX, fmt.Sprintf(format, v...))
}

func Warn(v ...interface{}) {
	log.Printf("%v %v\n", WARN_PREFIX, fmt.Sprint(v...))
}

func Warnf(format string, v ...interface{}) {
	log.Printf("%v %v\n", WARN_PREFIX, fmt.Sprintf(format, v...))
}

func Error(v ...interface{}) {
	log.Printf("%v %v\n", ERROR_PREFIX, fmt.Sprint(v...))
}

func Errorf(format string, v ...interface{}) {
	log.Printf("%v %v\n", ERROR_PREFIX, fmt.Sprintf(format, v...))
}

func Fatal(v ...interface{}) {
	log.Printf("%v %v\n", FATAL_PREFIX, fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	log.Printf("%v %v\n", FATAL_PREFIX, fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Panic(v ...interface{}) {
	message := fmt.Sprint(v...)
	log.Printf("%v %v\n", PANIC_PREFIX, message)
	panic(message)
}

func Panicf(format string, v ...interface{}) {
	message := fmt.Sprintf(format, v...)
	log.Printf("%v %v\n", PANIC_PREFIX, message)
	panic(message)
}
