package helpers

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/shared"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

func InitializeLogDir() {
	logDir := Config.LogDir
	_ = os.Mkdir(logDir, os.ModePerm)
	f, err := os.OpenFile(logDir+Config.LogFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file:%v", err)
	}
	log.SetFlags(0)
	log.SetOutput(f)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func LogRequest(ctx *gin.Context) {
	blw := &bodyLogWriter{
		body:           bytes.NewBufferString(""),
		ResponseWriter: ctx.Writer,
	}

	ctx.Writer = blw
	ctx.Next()
	statusCode := ctx.Writer.Status()
	response := shared.NoErrorsFound
	level := "INFO"
	if statusCode >= 400 {
		response = blw.body.String()
		level = "ERROR"
	}
	data, err := json.Marshal(&LogStruct{
		Method:          ctx.Request.Method,
		Level:           level,
		StatusCode:      strconv.Itoa(statusCode),
		Path:            ctx.Request.URL.String(),
		UserAgent:       ctx.Request.Header.Get("User-Agent"),
		RemoteIP:        ctx.ClientIP(),
		ResponseTime:    time.Since(time.Now()).String(),
		Message:         http.StatusText(statusCode) + ":" + response,
		Version:         "1",
		CorrelationId:   uuid.New().String(),
		AppName:         Config.AppName,
		ApplicationHost: ctx.Request.Host,
		LoggerName:      "",
		TimeStamp:       time.Now().Format(time.RFC3339),
	})

	if err != nil {
		log.Fatal(err)
	}

	log.Printf("%s\n", data)
	ctx.Next()
}

func LogEvent(level string, message interface{}) {
	data, err := json.Marshal(struct {
		TimeStamp string      `json:"time_stamp"`
		Level     string      `json:"level"`
		Message   interface{} `json:"message"`
	}{
		TimeStamp: time.Now().Format(time.RFC3339),
		Level:     level,
		Message:   message,
	})

	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%s\n", data)
}
