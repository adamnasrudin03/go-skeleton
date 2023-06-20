package usecase

import (
	"encoding/json"
	"fmt"

	"gitlab.spesolution.net/bni-merchant-management-system/go-sekeleton/entity"
	"gitlab.spesolution.net/bni-merchant-management-system/go-sekeleton/internal/helper"
	"gitlab.spesolution.net/bni-merchant-management-system/go-sekeleton/internal/queue"
)

type Log struct {
	queue queue.Queue
}

func NewLogUsecase(
	queue queue.Queue,
) *Log {
	return &Log{queue}
}

type LogUsecase interface {
	Log(status entity.LogType, message string, funcName string, err error, logFields map[string]string, processName string)
}

// Process writing log to file.
// Parameters :
//   - status: status of log (Check entity.LogType)
//   - message: message to descirbe the error (You can use it to indicate error dependencies/functions)
//   - funcName: source function that return error (Ex. walletUsecase.Create, etc.)
//   - err: error response from function
//   - logFields: additional data to track error (Ex. Indetifier ID, User ID, etc.)
//   - processName: name of process (optional, this can be use to track bug by process name) and make sure using Type Safety to write process name
func (w *Log) Log(status entity.LogType, message string, funcName string, err error, logFields map[string]string, processName string) {
	channel := entity.LogGeneralKey

	logData := entity.Log{
		Process:      processName,
		FuncName:     funcName,
		Message:      message,
		ErrorMessage: err.Error(),
		Status:       status,
		LogFields:    logFields,
	}

	log, _ := json.Marshal(logData)

	ts := fmt.Sprintf("[%s] %s:", helper.NowStrUTC(), status)
	o := fmt.Sprint(string(log))
	f := fmt.Sprintf("%s %s \r\n", ts, o)

	payload, _ := helper.Serialize(logData)
	errQueue := w.queue.Publish(queue.ProcessSyncLog, payload, 1)

	if errQueue != nil {
		helper.WriteLog(f, channel)
		return
	}
}
