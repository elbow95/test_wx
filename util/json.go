package util

import (
	"log"
	"strings"

	jsoniter "github.com/json-iterator/go"
)

func UnmarshalJson[T any](text string, t T) (T, error) {
	d := jsoniter.NewDecoder(strings.NewReader(text))
	d.UseNumber()
	err := d.Decode(t)
	return t, err
}

func UnmarshalJsonIgnoreError[T any](text string, t T) T {
	if len(text) == 0 {
		return t
	}
	t, err := UnmarshalJson(text, t)
	if err != nil {
		log.Printf("[UnmarshalJsonIgnoreError] unmarshal fail: %T, %+v", t, err)
	}
	return t
}

func MarshalJsonIgnoreError(obj interface{}) string {
	if obj == nil {
		return ""
	}
	s, err := jsoniter.MarshalToString(obj)
	if err != nil {
		log.Printf("[MarshalJsonIgnoreError] marshal fail: %+v", err)
	}
	return s
}
