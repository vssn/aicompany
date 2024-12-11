package main

import (
    "log"
	"github.com/vssn/aicompany/api"

    "github.com/pocketbase/pocketbase"
    "github.com/pocketbase/pocketbase/core"
)

func main() {
    app := pocketbase.New()

    app.OnServe().BindFunc(func(se *core.ServeEvent) error {
		apiv1 := se.Router.Group("/api/v1")
		apiv1.BindFunc(api.Authenticate)
		apiv1.GET("/run", api.Run)

        return se.Next()
    })

    if err := app.Start(); err != nil {
        log.Fatal(err)
    }
}