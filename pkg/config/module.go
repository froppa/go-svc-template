package config

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/fx"
)

// Module provides configuration to an Fx application.
var Module = fx.Module(
	"config",
	fx.Provide(NewConfig),
)

// Config holds service settings loaded from YAML or env vars.
type Config struct {
	Meta struct {
		Name string `mapstructure:"name"`
	} `mapstructure:"meta"`
	Server struct {
		HTTPPort int `mapstructure:"httpPort"`
		GRPCPort int `mapstructure:"grpcPort"`
	} `mapstructure:"server"`
}

// NewConfig loads the file given by --config (or CONFIG env), defaulting to "configs/config.yml".
func NewConfig() (*Config, error) {
	cfgFile := viper.GetString("config")
	if cfgFile == "" {
		cfgFile = "configs/config.yml"
	}
	viper.SetConfigFile(cfgFile)

	if err := viper.ReadInConfig(); err != nil {
		return nil, fmt.Errorf("read config %q: %w", cfgFile, err)
	}
	var cfg Config
	if err := viper.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}
	return &cfg, nil
}

// RegisterFlags defines `--config` on your Cobra root and binds it into Viper.
func RegisterFlags(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("config", "c", "configs/config.yml", "path to config file")
	viper.BindPFlag("config", cmd.PersistentFlags().Lookup("config"))
	viper.AutomaticEnv()
}
