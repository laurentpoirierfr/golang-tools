package log

import (
	"os"
	"time"

	"github.com/laurentpoirierfr/golang-tools/config"
	"github.com/sirupsen/logrus"
	"github.com/vladoatanasov/logrus_amqp"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
	log.Level = logrus.DebugLevel
	log.Formatter = &logrus.JSONFormatter{
		FieldMap: logrus.FieldMap{
			logrus.FieldKeyMsg:   "message",
			logrus.FieldKeyLevel: "level",
			logrus.FieldKeyTime:  "timestamp",
		},
		TimestampFormat: time.RFC3339Nano,
	}

	log.Out = os.Stdout

	addAmqpHook()

}

func addAmqpHook() {
	if config.GetStringValue("logger.enabled") == "true" {
		amqpLogServerName := config.GetStringValue("logger.amqp.server")
		if len(amqpLogServerName) > 0 {
			amqpLogPort := config.GetStringValue("logger.amqp.port")
			amqpLogUsername := config.GetStringValue("logger.amqp.username")
			amqpLogPassword := config.GetStringValue("logger.amqp.password")
			amqpLogExchange := config.GetStringValue("logger.amqp.exchange")
			amqpLogRoute := config.GetStringValue("logger.amqp.route")
			hook := logrus_amqp.NewAMQPHook(amqpLogServerName+":"+amqpLogPort, amqpLogUsername, amqpLogPassword, amqpLogExchange, amqpLogRoute)
			log.AddHook(hook)
		}
	}
}

// Debug :
func Debug(trace string) {
	log.WithFields(getServiceInformations()).Debug(trace)
}

// Info :
func Info(trace string) {
	log.WithFields(getServiceInformations()).Info(trace)
}

// Warning :
func Warning(trace string) {
	log.WithFields(getServiceInformations()).Warning(trace)
}

// Error :
func Error(trace string) {
	log.WithFields(getServiceInformations()).Error(trace)
}

func getServiceInformations() map[string]interface{} {
	infos := map[string]interface{}{
		"service": config.GetStringValue("application.name"),
		"version": config.GetStringValue("application.version"),
	}
	return infos
}
