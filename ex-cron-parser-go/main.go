package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nunosilva800/cron-parser-go/parser"
)

func main() {
	cronLine := os.Args[1]
	p, err := parser.New(cronLine)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	strs := p.PrettyPrint()
	fmt.Println(strings.Join(strs, "\n"))
}
