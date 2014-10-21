package llog

import (
	"fmt"
	"log"
	"os"
	"time"
)

type LogLevel int8
type LogType int8

var logger *log.Logger
var filelog *log.Logger
var minLevel LogLevel

const (
	TRACE LogLevel = iota
	DEBUG
	INFO
	WARN
	ERROR
	FATAL
	PANIC
)

const (
	CONSOLE LogType = iota + 1
	FILE
)

// Call Initialize after setting (or not setting) SyslogHost and SyslogPort when
// they're read from configuration source.
func init() {
	minLevel = TRACE
	logger = log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
}

func SetLogLevel(level LogLevel) {
	minLevel = level
}

func GetLogLevel() LogLevel {
	return minLevel
}

func ResetLogger() {
	logger = nil
	filelog = nil
}

func AddLogger(logtype LogType) {
	if logtype == CONSOLE {
		logger = log.New(os.Stdout, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	} else if logtype == FILE {
		filename := fmt.Sprintf("./%s.llog", time.Now().String())
		fmt.Printf("Log initialize:%s\n", filename)
		logfile, err := os.OpenFile(filename, os.O_RDWR|os.O_CREATE, 0600)
		if err != nil {
			Panic(err)
		}
		filelog = log.New(logfile, "\r\n", log.Ldate|log.Ltime|log.Lshortfile)
	}
}

func DelLogger(logtype LogType) {
	if logtype == CONSOLE {
		logger = nil
	} else if logtype == FILE {
		filelog = nil
	}
}

func _log(s string) {
	if logger != nil {
		logger.Output(4, s)
	}
	if filelog != nil {
		filelog.Output(4, s)
	}
}

func Panic(messages ...interface{}) {
	panic_(messages...)
}

func Panicf(format string, messages ...interface{}) {
	panicf(format, messages...)
}

func Fatal(messages ...interface{}) {
	if minLevel <= FATAL {
		fatal(messages...)
	}
}

func Fatalf(format string, messages ...interface{}) {
	if minLevel <= FATAL {
		fatalf(format, messages...)
	}
}

func Error(messages ...interface{}) {
	if minLevel <= ERROR {
		print(messages...)
	}
}

func Errorf(format string, messages ...interface{}) {
	if minLevel <= ERROR {
		printf(format, messages...)
	}
}

func Warn(messages ...interface{}) {
	if minLevel <= WARN {
		print(messages...)
	}
}

func Warnf(format string, messages ...interface{}) {
	if minLevel <= WARN {
		printf(format, messages...)
	}
}

func Info(messages ...interface{}) {
	if minLevel <= INFO {
		print(messages...)
	}
}

func Infof(format string, messages ...interface{}) {
	if minLevel <= INFO {
		printf(format, messages...)
	}
}

func Debug(messages ...interface{}) {
	if minLevel <= DEBUG {
		print(messages...)
	}
}

func Debugf(format string, messages ...interface{}) {
	if minLevel <= DEBUG {
		printf(format, messages...)
	}
}

func Trace(messages ...interface{}) {
	if minLevel <= TRACE {
		print(messages...)
	}
}

func Tracef(format string, messages ...interface{}) {
	if minLevel <= TRACE {
		printf(format, messages...)
	}
}

// Printf calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Printf.
func printf(format string, v ...interface{}) {
	_log(fmt.Sprintf(format, v...))
}

// Print calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Print.
func print(v ...interface{}) { _log(fmt.Sprint(v...)) }

// Println calls l.Output to print to the logger.
// Arguments are handled in the manner of fmt.Println.
func println(v ...interface{}) { _log(fmt.Sprintln(v...)) }

// Fatal is equivalent to l.Print() followed by a call to os.Exit(1).
func fatal(v ...interface{}) {
	_log(fmt.Sprint(v...))
	os.Exit(1)
}

// Fatalf is equivalent to l.Printf() followed by a call to os.Exit(1).
func fatalf(format string, v ...interface{}) {
	_log(fmt.Sprintf(format, v...))
	os.Exit(1)
}

// Fatalln is equivalent to l.Println() followed by a call to os.Exit(1).
func fatalln(v ...interface{}) {
	_log(fmt.Sprintln(v...))
	os.Exit(1)
}

// Panic is equivalent to l.Print() followed by a call to panic().
func panic_(v ...interface{}) {
	s := fmt.Sprint(v...)
	_log(s)
	panic(s)
}

// Panicf is equivalent to l.Printf() followed by a call to panic().
func panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	_log(s)
	panic(s)
}

// Panicln is equivalent to l.Println() followed by a call to panic().
func panicln(v ...interface{}) {
	s := fmt.Sprintln(v...)
	_log(s)
	panic(s)
}

func Println(level LogLevel, messages ...interface{}) {

	switch level {
	case DEBUG:
		Debugf("%v", messages)
	case TRACE:
		Tracef("%v", messages)
	case INFO:
		Infof("%v", messages)
	case WARN:
		Warnf("%v", messages)
	case ERROR:
		Errorf("%v", messages)
	case FATAL:
		Fatalf("%v", messages)
	case PANIC:
		Panicf("%v", messages)
	}

	return
}

func Printf(level LogLevel, format string, messages ...interface{}) {

	switch level {
	case DEBUG:
		Debugf(format, messages)
	case TRACE:
		Tracef(format, messages)
	case INFO:
		Infof(format, messages)
	case WARN:
		Warnf(format, messages)
	case ERROR:
		Errorf(format, messages)
	case FATAL:
		Fatalf(format, messages)
	case PANIC:
		Panicf(format, messages)
	}

	return
}

func fromMulti(messages ...interface{}) string {
	var r string
	for x := 0; x < len(messages); x++ {
		r = r + messages[x].(string)
		if x < len(messages) {
			r = r + "  "
		}
	}
	return r
}

func LevelFromString(l string) (level LogLevel) {
	switch l {
	case "DEBUG":
		level = DEBUG
	case "TRACE":
		level = TRACE
	case "INFO":
		level = INFO
	case "WARN":
		level = WARN
	case "ERROR":
		level = ERROR
	case "FATAL":
		level = FATAL
	case "PANIC":
		level = PANIC
	}

	return
}
