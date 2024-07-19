package settings

import (
	"os"
)

type Settings struct {
	DefaultDirectory string
	ClearScreen      bool
	Shell            string
}

func Defaults() Settings {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		panic(err)
	}
	return Settings{
		DefaultDirectory: homeDir,
		ClearScreen:      true,
		Shell:            "bash",
	}
}

