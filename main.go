package main

import (
	"gogofly/cmd"
	_ "gogofly/docs"
)

// @title GoWeb
// @version 0.0.1
// @description Golang Web开发实现
func main() {
	defer cmd.Clean()
	cmd.Start()
}
