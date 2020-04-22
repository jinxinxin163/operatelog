package main

import (
	"fmt"
	"github.com/jinxinxin163/operatelog/pkg/logger"
	"os"
	"time"
)

const (
	OPERATORLOGDIR = "/var/log/k8soperator"
	FILENUMBER = 4
	FILESIZE = 10
)
func InitLog() bool {
	var fileName string
	podName := os.Getenv("POD_NAME")
	namespace := os.Getenv("POD_NAMESPACE")
	if podName != "" && namespace != "" {
		fileName = namespace + "_" + podName + ".log"
	} else {
		return false
	}
	logger.SetConsole(true)
	logger.SetRollingFile(OPERATORLOGDIR, fileName, FILENUMBER, FILESIZE, logger.MB)
	// set logger.DEBUG, the Debug，Info，Warn, Error ,Fatal will print
	logger.SetLevel(logger.DEBUG)
	return true
}

func main() {
	ret :=  InitLog()
	if ret != true {
		fmt.Printf("can't get podName and namespace\n")
		return
	}
	var operateLog string
	operateLog = "nil sso  create user jin.xin admin  subid nil nil nil nil nil [create a user]"

	// Debug < Info < Warn < Error < Fatal
	// debug
	logger.Debug(operateLog)
	time.Sleep(1 * time.Second)
	// info
	logger.Info(operateLog)
	time.Sleep(1 * time.Second)
	// warn
	logger.Warn(operateLog)
	time.Sleep(1 * time.Second)
	// error
	logger.Error(operateLog)
	time.Sleep(1 * time.Second)
	// fatal
	logger.Fatal(operateLog)
	time.Sleep(1 * time.Second)
}
