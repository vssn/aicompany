package types

import (
    "github.com/pocketbase/pocketbase/core"
)

type Account struct {
	Id string
	Credits float64
}

func NewAccountFromRecord(r *core.Record) Account {
	return Account {
		Id: r.Id,
		Credits: r.GetFloat("credits"),
	}
}