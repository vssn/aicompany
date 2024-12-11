package api

import (
	"fmt"
	"github.com/vssn/aicompany/types"

    "github.com/pocketbase/pocketbase/core"
)


func Run(r *core.RequestEvent) error {
	records, err := r.App.FindAllRecords("accounts")
	if err != nil {
		return err
	}

	for _, record := range records {
		fmt.Println(types.NewAccountFromRecord(record))
	}

	return r.JSON(200, records)
}