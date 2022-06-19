package models

import (
	"fmt"
	"sort"
	"strconv"

	"github.com/harnyk/go-wirebird/internal/models/compat"
)

func Upgrade(v1Model *compat.SerializedLoggerEvent) (*LoggerEvent, error) {
	e := &LoggerEvent{}
	request, err := upgradeRequest(v1Model.Request)
	if err != nil {
		return nil, err
	}
	e.Request = *request

	e.Response, err = upgradeResponse(v1Model.Response)
	if err != nil {
		return nil, err
	}
	e.Error, err = upgradeError(v1Model.Error)
	if err != nil {
		return nil, err
	}
	e.Process, err = upgradeProcess(v1Model.ProcessData)
	if err != nil {
		return nil, err
	}
	return e, nil
}

func upgradeRequest(v1Request compat.SerializedLoggerRequest) (*RequestData, error) {
	r := &RequestData{}
	r.ID = v1Request.ID
	r.TimeStart = v1Request.TimeStart
	r.URL = v1Request.URL
	r.Body = v1Request.Body
	r.Method = v1Request.Method
	r.RemoteAddress = v1Request.RemoteAddress

	headers, err := UpgradeHeaders(v1Request.Headers)
	if err != nil {
		return nil, err
	}
	r.Headers = headers
	return r, nil
}

func upgradeResponse(v1Response *compat.SerializedLoggerResponse) (*ResponseData, error) {
	if v1Response == nil {
		return nil, nil
	}
	r := &ResponseData{}
	r.TimeStart = v1Response.TimeStart
	r.Body = v1Response.Body
	r.StatusCode = v1Response.Status

	headers, err := UpgradeHeaders(v1Response.Headers)
	if err != nil {
		return nil, err
	}
	r.Headers = headers
	return r, nil
}

func upgradeError(v1Error *compat.LoggerError) (*ErrorData, error) {
	if v1Error == nil {
		return nil, nil
	}
	e := &ErrorData{}
	e.Code = v1Error.Code
	e.Message = v1Error.Message
	e.Stack = v1Error.Stack
	return e, nil
}

func upgradeProcess(v1ProcessData *compat.ProcessData) (*ProcessData, error) {
	if v1ProcessData == nil {
		return nil, nil
	}
	p := &ProcessData{}
	p.PID = v1ProcessData.PID
	p.Title = v1ProcessData.Title
	p.MainModule = v1ProcessData.MainModule
	return p, nil
}

func UpgradeHeaders(v1Headers compat.LoggerHeaders) (HeadersData, error) {
	headers := HeadersData{}

	keys := make([]string, 0, len(v1Headers))
	for k := range v1Headers {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	for _, k := range keys {
		chunk, err := upgradeHeaderItem(k, v1Headers[k])
		if err != nil {
			return nil, err
		}
		headers = append(headers, chunk...)
	}
	return headers, nil
}

func upgradeHeaderItem(k string, v interface{}) ([]HeaderItem, error) {
	switch vt := v.(type) {
	case string:
		return []HeaderItem{{k, vt}}, nil
	case float64:
		return []HeaderItem{{k, strconv.FormatInt(int64(vt), 10)}}, nil
	case []interface{}:
		items := make([]HeaderItem, 0, len(vt))
		for _, v := range vt {
			subItems, err := upgradeHeaderItem(k, v)
			if err != nil {
				return nil, err
			}
			items = append(items, subItems...)
		}
		return items, nil
	default:
		return nil, fmt.Errorf("unsupported type for header item: %T", v)
	}
}
