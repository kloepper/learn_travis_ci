package main

import (
	"os"
	"strings"

	"github.com/kloepper/learn_travis_ci/message"
)

func main() {
	text := strings.Join(os.Args[1:len(os.Args)], " ")
	m := message.NewMessage(text)
	m.Send()
}
