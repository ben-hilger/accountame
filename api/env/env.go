package env

import "github.com/joho/godotenv"

type Environment interface {
	LoadEnvironment() error
}

type EnvironmentFile struct {
	File string
}

func (e EnvironmentFile) LoadEnvironment() error {
	return godotenv.Load(e.File)
}
