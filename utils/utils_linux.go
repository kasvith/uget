// +build !windows,!darwin

package utils

import (
	"os/user"
	path "path/filepath"
)

func AppData() string {
	user, _ := user.Current()
	return path.Join(user.HomeDir, ".config")
}
