package main

import (
	"bytes"
	"encoding/json"
	"fmt"
)

func main() {
	m := map[string]interface{}{
		"abc":  "abc",
		"xyz":  "xyz",
		"123":  123,
		"bool": true,
		"html": "<html><body>",
		"int":  456,
	}

	jsonData, err := encodeJSON(m)
	fmt.Println("jsonData = ", string(jsonData), "err = ", err)
}

func encodeJSON(v interface{}) ([]byte, error) {
	var buf bytes.Buffer
	encoder := json.NewEncoder(&buf)
	encoder.SetEscapeHTML(false)
	if err := encoder.Encode(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}
