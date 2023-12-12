package util

import (
	"bufio"
	"bytes"
)

func ReadStrings(file []byte) []string {
	s := bufio.NewScanner(bytes.NewBuffer(file))
	var out []string
	for s.Scan() {
		out = append(out, s.Text())
	}
	return out
}

func ReadBytes(file []byte) [][]byte {
	s := bufio.NewScanner(bytes.NewBuffer(file))
	var out [][]byte
	for s.Scan() {
		out = append(out, s.Bytes())
	}
	return out
}
