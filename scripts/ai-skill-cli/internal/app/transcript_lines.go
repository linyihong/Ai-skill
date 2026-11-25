package app

import (
	"bufio"
	"bytes"
	"io"
)

// transcriptLineReader iterates over JSONL transcript lines with no line
// length limit. It replaces bufio.Scanner in the transcript readers: a
// Scanner with a fixed buffer stops at the first line that does not fit
// (a large tool result or attachment can exceed 2 MB) and silently drops
// every later line, which hid receipts, reads and user turns from the hooks.
//
// Semantics otherwise match bufio.ScanLines: the trailing "\n" and a "\r"
// before it are removed, and a final line without a newline is returned.
type transcriptLineReader struct {
	r    *bufio.Reader
	line []byte
	long []byte // reused for lines longer than the bufio buffer
	done bool
}

func newTranscriptLineReader(r io.Reader) *transcriptLineReader {
	return &transcriptLineReader{r: bufio.NewReaderSize(r, 256*1024)}
}

// Scan advances to the next line; it returns false at EOF or on a read error.
// Like bufio.Scanner, Bytes is only valid until the next Scan.
func (t *transcriptLineReader) Scan() bool {
	if t.done {
		return false
	}
	line, err := t.r.ReadSlice('\n')
	if err == bufio.ErrBufferFull {
		t.long = append(t.long[:0], line...)
		for err == bufio.ErrBufferFull {
			line, err = t.r.ReadSlice('\n')
			t.long = append(t.long, line...)
		}
		line = t.long
	}
	if err != nil {
		t.done = true
		if len(line) == 0 {
			return false
		}
	}
	line = bytes.TrimSuffix(line, []byte("\n"))
	line = bytes.TrimSuffix(line, []byte("\r"))
	t.line = line
	return true
}

func (t *transcriptLineReader) Bytes() []byte { return t.line }

func (t *transcriptLineReader) Text() string { return string(t.line) }
