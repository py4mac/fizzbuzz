// Package pflagx provide utils functions to manage input flags
package pflagx

import "os"

// LookupEnvOrString adds the capability to get env instead of param flag
func LookupEnvOrString(key string, defaultVal string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}

	return defaultVal
}
