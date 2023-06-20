package log

import (
	"fmt"
	"github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"log"
	"os"
	"time"
)

func init() {
	filePaths := "./log/info.log"
	writer, _ := rotatelogs.New(
		filePaths+".%Y%m%d%H%M",
		rotatelogs.WithLinkName(filePaths),
		rotatelogs.WithMaxAge(time.Duration(604800)*time.Second),
		rotatelogs.WithRotationTime(time.Duration(86400)*time.Second),
	)
	writerMap := lfshook.WriterMap{
		logrus.InfoLevel:  writer,
		logrus.FatalLevel: writer,
		logrus.DebugLevel: writer,
		logrus.WarnLevel:  writer,
		logrus.ErrorLevel: writer,
		logrus.PanicLevel: writer,
	}
	formatter := logrus.JSONFormatter{}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.SetOutput(os.Stdout)
	logrus.SetFormatter(&formatter)
	logrus.AddHook(lfshook.NewHook(writerMap, &formatter))
}

func Infoln(format string, v ...any) {
	logrus.Infoln(fmt.Sprintf(format, v))
}

func Warnln(format string, v ...any) {
	logrus.Warnln(fmt.Sprintf(format, v))

}

func Errorln(format string, v ...any) {
	logrus.Errorln(fmt.Sprintf(format, v))
}

func Debugln(format string, v ...any) {
	logrus.Debugln(fmt.Sprintf(format, v))
}

func Fatalln(format string, v ...any) {
	log.Fatalf(format, v...)
}
