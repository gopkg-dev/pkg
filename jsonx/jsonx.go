package jsonx

import (
	"bytes"

	"github.com/gopkg-dev/pkg/jsonx/jsontime"

	"github.com/tidwall/gjson"
)

// 定义JSON操作
var (
	json                 = jsontime.ConfigWithCustomTimeFormat
	Marshal              = json.Marshal
	Unmarshal            = json.Unmarshal
	MarshalIndent        = json.MarshalIndent
	NewDecoder           = json.NewDecoder
	NewEncoder           = json.NewEncoder
	SetDefaultTimeFormat = jsontime.SetDefaultTimeFormat
)

// Valid reports whether data is a valid JSON encoding.
func Valid(data []byte) bool {
	return json.Valid(data)
}

// GetStringFromJson get the string value from json path
func GetStringFromJson(json, path string) string {
	return gjson.Get(json, path).String()
}

// MarshalToString JSON编码为字符串
func MarshalToString(v interface{}) string {
	s, err := json.MarshalToString(v)
	if err != nil {
		return ""
	}
	return s
}

func MarshalToBytes(v interface{}) []byte {
	s, err := json.Marshal(v)
	if err != nil {
		return []byte{}
	}
	return s
}

// MarshalIndentToString JSON编码为格式化字符串
func MarshalIndentToString(v interface{}) string {
	bf := bytes.NewBuffer([]byte{})
	encoder := NewEncoder(bf)
	encoder.SetEscapeHTML(false)
	encoder.SetIndent("", "\t")
	_ = encoder.Encode(v)
	return bf.String()
}
