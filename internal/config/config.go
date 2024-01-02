// package config
package config

import (
	"net/url"
	"os"
	"sync"

	"github.com/avocagrow/plant-up/internal/errs"
)

var (
	env                = os.Getenv("RUN_MODE")
	secretStoreBaseURL = "https://api.doppler.com/v3/configs/config/secrets"
	secretStoreToken   = os.Getenv("SECRET_STORE_TOKEN")
	environment        = os.Getenv("RUN_MODE")
	defaultSecrets     = map[string]string{
		"ROACH_CONN": "postgresql://root@roach1:26257/defaultdb?sslmode=disable",
	}
)

type Configurer interface {
	Configure()
	GetSecrets() error
	Secrets() map[string]string
	ProjectName() string
}

type Config struct {
	projectName string
	environment string
	secrets     map[string]string
	serverURL   *url.URL
	l           sync.RWMutex
}

func NewConfig() (*Config, error) {
	return nil, errs.ErrNotImplemented
}

func (c *Config) Secrets() map[string]interface{} {
	s := make(map[string]interface{})
	return s
}

// ProjectName returns the name of the service
func (c *Config) ConfigServiceName() string {
	return c.projectName
}

func (c *Config) GetSecrets() error {
	return errs.ErrNotImplemented
}
