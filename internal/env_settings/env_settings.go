package envsettings

import "os"

type EnvSettings struct {
	Port     string
	DbPath   string
	Password string
}

var Env *EnvSettings

func Init() {
	Env = &EnvSettings{
		Port:   ":" + os.Getenv("TODO_PORT"),
		DbPath: os.Getenv("TODO_DBFILE"),
		// Password: os.Getenv("TODO_PASSWORD"),
	}
}
