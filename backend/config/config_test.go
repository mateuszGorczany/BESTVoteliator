package config

import (
	"testing"

	common "github.com/mateuszGorczany/BESTVoteliator/utils"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func init() {
	common.LoadConfig()
}

func TestDBConfig(t *testing.T) {
	assert.Equal(t, viper.GetString("Database.Sheets.SheetID"), "1pl-MqjFebhLKDL51B1D9qXgG1J3Yr7dq7b7XaM63yJY")
	assert.Equal(t, viper.GetString("Database.Sheets.Credentials"), "bestvoteliator-361500-7f0eeec22d74.json")
	common.LoadSpecificConfig(viper.GetString("Database.Sheets.Credentials"))
	assert.Equal(t, viper.GetString("type"), "service_account")
}
