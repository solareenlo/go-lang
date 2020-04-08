package main

import (
	"io"
	"log"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (rot13 *rot13Reader) Read(buf []byte) (int, error) {
	n, e := rot13.r.Read(buf)
	for i := 0; i < n; i++ {
		if buf[i] >= 'A' && buf[i] <= 'M' {
			buf[i] += 13
		} else if buf[i] >= 'N' && buf[i] <= 'Z' {
			buf[i] -= 13
		} else if buf[i] >= 'a' && buf[i] <= 'm' {
			buf[i] += 13
		} else if buf[i] >= 'n' && buf[i] <= 'z' {
			buf[i] -= 13
		}
	}
	return n, e
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	// io.Copy(os.Stdout, &r)
	if _, err := io.Copy(os.Stdout, &r); err != nil {
		log.Fatal(err)
	}
}
