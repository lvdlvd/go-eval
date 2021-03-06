package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/lvdlvd/go-eval"
)

func main() {
	in := bufio.NewReader(os.Stdin)
	for {
		if _, err := os.Stdout.WriteString("> "); err != nil {
			log.Fatalf("WriteString: %s", err)
		}
		line, err := in.ReadBytes('\n')
		if err == io.EOF {
			return
		}
		if err != nil {
			log.Fatalf("ReadBytes: %s", err)
		}

		e, err := eval.Parse(string(line))
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(e)
		fmt.Println(e.Evaluate(map[string]float64{"foo": 42}))
	}
}
