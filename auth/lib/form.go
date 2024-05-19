package lib

import (
	"net/http"
)

// GetFormValue returns a function that gets the value of a form field
func GetFormValue(r *http.Request) func(key string) string {
	return func(key string) string {
		if val, ok := r.Form[key]; ok {
			return val[0]
		}
		return ""
	}
}

// AllFieldsValid checks if all fields are valid
func AllFieldsValid(fields ...string) bool {
	for _, field := range fields {
		if field == "" {
			return false
		}
	}
	return true
}
