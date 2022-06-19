package compat

type SerializedLoggerEvent struct {
	Request     SerializedLoggerRequest   `json:"request"`
	Response    *SerializedLoggerResponse `json:"response,omitempty"`
	Error       *LoggerError              `json:"error,omitempty"`
	ProcessData *ProcessData              `json:"processData,omitempty"`
}

type SerializedLoggerRequest struct {
	ID            string        `json:"id"`
	TimeStart     int64         `json:"timeStart"`
	URL           string        `json:"url"`
	Body          string        `json:"body,omitempty"`
	Headers       LoggerHeaders `json:"headers"`
	Method        string        `json:"method"`
	RemoteAddress string        `json:"remoteAddress,omitempty"`
}

type SerializedLoggerResponse struct {
	TimeStart  int64                    `json:"timeStart"`
	Body       string                   `json:"body,omitempty"`
	Headers    LoggerHeaders            `json:"headers"`
	RawHeaders LoggerResponseRawHeaders `json:"rawHeaders"`
	Status     int                      `json:"status"`
}

type LoggerError struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Stack   string `json:"stack"`
}

type ProcessData struct {
	PID        int64  `json:"pid"`
	Title      string `json:"title"`
	MainModule string `json:"mainModule"`
}

type LoggerHeaders map[string]interface{}
type LoggerResponseRawHeaders []string
