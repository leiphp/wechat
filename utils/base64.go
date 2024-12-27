package utils

import (
	"bytes"
	"encoding/base64"
)

// 字节转base64
func ByteToBase64(d []byte) string {
	b := bytes.Buffer{}
	w := base64.NewEncoder(base64.StdEncoding, &b)
	w.Write(d)
	defer w.Close()
	return b.String()
}
