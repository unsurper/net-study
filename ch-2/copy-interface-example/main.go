package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

type CreateReadWriter interface {
	Create() FooReadWriter //构造初始化接口
}
type CreateRW struct {
}
type FooReadWriter interface {
	Read(p []byte) (n int, err error)
	Write(p []byte) (n int, err error)
}
type ReadWriter struct {
}

func (Use *CreateRW) Create() FooReadWriter {
	it := new(ReadWriter)
	return it
}

// Read reads data from stdin.
func (fooReader *ReadWriter) Read(b []byte) (int, error) {
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

// Write writes data to Stdout.
func (fooWriter *ReadWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}
func main() {
	// Instantiate reader and writer.
	var Create CreateRW
	reader := Create.Create()
	writer := Create.Create()
	if _, err := io.Copy(writer, reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
