package util

import (
	"flag"
	"os"
	"strings"
)

// String returns the flag value specified by the name, or the equivalent environment variable if specified
func String(name string, value string, usage string) *string {
	envName := strings.ToUpper(name)
	if envVal, ok := os.LookupEnv(envName); ok {
		return flag.String(name, envVal, usage)
	}

	// fallback to flag default
	return flag.String(name, value, usage)
}

// ToEnvironmentVariableName Takes a flag name and returns the equivalent environment variable name. For example, "digitalocean-api-token" becomes "DIGITALOCEAN_API_TOKEN".
func ToEnvironmentVariableName(flagName string) string {
	return strings.ToUpper(strings.ReplaceAll(flagName, "-", "_"))
}
