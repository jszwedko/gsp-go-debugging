package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/nu7hatch/gouuid"

	"../."
)

var (
	logLevels = map[string]logrus.Level{
		"debug": logrus.DebugLevel,
		"info":  logrus.InfoLevel,
		"warn":  logrus.WarnLevel,
		"error": logrus.ErrorLevel,
		"fatal": logrus.FatalLevel,
		"panic": logrus.PanicLevel,
	}

	logFormats = map[string]logrus.Formatter{
		"text": new(logrus.TextFormatter),
		"json": new(logrus.JSONFormatter),
	}

	logLevelString, logFormatString string
)

func init() {
	var (
		logLevelOptions  []string
		logFormatOptions []string
	)

	for s := range logLevels {
		logLevelOptions = append(logLevelOptions, s)
	}

	for s := range logFormats {
		logFormatOptions = append(logFormatOptions, s)
	}

	flag.StringVar(
		&logLevelString,
		"log-level",
		"info",
		fmt.Sprintf("Log level (options: %s)", strings.Join(logLevelOptions, ",")),
	)
	flag.StringVar(
		&logFormatString,
		"log-format",
		"text",
		fmt.Sprintf("Log format (options: %s)", strings.Join(logFormatOptions, ",")),
	)
}

func main() {
	flag.Parse()

	logger, err := newLogger(logLevelString, logFormatString)

	if err != nil {
		log.Fatal(err)
	}

	http.HandleFunc("/factorial/", func(w http.ResponseWriter, r *http.Request) {
		u4, err := uuid.NewV4()
		if err != nil {
			logger.Errorf("could not generate uuid: %v", err)
			http.Error(w, "", http.StatusInternalServerError)
		}

		logger.WithFields(logrus.Fields{
			"request": u4.String(),
		}).Infof("received request for %s", r.URL)

		pathParts := strings.Split(r.URL.Path, "/")

		if len(pathParts) < 3 || pathParts[2] == "" {
			logger.WithFields(logrus.Fields{
				"request": u4.String(),
			}).Infof("returning 404")
			http.Error(w, "", http.StatusNotFound)
			return
		}

		logger.WithFields(logrus.Fields{
			"request": u4.String(),
		}).Debugf("parsing resource %s as integer", pathParts[2])
		n, err := strconv.ParseInt(pathParts[2], 10, 64)
		if err != nil {
			logger.WithFields(logrus.Fields{
				"request": u4.String(),
			}).Infof("returning 400: %v", err)
			http.Error(w, fmt.Sprintf("could not parse as integer: %s", pathParts[2]), http.StatusBadRequest)
			return
		}

		time.Sleep(200 * time.Millisecond)

		logger.WithFields(logrus.Fields{
			"request": u4.String(),
		}).Debugf("calculating factorial for %d", n)

		f := math2.Factorial(n)

		time.Sleep(200 * time.Millisecond)

		logger.WithFields(logrus.Fields{
			"request": u4.String(),
		}).Infof("responding with %d", f)

		fmt.Fprintf(w, "%d", f)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newLogger(level, format string) (l *logrus.Logger, err error) {
	var ok bool

	l = logrus.New()

	l.Level, ok = logLevels[level]
	if !ok {
		return nil, fmt.Errorf("invalid log level %s", level)
	}

	l.Formatter, ok = logFormats[format]
	if !ok {
		return nil, fmt.Errorf("invalid log format %s", format)
	}

	return l, nil
}
