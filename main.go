package main

import (
	"github.com/cloudquery/cq-provider-sdk/serve"
)

func main() {
	serve.Serve(&serve.Options{
		Name:     "nil",
		Provider: nil,
	})
}
