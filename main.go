package main

import (
	"beer/cmd"
	"beer/util"
	"path/filepath"
)

func main() {

	envFile, _ := filepath.Abs(".env")
	util.LoadDotEnv(envFile)

	cmd.Execute()

}
