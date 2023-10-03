package configs

import (
	"sync"

	"github.com/go-playground/validator/v10"
)

var (
	once     sync.Once
	validate *validator.Validate
)

func GetValidator() *validator.Validate {
	once.Do(func() {
		validate = validator.New()
		// Register custom validator
	})

	return validate
}
