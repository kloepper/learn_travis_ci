package message

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"io"
	"os"
)

type Message interface {
	Prepare(string)
	Send()
}

func NewMessage(content string) *localGzipMessage {
	m := &localGzipMessage{}
	m.Prepare(content)
	return m
}

type localGzipMessage struct {
	Content []byte
}

func (this *localGzipMessage) Prepare(content string) {
	out := &bytes.Buffer{}
	compressor := gzip.NewWriter(out)

	contentBuffer := bytes.NewBufferString(content)
	_, copyErr := io.Copy(compressor, contentBuffer)
	if copyErr != nil {
		panic(copyErr)
	}
	if err := compressor.Close(); err != nil {
		panic(err)
	}
	this.Content = out.Bytes()

	fmt.Printf("Prepared message of size %v (gzip size %v).\n", len(content), len(this.Content))
}

func (this *localGzipMessage) Send() {
	byteReader := bytes.NewReader(this.Content)
	gzipReader, gzipCreateErr := gzip.NewReader(byteReader)
	if gzipCreateErr != nil {
		panic(gzipCreateErr)
	}

	contentSize, inflateErr := io.Copy(os.Stdout, gzipReader)
	if inflateErr != nil {
		panic(inflateErr)
	}

	fmt.Println()
	fmt.Printf("Sent message of size %v.\n", contentSize)
}
