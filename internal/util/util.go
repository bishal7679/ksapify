package util

import "os"

func HomeDir() string {
	home := os.Getenv("HOME")
	if home == "" {
		// for windows os
		home = os.Getenv("USERPROFILE")
	}
	return home
}
