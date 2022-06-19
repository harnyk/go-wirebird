package main

import (
	"log"
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
	"github.com/harnyk/go-wirebird/internal/models"
)

type LoggerResult struct {
	Accepted bool `json:"accepted"`
}

type V1LoggerEventInput struct {
	Body models.LoggerEvent `json:"body"`
}

func main() {
	app := cli.NewRouter("Wirebird API", "0.0.1")
	app.DocsHandler(huma.SwaggerUIHandler(app.Router))

	app.Resource("/updates").
		Post(
			"post-logger-event-v1",
			"Save a logger event (backwards compatible)",
			responses.
				OK().
				Model(&LoggerResult{}),
		).
		Run(
			func(ctx huma.Context, input V1LoggerEventInput) {
				ctx.Header().Set("Content-Type", "application/json")
				log.Printf("Received event: %+v", input)
				ctx.WriteModel(http.StatusOK, &LoggerResult{Accepted: true})
			},
		)
	app.Run()

}
