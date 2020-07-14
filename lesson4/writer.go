package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

type Capper struct {
	wtr io.Writer
}

func (c *Capper) Write(p []byte) (n int, err error) {
	out := make([]byte, len(p))
	for i := 0; i < len(p); i++ {
		c := string(p[i])
		out[i] = byte(strings.ToUpper(c)[0])
	}

	return c.wtr.Write(out)
}

func main() {
	c := &Capper{wtr: os.Stdout}
	fmt.Fprintf(c, "Hello there")
}
