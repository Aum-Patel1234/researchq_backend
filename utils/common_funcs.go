package utils

import "os"

func IsDev() bool {
	return os.Getenv("GIN_MODE") == "debug"
}
