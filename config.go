package log

import (
	"os"
	"path/filepath"
)

// --------------------------------------------------------------------------------------------------

const defaultOutputFolder = "/tmp/"

// --------------------------------------------------------------------------------------------------

// Config logger config
type Config struct {
	Level   string      `mapstructure:"level" json:"level" yaml:"level" toml:"level"`
	Context LineContext `mapstructure:"context" json:"context" yaml:"context" toml:"context"`
	Output  Output      `mapstructure:"output" json:"output" yaml:"output" toml:"output"`
}

// LineContext log context config
type LineContext struct {
	Format     string            `mapstructure:"format" json:"format" yaml:"format" toml:"format"`
	Fileds     map[string]string `mapstructure:"fileds" json:"fileds" yaml:"fileds" toml:"fileds"`
	ShowCaller bool              `mapstructure:"show_caller" json:"show_caller" yaml:"show_caller" toml:"show_caller"`
}

// Output log output config
type Output struct {
	File       bool   `mapstructure:"file" json:"file" yaml:"file" toml:"file"`
	FilePath   string `mapstructure:"file_path" json:"file_path" yaml:"file_path" toml:"file_path"`
	MaxSize    int    `mapstructure:"max_size" json:"max_size" yaml:"max_size" toml:"max_size"`
	MaxBackups int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups" toml:"max_backups"`
	MaxAge     int    `mapstructure:"max_age" json:"max_age" yaml:"max_age" toml:"max_age"`
	Compress   bool   `mapstructure:"compress" json:"compress" yaml:"compress" toml:"compress"`
	Console    bool   `mapstructure:"console" json:"console" yaml:"console" toml:"console"`
}

// Default make config default
func (conf *Config) Default() {
	conf.Output.File = false
	conf.Output.FilePath = getDefOutputPath()
	conf.Output.MaxSize = 10
	conf.Output.MaxBackups = 30
	conf.Output.MaxAge = 7
	conf.Output.Compress = false
	conf.Output.Console = true
	conf.Level = "INFO"
	conf.Context.Format = "console"
	conf.Context.Fileds = nil
	conf.Context.ShowCaller = true
}

// Fix validate and fix config
func (conf *Config) Fix() {
	if conf.Output.File {
		if conf.Output.FilePath == "" {
			conf.Output.FilePath = getDefOutputPath()
		}
		if conf.Output.MaxSize == 0 {
			conf.Output.MaxSize = 10
		}
		if conf.Output.MaxBackups == 0 {
			conf.Output.MaxSize = 30
		}
		if conf.Output.MaxAge == 0 {
			conf.Output.MaxAge = 7
		}
	} else {
		conf.Output.Console = true
	}
	if conf.Context.Format != "json" && conf.Context.Format != "console" {
		conf.Context.Format = "console"
	}
}

// getDefOutputPath get default output
func getDefOutputPath() string {
	return filepath.Join(defaultOutputFolder, filepath.Base(os.Args[0])+".log")
}

// --------------------------------------------------------------------------------------------------
