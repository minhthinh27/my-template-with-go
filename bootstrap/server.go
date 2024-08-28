package bootstrap

type Server struct {
	Env  Env
	Http Http
}

type Env struct {
	Mode string
}

type Http struct {
	Address string
	Timeout int
}
