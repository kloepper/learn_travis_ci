package message

import (
	"testing"

	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type MessageSuite struct{}

var _ = Suite(&MessageSuite{})

func (*MessageSuite) TestSend(c *C) {
	m := localGzipMessage{}
	m.Prepare("Cat")
	m.Send()
}
