package main

import (
	"os"

	"github.com/code-raushan/static-file-server/cmd/server"
)

func main(){
	os.Exit(server.Main())
}