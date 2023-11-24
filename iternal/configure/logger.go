package configure

import (
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

type logLevel int

const (
	info logLevel = iota
	warn
	err
)

func (ll logLevel) toString() string {
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
	env       *env
	file      *os.File
	timestamp time.Time
}

func newLogger() iLogger {
	return &logger{
		timestamp: time.Now().Round(time.Second),
	}
}

func (l *logger) configure(env *env) iStartLogger {
	l.env = env
	return l
}
func (l *logger) Start() iLogLevel {
	file, err := l.openLogFile(l.env.LogFolder)
	if err != nil {
		log.Fatal(err)
	}
	l.file = file
	return l
}

func (l *logger) openLogFile(folder string) (*os.File, error) {
	file := l.createLogFile()
	path := filepath.Join(folder, file)
	logFile, err := os.OpenFile(path, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	return logFile, nil
}

func (l *logger) createLogFile() string {
	var file strings.Builder
	file.WriteString("log-")
	file.WriteString(l.timestamp.Format("2006-01-02_150405"))
	file.WriteString(".log")
	return file.String()
}

func (l *logger) Information(message string) {
	var ll logLevel = info
	l.print(ll, message)
}
func (l *logger) Warning(message string) {
	var ll logLevel = warn
	l.print(ll, message)
}

func (l *logger) Error(message string) {
	var ll logLevel = err
	l.print(ll, message)
}

func (l *logger) print(ll logLevel, message string) {
	customLog := log.New(l.file, ll.toString(), log.LstdFlags|log.Lshortfile|log.Lmicroseconds)
	customLog.Println(message)
}
