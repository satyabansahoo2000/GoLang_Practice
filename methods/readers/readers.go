package readers

import (
	"fmt"
	"io"
	"strings"
)

func PrintReader() {
	r := strings.NewReader("Hello World!!!")

	batch := make([]byte, 8)
	for {
		n, err := r.Read(batch)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, batch)
		fmt.Printf("b[:n] = %v\n", batch[:n])
		if err == io.EOF {
			break
		}
	}
}