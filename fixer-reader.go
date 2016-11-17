package kiwitaxi

import (
	"bufio"
	"bytes"
	"io"
	"strings"
)

type FixerReader struct {
	rBuf     *bufio.Reader
	wBuf     *bytes.Buffer
	replacer *strings.Replacer
	columns  int
}

func NewFixerReader(r io.Reader) *FixerReader {
	return &FixerReader{
		rBuf:     bufio.NewReader(r),
		wBuf:     bytes.NewBuffer([]byte{}),
		replacer: strings.NewReplacer("\\N", ""),
		columns:  -1,
	}
}

func (f *FixerReader) Read(p []byte) (n int, err error) {
	if f.wBuf.Len() == 0 {
		data, _, err := f.rBuf.ReadLine()
		if err != nil {
			return n, err
		}
		if len(data) == 0 {
			return 0, io.EOF
		}
		// skip bad line
		if !f.checkColumn(data) {
			return 0, nil
		}

		f.wBuf.Write(f.fixLine(data))
	}

	n, _ = f.wBuf.Read(p)
	return
}

func (f *FixerReader) checkColumn(line []byte) bool {
	countColumns := strings.Count(string(line), "\t")
	if f.columns == -1 {
		f.columns = countColumns
	}

	if f.columns != countColumns {
		return false
	}
	return true
}

func (f *FixerReader) fixLine(line []byte) []byte {
	return []byte(f.replacer.Replace(string(line)) + "\n")
}
