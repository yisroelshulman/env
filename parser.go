package env

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"unicode"
)

const (
	singleQuote = '\''
	doubleQuote = '"'
)

// parse reads the content from the provided reader and returns a map with the key, value pairs
// contained in the content
func parse(r io.Reader) (map[string]string, error) {
	var buf bytes.Buffer
	_, err := io.Copy(&buf, r)
	if err != nil {
		return nil, err
	}

	return unmarshalBytes(buf.Bytes())
}

// unmarshalBytes creates and returns a map with the key value pairs from src
func unmarshalBytes(src []byte) (map[string]string, error) {
	out := make(map[string]string)
	err := parseBytes(src, out)

	return out, err
}

// parseBytes extracts the key value pairs from the src and puts them in the provided map
func parseBytes(src []byte, out map[string]string) error {
	src = bytes.Replace(src, []byte("\r\n"), []byte("\r"), -1)
	remBuf := src

	for {
		remBuf = getStatementStart(remBuf)
		if len(remBuf) == 0 {
			break
		}

		key, left, err := getKey(remBuf)
		if err != nil {
			return err
		}

		value, left, err := getValue(left)
		if err != nil {
			return err
		}

		out[key] = value
		remBuf = left
	}

	return nil
}

// getStatementStart extracts the sub array from the start of the next statement to the end of src
func getStatementStart(src []byte) []byte {
	pos := indexOfNonSpaceChar(src)
	if pos < 0 {
		return nil
	}
	return src[pos:]
}

func indexOfNonSpaceChar(src []byte) int {
	return bytes.IndexFunc(src, func(r rune) bool {
		return !unicode.IsSpace(r)
	})
}

// getKey splits a src array into a key and rest based by the first '=' then valildates the shape
// of the key.
func getKey(src []byte) (key string, remBuf []byte, err error) {
	src = bytes.TrimLeftFunc(src, isSpace)
	if len(src) == 0 {
		return "", nil, errors.New("string is empty")
	}

	offset := 0
	for i, char := range src {
		if char == '=' {
			key = string(src[:i])
			offset = i + 1
			break
		}
	}

	key = strings.TrimRightFunc(key, unicode.IsSpace)
	err = validateEnvVariableKey(key)
	if err != nil {
		return "", nil, err
	}

	remBuf = bytes.TrimLeftFunc(src[offset:], isSpace)

	return key, remBuf, nil
}

// validateEnvVariableKey validates the key to make sure it has the correct shape
// a valid key has the shape starts with a letter followed by 0 or more alphanumeric or underscore
// charachters
func validateEnvVariableKey(key string) error {
	if len(key) == 0 || !unicode.IsLetter(rune(key[0])) {
		return errors.New("invalid key: zero len")
	}

	for _, char := range key {
		if unicode.IsNumber(char) || unicode.IsLetter(char) || char == '_' {
			continue
		}
		return errors.New("invalid key")
	}
	return nil
}

// getValue parses for a valid substring in the src
// valid substrings are enclosed in either singlQuotes or doubleQuotes
func getValue(src []byte) (value string, remBuf []byte, err error) {
	if len(src) == 0 {
		return "", nil, errors.New("no value")
	}

	offset := 0
	switch src[0] {
	case singleQuote, doubleQuote:
		for i, char := range src[1:] {
			if (char == singleQuote && src[0] == singleQuote) || (char == doubleQuote && src[0] == doubleQuote) {
				offset = i + 1
				break
			}
		}
	default:
		break
	}

	if offset == 0 {
		return "", nil, errors.New("invalid value")
	}

	value = string(src[1:offset])
	remBuf = src[offset+1:]

	return value, remBuf, nil
}

// isSpace checks for all space runes execpt for new line ones
func isSpace(r rune) bool {
	switch r {
	case '\f', '\r', '\t', '\v', ' ', 0x85, 0xA0:
		return true
	}
	return false
}
