package player

type Player struct {
	Name string
}

func (u Player) GetName() string {
	return u.Name
}
