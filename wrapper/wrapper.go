package wrapper

import "fmt"

func Wrap(s string) string {
	return fmt.Sprintf("{%s}", s)
}
