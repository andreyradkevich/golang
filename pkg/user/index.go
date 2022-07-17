package user

type User struct {
	Age  string
	Name string
}

type Getter interface {
	GetName() string
	GetAge() string
}

func (u User) GetAge() string {
	return u.Age
}

func (u User) GetName() string {
	return u.Name
}

func GetAgeNameString(getter Getter) string {
	return "Age: " + getter.GetAge() + "; Name :" + getter.GetName()
}
