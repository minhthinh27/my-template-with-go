package bootstrap

type Server struct {
	Env  Env
	Http Http
}

func (p *Server) GetEnv() *Env {
	if p == nil {
		return nil
	}
	return &p.Env
}

func (p *Server) GetHttp() *Http {
	if p == nil {
		return nil
	}
	return &p.Http
}

type Env struct {
	Mode string
}

func (p *Env) GetMode() string {
	if p == nil {
		return ""
	}
	return p.Mode
}

type Http struct {
	Address string
	Timeout int
}

func (p *Http) GetAddress() string {
	if p == nil {
		return ""
	}
	return p.Address
}

func (p *Http) GetTimeout() int {
	if p == nil {
		return 0
	}
	return p.Timeout
}
