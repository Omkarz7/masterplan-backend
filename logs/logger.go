package logs

import (
	"masterplan-backend/models"
	"os"

	"github.com/natefinch/lumberjack"
	"github.com/sirupsen/logrus"
)

var HTTPlog = make(chan logrus.Fields, 100) //buffer for logs

func LogrusHTTP() {
	logPath, err := os.Getwd()
	if err != nil {
		panic(err.Error())
	}
	logPath = logPath + models.Config.LogPath + ".log"
	for {
		select {
		case value := <-HTTPlog:

			lumberjackLogger := &lumberjack.Logger{
				Filename:  logPath,
				MaxSize:   2,    // megabytes after which new file is created
				LocalTime: true, //use local time and not utc time for formaating timestamp
			}
			logrus.SetOutput(lumberjackLogger)
			if value["status"] == 200 {
				logrus.WithFields(value).Info() //write the data as info

			} else {
				logrus.WithFields(value).Error() //write the data as info
			}
			n := len(HTTPlog) //check if any data is left in buffer, write that too
			for i := 0; i < n; i++ {
				value = <-HTTPlog

				if value["status"] == 200 {
					logrus.WithFields(value).Info() //write the data as info

				} else {
					logrus.WithFields(value).Error() //write the data as info
				}
			}
			lumberjackLogger.Close()
		}
	}

}
