package models

type ProcessData struct {
	PID        int    `json:"pid"`
	Title      string `json:"title,omitempty"`
	MainModule string `json:"mainModule,omitempty"`
}

type HeadersData []HeaderItem

type HeaderItem struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type RequestData struct {
	ID            string      `json:"id"`
	TimeStart     int64       `json:"timeStart"`
	TimeEnd       int64       `json:"timeEnd"`
	URL           string      `json:"url"`
	Method        string      `json:"method"`
	RemoteAddress string      `json:"remoteAddress,omitempty"`
	Headers       HeadersData `json:"headers"`
	Body          string      `json:"body,omitempty"`
}

type ResponseData struct {
	//Example comment
	StatusCode int         `json:"statusCode"`
	TimeStart  int64       `json:"timeStart"`
	TimeEnd    int64       `json:"timeEnd"`
	Headers    HeadersData `json:"headers"`
	Body       string      `json:"body,omitempty"`
}

type ErrorData struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message"`
	Stack   string `json:"stack,omitempty"`
}

type LoggerEvent struct {
	Request  RequestData  `json:"request"`
	Response ResponseData `json:"response,omitempty"`
	Error    ErrorData    `json:"error,omitempty"`
	Process  ProcessData  `json:"process,omitempty"`
}
