package conf_test

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/wangjiandev/cloud-api/conf"
)

func TestLoadConfigFromToml(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/cloud.toml")
	if should.NoError(err) {
		should.Equal("demo", conf.C().App.Name)
	}
}

func TestLoadConfigFromTomlError(t *testing.T) {
	should := assert.New(t)
	err := conf.LoadConfigFromToml("../etc/error.toml")
	should.Error(err)
}

func TestLoadConfigFromEnv(t *testing.T) {
	os.Setenv("APP_HOST", "127.0.0.1")
	should := assert.New(t)
	err := conf.LoadConfigFromEnv()
	if should.NoError(err) {
		should.Equal("127.0.0.1", conf.C().App.Host)
	}
}
