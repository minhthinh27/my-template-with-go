package bootstrap

type Database struct {
	Main  Main
	Slave Slave
}

func (p *Database) GetMain() *Main {
	if p == nil {
		return nil
	}
	return &p.Main
}

type Main struct {
	Host       string
	Port       string
	DBName     string `mapstructure:"db_name"`
	Username   string
	Password   string
	MaxCon     int `mapstructure:"max_con"`
	MaxIdleCon int `mapstructure:"max_idle_con"`
}

func (p *Main) GetHost() string {
	if p == nil {
		return ""
	}
	return p.Host
}

func (p *Main) GetPort() string {
	if p == nil {
		return ""
	}
	return p.Port
}

func (p *Main) GetName() string {
	if p == nil {
		return ""
	}
	return p.DBName
}

func (p *Main) GetUserName() string {
	if p == nil {
		return ""
	}
	return p.Username
}

func (p *Main) GetPassword() string {
	if p == nil {
		return ""
	}
	return p.Password
}

func (p *Main) GetMaxCon() int {
	if p == nil {
		return 0
	}
	return p.MaxCon
}

func (p *Main) GetMaxIdleCon() int {
	if p == nil {
		return 0
	}
	return p.MaxIdleCon
}

func (p *Database) GetSlave() *Slave {
	if p == nil {
		return nil
	}
	return &p.Slave
}

type Slave struct {
	Host       string
	Port       string
	DBName     string `mapstructure:"db_name"`
	Username   string
	Password   string
	MaxCon     int `mapstructure:"max_con"`
	MaxIdleCon int `mapstructure:"max_idle_con"`
}

func (p *Slave) GetHost() string {
	if p == nil {
		return ""
	}
	return p.Host
}

func (p *Slave) GetPort() string {
	if p == nil {
		return ""
	}
	return p.Port
}

func (p *Slave) GetName() string {
	if p == nil {
		return ""
	}
	return p.DBName
}

func (p *Slave) GetUserName() string {
	if p == nil {
		return ""
	}
	return p.Username
}

func (p *Slave) GetPassword() string {
	if p == nil {
		return ""
	}
	return p.Password
}

func (p *Slave) GetMaxCon() int {
	if p == nil {
		return 0
	}
	return p.MaxCon
}

func (p *Slave) GetMaxIdleCon() int {
	if p == nil {
		return 0
	}
	return p.MaxIdleCon
}
