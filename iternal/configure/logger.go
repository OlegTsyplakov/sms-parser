package configure

import (
	"log"
	"os"
)

type logLevel int

const (
	info logLevel = iota
	warn
	err
)

func (ll logLevel) string() string {
	return [...]string{"[info]", "[warn]", "[error]"}[ll]
}

type iLogger interface {
	configure(env *env) iStartLogger
}

type iStartLogger interface {
	Start() iLogLevel
}
type iLogLevel interface {
	Information(message string)
	Warning(message string)
	Error(message string)
}
type logger struct {
	env  *env
	file *os.File
}

func newLogger() iLogger {
	return &logger{}
}

func (l *logger) configure(env *env) iStartLogger {
	l.env = env
	return l
}
func (l *logger) Start() iLogLevel {
	file, err := openLogFile(l.env.LogFolder)
	if err != nil {
		log.Fatal(err)
	}
	l.file = file
	return l
}

func openLogFile(path string) (*os.File, error) {

	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func (l *logger) Information(message string) {
	var ll logLevel = info
	customLog := log.New(l.file, ll.string(), log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	customLog.Println(message)
}

func (l *logger) Warning(message string) {
	var ll logLevel = warn
	customLog := log.New(l.file, ll.string(), log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	customLog.Println(message)
}

func (l *logger) Error(message string) {
	var ll logLevel = err
	customLog := log.New(l.file, ll.string(), log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	customLog.Println(message)
}
