package file

import (
	"bufio"
	"io"
	"iter"
	"os"
)

func Lines(path string) iter.Seq2[[]byte, error] {
	return func(yield func([]byte, error) bool) {
		f, err := os.Open(path)
		if err != nil {
			yield(nil, err)
			return
		}
		defer f.Close()

		bufReader := bufio.NewReader(f)

		for {
			bytes, err := bufReader.ReadBytes('\n')
			if err != nil && err != io.EOF {
				yield(nil, err)
				return
			}

			if len(bytes) > 0 {
				if !yield(bytes, nil) {
					return
				}
			}

			if err == io.EOF {
				return
			}
		}
	}
}
