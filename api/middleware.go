package api

import(
	"fmt"
    "github.com/pocketbase/pocketbase/core"

)

func Authenticate(r *core.RequestEvent) error {
	fmt.Println("Middleware...")
	return r.Next()
}