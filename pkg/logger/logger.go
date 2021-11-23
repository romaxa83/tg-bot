package logger

// оберта для logrus
import (
	"fmt"

	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

// создаем свой интерфейс для логера
type Logger interface {
	// встраиваем интерфейс logrus
	logrus.FieldLogger
	// добавляем трейс
	TraceWrap(error) logrus.Fields
}

type Log struct {
	*logrus.Logger
}

func (l *Log) TraceWrap(err error) logrus.Fields {
	st := errors.WithStack(err)
	return logrus.Fields{
		"trace": fmt.Sprintf("%+v", st),
	}
}

// форматер - форматирует логи, hooks - отдать логи сторонему сервису
func NewLogger(fmter logrus.Formatter, hooks []logrus.Hook) Logger {
	log := &Log{
		logrus.New(),
	}
	log.Formatter = fmter
	for _, h := range hooks {
		log.AddHook(h)
	}

	return log
}
