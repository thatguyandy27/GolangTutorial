package main

import (
	"fmt"
	"io/ioutil"
	"strconv"

	"github.com/pkg/errors"
)

func killServer(pidFile string) error {
	contents, err := ioutil.ReadFile(pidFile)
	if err != nil {
		return errors.Wrap(err, "Can't open config file")
	}

	// str := strings.TrimSpace(string(contents))
	num, err := strconv.Atoi(string(contents[:len(contents)-1]))
	if err != nil {
		return errors.Wrap(err, "Contents are not an integer")
	}

	fmt.Printf("Killed %d\n", num)

	return nil
}

func main() {
	err := killServer("server.bad")
	if err != nil {
		fmt.Printf("Error %s \n", err)
	}
}
