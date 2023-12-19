package vars

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

var logs = logrus.New()
var IsDebugEnabled = false

func InitLogger() {
	logs.SetOutput(os.Stdout)
	if IsDebugEnabled {
		logs.SetLevel(logrus.DebugLevel)
	} else {
		logs.SetLevel(logrus.InfoLevel)
	}

	logs.SetFormatter(&logrus.TextFormatter{
		ForceColors:      true,
		FullTimestamp:    true,
		TimestampFormat:  "2006/01/02 15:04:05",
		DisableSorting:   false,
		DisableTimestamp: true,
	})
}

func GetLogger() *logrus.Logger {
	return logs
}

func ToggleDebugLogging(c *gin.Context) {
	IsDebugEnabled = !IsDebugEnabled
	InitLogger()

	c.JSON(http.StatusOK, gin.H{
		"debug_logging": IsDebugEnabled,
	})
}
