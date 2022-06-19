package main

import (
	"net/http"

	"github.com/danielgtaylor/huma"
	"github.com/danielgtaylor/huma/cli"
	"github.com/danielgtaylor/huma/responses"
)

type LoggerResult struct {
	Accepted bool `json:"accepted"`
}

func main() {
	app := cli.NewRouter("Wirebird API", "0.0.1")
	// app.DocsHandler(huma.SwaggerUIHandler(app.Router))

	app.Resource("/").Post("post-logger-event", "Save a logger event",
		responses.OK().Model(&LoggerResult{}),
	).Run(func(ctx huma.Context) {
		ctx.Header().Set("Content-Type", "application/json")
		ctx.WriteModel(http.StatusOK, &LoggerResult{Accepted: true})
	})

	app.Run()

}
