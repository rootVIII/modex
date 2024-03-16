package modex

import "fmt"

// semver 1.2.3
// 1 is a major and possibly breaking change
// 2 is for backward compatible changes
// 3 is patch version and should not be for breaking changes to public API

func Combine(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s", val)
}
