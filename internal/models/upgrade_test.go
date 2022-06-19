package models_test

import (
	"encoding/json"
	"testing"

	"github.com/harnyk/go-wirebird/internal/models"
	"github.com/harnyk/go-wirebird/internal/models/compat"
	"github.com/ysmood/got"
)

func Test_upgradeHeaders(t *testing.T) {
	jsonData := `{
		"Content-Type": "application/json",
		"Content-Length": 123,
		"X-Foo": [
			"bar",
			100,
			["hello","world"]
		]
	}`

	testData := compat.LoggerHeaders{}
	err := json.Unmarshal([]byte(jsonData), &testData)
	if err != nil {
		t.Fatal(err)
	}

	t.Run("Success", func(t *testing.T) {
		headers, err := models.UpgradeHeaders(testData)
		if err != nil {
			t.Error(err)
		}

		got.T(t).Eq(headers, models.HeadersData{
			{"Content-Length", "123"},
			{"Content-Type", "application/json"},
			{"X-Foo", "bar"},
			{"X-Foo", "100"},
			{"X-Foo", "hello"},
			{"X-Foo", "world"},
		})
	})

	t.Run("Error", func(t *testing.T) {
		testData2 := compat.LoggerHeaders{
			"Hello":        "World",
			"Content-Type": true,
		}
		headers, err := models.UpgradeHeaders(testData2)
		got.T(t).Eq(err.Error(), "unsupported type for header item: bool")
		got.T(t).Eq(len(headers), 0)
	})
}

func Test_upgrade(t *testing.T) {
	jsonData := `{
		"request": {
			"id": "123",
			"timeStart": 1514764800000,
			"url": "http://example.com",
			"body": "Hello World",
			"method": "GET",
			"remoteAddress": "127.0.0.1",
			"headers": {
				"Content-Type": "application/json",
				"Content-Length": 123,
				"X-Foo": [
					"bar",
					100,
					["hello","world"]
				]
			}
		},
		"response": {
			"timeStart": 1514764800000,
			"body": "Hello World",
			"headers": {
				"Content-Type": "application/json",
				"Content-Length": 123
			},
			"rawHeaders": [
				"Content-Type",
				"application/json",
				"Content-Length",
				"123"
			],
			"status": 200
		},
		"error": {
			"code": "ERR_UNKNOWN",
			"message": "Unknown error",
			"stack": "Error: Unknown error\n\tat Object.error (http://example.com)\n\tat Object.error (http://example.com)"
		},
		"processData": {
			"pid": 123,
			"title": "Example",
			"mainModule": "example.exe"
		}
	}`

	testData := &compat.SerializedLoggerEvent{}
	err := json.Unmarshal([]byte(jsonData), testData)
	if err != nil {
		t.Fatal(err)
	}

	event, err := models.Upgrade(testData)
	if err != nil {
		t.Error(err)
	}

	got.T(t).Eq(event, &models.LoggerEvent{
		Request: models.RequestData{
			ID:            "123",
			TimeStart:     1514764800000,
			URL:           "http://example.com",
			Body:          "Hello World",
			Method:        "GET",
			RemoteAddress: "127.0.0.1",
			Headers: models.HeadersData{
				{"Content-Length", "123"},
				{"Content-Type", "application/json"},
				{"X-Foo", "bar"},
				{"X-Foo", "100"},
				{"X-Foo", "hello"},
				{"X-Foo", "world"},
			},
		},
		Response: &models.ResponseData{
			TimeStart: 1514764800000,
			Body:      "Hello World",
			Headers: models.HeadersData{
				{"Content-Length", "123"},
				{"Content-Type", "application/json"},
			},
			StatusCode: 200,
		},
		Error: &models.ErrorData{
			Code:    "ERR_UNKNOWN",
			Message: "Unknown error",
			Stack:   "Error: Unknown error\n\tat Object.error (http://example.com)\n\tat Object.error (http://example.com)",
		},
		Process: &models.ProcessData{
			PID:        123,
			Title:      "Example",
			MainModule: "example.exe",
		},
	})
}
