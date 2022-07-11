package util

import (
	"flag"
	"log/syslog"
	"strconv"

	"k8s.io/klog/v2"
)

type LoggingVerbosityLevel int32

const (
	LOG_LEVEL_DEBUGGING	LoggingVerbosityLevel = 4
)

func SetLoggingVerbosity(level LoggingVerbosityLevel) {
	flag.Set("v", strconv.Itoa(int(level)))
	//flag.Parse()
}

func SetLoggingOutputFile(file string) {
	flag.Set("log_file", file)
	//flag.Parse()
}

type LoggingOutputType int32

const (
	LOG_OUTPUT_TYPE_CONSOLE		LoggingOutputType = iota
	LOG_OUTPUT_TYPE_FILE
	LOG_OUTPUT_TYPE_SYSLOG
)

func SetLoggingOutputType(t LoggingOutputType) error {
	switch t {
	case LOG_OUTPUT_TYPE_CONSOLE:
		flag.Set("logtostderr", "true")
	case LOG_OUTPUT_TYPE_FILE:
		flag.Set("logtostderr", "false")
	case LOG_OUTPUT_TYPE_SYSLOG:
		flag.Set("logtostderr", "false")
		logwriter, err := syslog.New(syslog.LOG_NOTICE, "libopenflow")
		if err != nil {
			return err
		}
		klog.SetOutput(logwriter)
	}
	//flag.Parse()
	return nil
}

func init() {
	klog.InitFlags(nil)
}
