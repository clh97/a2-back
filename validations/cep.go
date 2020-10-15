package validations

import (
	"strconv"
)

// ValidateCep is a simple CEP validator implementation
func ValidateCep(cep string) bool {
	// checks up if cep string consists of integers
	if _, err := strconv.Atoi(cep); err == nil {
		// checks up if cep length is 8 chars
		if len(cep) == 8 {
			return true
		}
	}
	return false
}
