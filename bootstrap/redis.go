package bootstrap

type Cache struct {
	Redis Redis
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Db       int
}
