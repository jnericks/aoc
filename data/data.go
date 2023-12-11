package data

import (
	"bufio"
	"bytes"
)

func Strings(file []byte) []string {
	s := bufio.NewScanner(bytes.NewBuffer(file))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func Bytes(file []byte) [][]byte {
	s := bufio.NewScanner(bytes.NewBuffer(file))
	var out [][]byte
	for s.Scan() {
		out = append(out, s.Bytes())
	}
	return out
}
