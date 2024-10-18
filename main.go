package main

import (
	"beer/util"

	"beer/cmd"
)

func main() {

	util.LoadDotEnv()

	cmd.Execute()

}
