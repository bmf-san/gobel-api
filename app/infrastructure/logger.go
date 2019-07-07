package infrastructure

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/bmf-san/gobel-api/app/usecases"
	"github.com/pkg/errors"
)

// errorLogFile is path to error log file.
const errorLogFile = "./log/error.log"

// accessLogFile is path to access log file.
const accessLogFile = "./log/access.log"

// A Logger represents a logger.
type Logger struct {
	errorLogFile  string
	accessLogFile string
}

// NewLogger creates a logger.
func NewLogger() usecases.Logger {
	return &Logger{
		errorLogFile:  errorLogFile,
		accessLogFile: accessLogFile,
	}
}

// LogError writes a log for an error log.
func (l *Logger) LogError(e error) {
	file, err := os.OpenFile(l.errorLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%+v\n", errors.WithStack(err))
	}
	defer file.Close()

	// TODO: Optimize log format
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf("%+v\n", errors.WithStack(e))
}

// LogAccess writes a log for an access log.
func (l *Logger) LogAccess(r *http.Request) {
	// TODO: To implement log lotation.
	file, err := os.OpenFile(l.accessLogFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Printf("%+v\n", errors.WithStack(err))
	}
	defer file.Close()

	// TODO: Optimize log format
	log.SetOutput(io.MultiWriter(file, os.Stdout))
	log.SetFlags(log.Ldate | log.Ltime)

	log.Printf("%s %s %s\n", r.RemoteAddr, r.Method, r.URL)
}
