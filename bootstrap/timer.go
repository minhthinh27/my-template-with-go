package bootstrap

type Timer struct {
	Zone string
}

func (p *Timer) GetZone() string {
	if p == nil {
		return ""
	}
	return p.Zone
}
