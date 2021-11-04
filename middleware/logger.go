package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	retalog "github.com/lestrrat-go/file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"math"
	"os"
	"time"
)

func Logger() gin.HandlerFunc {
	filePath := "log/log"
	// 软连接的路径 在windows下需要管理员权限启动 linux下给了文件权限
	linkName := "latest_log.log"
	scr, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE, 0075)
	if err != nil {
		fmt.Println("err", err)
	}
	logger := logrus.New()
	logger.Out = scr
	logger.SetLevel(logrus.DebugLevel)
	logWriter, _ := retalog.New(
		filePath+"%Y%m%d.log",
		retalog.WithMaxAge(12*30*24*time.Hour),
		//24小时分割一次
		retalog.WithRotationTime(24*time.Hour),
		retalog.WithLinkName(linkName),
	)
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.DebugLevel: logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		logrus.PanicLevel: logWriter,
	}
	Hook := lfshook.NewHook(writeMap, &logrus.TextFormatter{
		TimestampFormat: "2006-01-02 15:04:05",
	})
	logger.AddHook(Hook)
	return func(context *gin.Context) {
		data, err := context.GetRawData()
		if err != nil {
			fmt.Println(err.Error())
		}
		context.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))
		startTime := time.Now()
		// 洋葱模型
		context.Next()
		// 从开始到结束的时间
		stopTime := time.Since(startTime)
		spendTime := fmt.Sprintf("%d ms", int(math.Ceil(float64(stopTime.Nanoseconds())/1000000.0)))
		//hostName, err := os.Hostname()
		//if err != nil {
		//	hostName = "unknown"
		//}
		statusCode := context.Writer.Status()
		//clientIp := context.ClientIP()
		//userAgent := context.Request.UserAgent()
		//dataSize := context.Writer.Size()
		//fmt.Println(string(body1))
		//if dataSize < 0 {
		//	dataSize = 0
		//}
		method := context.Request.Method
		path := context.Request.RequestURI
		entry := logger.WithFields(logrus.Fields{
			//"HostName":  hostName,
			"status":    statusCode,
			"SpendTime": spendTime,
			//"Ip":        clientIp,
			"Method": method,
			"Path":   path,
			"Body":   string(data),
			//"DataSize":  dataSize,
			//"Agent":     userAgent,
		})
		//记录系统内部的错误
		if len(context.Errors) > 0 {
			entry.Error(context.Errors.ByType(gin.ErrorTypePrivate).String())
		}
		if statusCode >= 500 {
			entry.Error()
		} else if statusCode >= 400 {
			entry.Warn()
		} else {
			entry.Info()
		}
	}
}
