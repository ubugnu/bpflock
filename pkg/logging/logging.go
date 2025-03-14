// SPDX-License-Identifier: Apache-2.0
// Copyright 2021 Djalal Harouni
// Copyright 2016-2021 Authors of Cilium

package logging

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"sync/atomic"

	"github.com/linux-lock/bpflock/pkg/logging/logfields"
	"github.com/sirupsen/logrus"
)

type LogFormat string

const (
	Syslog    = "syslog"
	LevelOpt  = "level"
	FormatOpt = "format"

	LogFormatText LogFormat = "text"
	LogFormatJSON LogFormat = "json"

	// DefaultLogFormat is the string representation of the default logrus.Formatter
	// we want to use (possible values: text or json)
	DefaultLogFormat LogFormat = LogFormatText

	// DefaultLogLevel is the default log level we want to use for our logrus.Formatter
	DefaultLogLevel logrus.Level = logrus.InfoLevel
)

var (
	// DefaultLogger is the base logrus logger. It is different from the logrus
	// default to avoid external dependencies from writing out unexpectedly
	DefaultLogger = InitializeDefaultLogger()
)

// LogOptions maps configuration key-value pairs related to logging.
type LogOptions map[string]string

// InitializeDefaultLogger returns a logrus Logger with a custom text formatter.
func InitializeDefaultLogger() (logger *logrus.Logger) {
	logger = logrus.New()
	logger.SetFormatter(GetFormatter(DefaultLogFormat))
	logger.SetLevel(DefaultLogLevel)
	return
}

// GetLogLevel returns the log level specified in the provided LogOptions. If
// it is not set in the options, it will return the default level.
func (o LogOptions) GetLogLevel() (level logrus.Level) {
	levelOpt, ok := o[LevelOpt]
	if !ok {
		return DefaultLogLevel
	}

	var err error
	if level, err = logrus.ParseLevel(levelOpt); err != nil {
		logrus.WithError(err).Warning("Ignoring user-configured log level")
		return DefaultLogLevel
	}

	return
}

// GetLogFormat returns the log format specified in the provided LogOptions. If
// it is not set in the options or is invalid, it will return the default format.
func (o LogOptions) GetLogFormat() LogFormat {
	formatOpt, ok := o[FormatOpt]
	if !ok {
		return DefaultLogFormat
	}

	formatOpt = strings.ToLower(formatOpt)
	re := regexp.MustCompile(`^(text|json)$`)
	if !re.MatchString(formatOpt) {
		logrus.WithError(
			fmt.Errorf("incorrect log format configured '%s', expected 'text' or 'json'", formatOpt),
		).Warning("Ignoring user-configured log format")
		return DefaultLogFormat
	}

	return LogFormat(formatOpt)
}

// Sets the subsys logger and returns a new log entry from a logrus Logger
func GetLogSubsys(subsys string) *logrus.Entry {
	return DefaultLogger.WithField(logfields.LogSubsys, subsys)
}

// Sets the ebpf program loffer and returns a new log entry
func GetLogBpfsubsys(bpfprog string) *logrus.Entry {
	return DefaultLogger.WithField(logfields.LogBpfSubsys, bpfprog)
}

// SetLogOutput change the DefaultLogger output
func SetLogOutput(out io.Writer) {
	DefaultLogger.SetOutput(out)
}

func ResetLogOutput() {
	DefaultLogger.SetOutput(os.Stdout)
}

// SetLogLevel updates the DefaultLogger with a new logrus.Level
func SetLogLevel(logLevel logrus.Level) {
	DefaultLogger.SetLevel(logLevel)
}

// SetDefaultLogLevel updates the DefaultLogger with the DefaultLogLevel
func SetDefaultLogLevel() {
	DefaultLogger.SetLevel(DefaultLogLevel)
}

// SetLogLevelToDebug updates the DefaultLogger with the logrus.DebugLevel
func SetLogLevelToDebug() {
	DefaultLogger.SetLevel(logrus.DebugLevel)
}

// SetLogFormat updates the DefaultLogger with a new LogFormat
func SetLogFormat(logFormat LogFormat) {
	DefaultLogger.SetFormatter(GetFormatter(logFormat))
}

// SetLogLevel updates the DefaultLogger with the DefaultLogFormat
func SetDefaultLogFormat() {
	DefaultLogger.SetFormatter(GetFormatter(DefaultLogFormat))
}

// SetupLogging sets up each logging service provided in loggers and configures
// each logger with the provided logOpts.
func SetupLogging(loggers []string, logOpts LogOptions, tag string, debug bool) error {
	// Updating the default log format
	SetLogFormat(logOpts.GetLogFormat())

	// Set default logger to output to stdout if no loggers are provided.
	if len(loggers) == 0 {
		// TODO: switch to a per-logger version when we upgrade to logrus >1.0.3
		logrus.SetOutput(os.Stdout)
	}

	// Updating the default log level, overriding the log options if the debug arg is being set
	if debug {
		SetLogLevelToDebug()
	} else {
		SetLogLevel(logOpts.GetLogLevel())
	}

	// always suppress the default logger so libraries don't print things
	logrus.SetLevel(logrus.PanicLevel)

	// Iterate through all provided loggers and configure them according
	// to user-provided settings.
	for _, logger := range loggers {
		switch logger {
		case Syslog:
			if err := setupSyslog(logOpts, tag, debug); err != nil {
				return fmt.Errorf("failed to set up syslog: %w", err)
			}
		default:
			return fmt.Errorf("provided log driver %q is not a supported log driver", logger)
		}
	}

	return nil
}

// GetFormatter returns a configured logrus.Formatter with some specific values
// we want to have
func GetFormatter(format LogFormat) logrus.Formatter {
	switch format {
	case LogFormatText:
		return &logrus.TextFormatter{
			DisableTimestamp: true,
			DisableColors:    true,
		}
	case LogFormatJSON:
		return &logrus.JSONFormatter{
			DisableTimestamp: true,
		}
	}

	return nil
}

// validateOpts iterates through all of the keys and values in logOpts, and errors out if
// the key in logOpts is not a key in supportedOpts, or the value of corresponding key is
// not listed in the value of validKVs.
func (o LogOptions) validateOpts(logDriver string, supportedOpts map[string]bool, validKVs map[string][]string) error {
	for k, v := range o {
		if !supportedOpts[k] {
			return fmt.Errorf("provided configuration key %q is not supported as a logging option for log driver %s", k, logDriver)
		}
		if validValues, ok := validKVs[k]; ok {
			valid := false
			for _, vv := range validValues {
				if v == vv {
					valid = true
					break
				}
			}
			if !valid {
				return fmt.Errorf("provided configuration value %q is not a valid value for %q in log driver %s, valid values: %v", v, k, logDriver, validValues)
			}

		}
	}
	return nil
}

// getLogDriverConfig returns a map containing the key-value pairs that start
// with string logDriver from map logOpts.
func getLogDriverConfig(logDriver string, logOpts LogOptions) LogOptions {
	keysToValidate := make(LogOptions)
	for k, v := range logOpts {
		ok, err := regexp.MatchString(logDriver+".*", k)
		if err != nil {
			DefaultLogger.Fatal(err)
		}
		if ok {
			keysToValidate[k] = v
		}
	}
	return keysToValidate
}

// MultiLine breaks a multi line text into individual log entries and calls the
// logging function to log each entry
func MultiLine(logFn func(args ...interface{}), output string) {
	scanner := bufio.NewScanner(bytes.NewReader([]byte(output)))
	for scanner.Scan() {
		logFn(scanner.Text())
	}
}

// CanLogAt returns whether a log message at the given level would be
// logged by the given logger.
func CanLogAt(logger *logrus.Logger, level logrus.Level) bool {
	return GetLevel(logger) >= level
}

// GetLevel returns the log level of the given logger.
func GetLevel(logger *logrus.Logger) logrus.Level {
	return logrus.Level(atomic.LoadUint32((*uint32)(&logger.Level)))
}
