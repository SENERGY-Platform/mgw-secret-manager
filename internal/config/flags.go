package config

import (
	"flag"
)

type Flags struct {
	ConfPath *string
}

func NewFlags() *Flags {
	f := Flags{
		ConfPath: flag.String("config", "", "path to config JSON file"),
	}
	flag.Parse()
	return &f
}
