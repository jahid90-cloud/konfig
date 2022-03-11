package env

import (
	"log"
	"os"
)

type Env interface {
	MustGet(string) string
	Get(string) string
	GetOrDefault(string, string) string
}

type env struct{}

func NewEnv() Env {
	return &env{}
}

// MustGet Retrieves the requested env variable. Exits if it is not found.
// NOTE: Do not call during requests; web server would crash for unavailable env vars!
func (e *env) MustGet(key string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		log.Fatal("missing env variable [", key, "]")
	}

	return value
}

// Get Retrieves the requested env variable. Returns empty string when it is missing.
func (e *env) Get(key string) string {
	return os.Getenv(key)
}

// GetOrDefault Retrieves the requested env variable. Returns the default value provided when it is missing.
func (e *env) GetOrDefault(key string, def string) string {
	value, ok := os.LookupEnv(key)

	if !ok {
		return def
	}

	return value
}
