package app

import (
	"rules-explorer/internal/core/types"
)

type Config struct {
	Theme        types.Theme
	InitialFocus types.Focus
}

func NewConfig() *Config {
	return &Config{
		InitialFocus: types.FocusSearch,
	}
}