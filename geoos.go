// Package geoos It is a basic package that defines debugging information.
package geoos

import (
	"os"
	"strings"
)

// GeoosTestTag Decide whether to perform test control
const GeoosTestTag = true

// const GeoosTestTag = false

// EnvPath return work path
func EnvPath() string {
	env, _ := os.Getwd()
	env = env[0 : strings.Index(env, "geoos")+6]
	return env
}
