package json

import "github.com/tidwall/gjson"

func ParseHello(s string) string {
	return gjson.Get(s, "hello").String()
}
