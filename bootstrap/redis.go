package bootstrap

type Cache struct {
	Redis Redis
}

func (p *Cache) GetRedis() *Redis {
	if p == nil {
		return nil
	}
	return &p.Redis
}

type Redis struct {
	Host     string
	Port     string
	Password string
	Db       int
}

func (p *Redis) GetHost() string {
	if p == nil {
		return ""
	}
	return p.Host
}

func (p *Redis) GetPort() string {
	if p == nil {
		return ""
	}
	return p.Port
}

func (p *Redis) GetPassword() string {
	if p == nil {
		return ""
	}
	return p.Password
}

func (p *Redis) GetDB() int {
	if p == nil {
		return 0
	}
	return p.Db
}
