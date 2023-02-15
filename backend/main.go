// main.go
package main

import (
	"log"

	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/core"
)

func main() {
	app := pocketbase.New()

	app.OnRecordBeforeCreateRequest().Add(func(e *core.RecordCreateEvent) error {
		// overwrite the newly submitted "posts" record status to pending
		if e.Record.Collection().Name == "posts" {
			e.Record.Set("status", "pending1235")
		}
		return nil
	})

	if err := app.Start(); err != nil {
		log.Fatal(err)
	}

}
