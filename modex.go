package modex

import "fmt"

func Combine(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s", val)
}
