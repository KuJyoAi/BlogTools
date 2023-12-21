package main

import "blog-tools/cmd"

func main() {
	err := cmd.Execute()
	if err != nil {
		panic(err)
	}
}
