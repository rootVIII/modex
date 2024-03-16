package modex

import "fmt"

// This module was created via:
//go mod init github.com/rootVIII/modex

// semver 1.2.3
// 1 is a major and possibly breaking change
// 2 is for backward compatible changes
// 3 is patch version and should not be for breaking changes to public API

// After pushing changes, use git tag to version module
// git tag v0.1.0
// git push origin v0.1.0

// If new changes are made, add/commit/push and then tag/push to update versions

// In client codebase use go get to install to ~/go/pkg/mod/github.com
// go get github.com/rootVIII/modex@v0.1.0
// go get github.com/rootVIII/modex@v0.1.1
// Module version is now visible in client codebase's go.mod

// To import this module into client codebase via:
// import "github.com/rootVIII/modex"

// Run "go get" with version details if upgrading to new version (go.mod will update)
// If the module is no longer used, remove unused modules from "go.mod" via:
// go mod tidy

// Combine Uses v0.1.0
func Combine(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s", val)
}

// Combine2 Uses v0.1.1
func Combine2(val string) string {
	return fmt.Sprintf("%[1]s - %[1]s - %[1]s", val)
}
