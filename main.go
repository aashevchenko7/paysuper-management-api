package main

import (
	"github.com/paysuper/paysuper-management-api/cmd/casbin"
	"github.com/paysuper/paysuper-management-api/cmd/http"
	"github.com/paysuper/paysuper-management-api/cmd/root"
)

// @title PaySuper payment solution service.
// @desc Swagger Specification for PaySuper Management API.
//
// @ver 1.0.0
// @server https://api.pay.super.com Production API
func main() {
	args := []string{
		"http", "-c", "configs/local.yaml", "-d",
	}
	root.ExecuteDefault(args, http.Cmd, casbin.Cmd)
}
