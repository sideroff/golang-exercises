package main

import (
	"bytes"
	"fmt"
)

// interfaces should be <verb>er (writer, reader, walker...)

// Writer interface
type Writer interface {
	Write([]byte) (int, error)
}

// ConsoleWriter asd
type ConsoleWriter struct {
}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}

//

type Writer2 interface {
	Write([]byte) (int, error)
}

// Closer test
type Closer interface {
	Close() error
}

// WriterCloser test
type WriterCloser interface {
	Writer2
	Closer
}

// BufferedWriterCloser test
type BufferedWriterCloser struct {
	buffer *bytes.Buffer
}

func (bwc *BufferedWriterCloser) Write(data []byte) (int, error) {
	n, err := bwc.buffer.Write(data)

	if err != nil {
		return 0, err
	}

	v := make([]byte, 8)
	for bwc.buffer.Len() > 8 {
		_, err := bwc.buffer.Read(v)
		if err != nil {
			return 0, err
		}
		_, err = fmt.Println(string(v))
		if err != nil {
			return 0, err
		}
	}
	return n, nil
}

// Close test
func (bwc *BufferedWriterCloser) Close() error {
	for bwc.buffer.Len() > 0 {
		data := bwc.buffer.Next(8)
		_, err := fmt.Println(string(data))
		if err != nil {
			return err
		}
	}
	return nil
}

// NewBufferedWriterCloser test
func NewBufferedWriterCloser() *BufferedWriterCloser {
	return &BufferedWriterCloser{
		buffer: bytes.NewBuffer([]byte{}),
	}
}

func main() {
	var w Writer = ConsoleWriter{}
	w.Write([]byte("Hello Go!"))

	//

	var wc WriterCloser = NewBufferedWriterCloser()
	wc.Write(([]byte("Hello Youtube listeners this is a test")))
	wc.Close()

	bwc, ok := wc.(*BufferedWriterCloser)
	fmt.Println(bwc, ok)
}
