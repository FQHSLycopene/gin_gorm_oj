package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"os/exec"
)

func main() {
	//go run code-user/main.go
	cmd := exec.Command("go", "run", "code-user/main.go")
	var out, err bytes.Buffer
	cmd.Stderr = &err
	cmd.Stdout = &out
	pipe, err2 := cmd.StdinPipe()
	if err2 != nil {
		log.Fatalln(err2)
	}
	io.WriteString(pipe, "23 11\n")
	if err2 := cmd.Run(); err2 != nil {
		log.Fatalln(err2, err.String())
	}
	fmt.Println(out.String())
	println(out.String() == "34\n")
}
