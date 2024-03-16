package modex

import "fmt"

// Module created via:
//go mod init github.com/rootVIII/modex

// semver 1.2.3
// 1 is a major and possibly breaking change
// 2 is for backward compatible changes
// 3 is patch version and should not be for breaking changes to public API

// After pushing changes, use git tag to version module
// git tag v0.1.0
// git push origin v0.1.0

// In client codebase use go get to install to ~/go/pkg/mod/github.com
// go get github.com/rootVIII/modex@v0.1.0
// go get github.com/rootVIII/modex@v0.1.1

// Module version is visible in client codebase's go.mod

// Combine Uses v0.1.0
func Combine(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s", val)
}

// Combine2 Uses v0.1.1
func Combine2(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s - %[1]s", val)
}
