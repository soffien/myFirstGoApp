package resources

type Person struct {
	name    string
	address string
	age     int
}

func NewPerson(n string, add string, ag int) Person {
	return Person{name: n, address: add, age: ag}
}

func (p *Person) ChangeName(newName string) error {
	p.name = newName
	return nil
}

func (p *Person) AddAge(numberToAdd int) int {
	p.age = p.age + numberToAdd
	return p.age
}

func (p *Person) GetName() string {
	return p.name
}
func (p *Person) GetAge() int {
	return p.age
}
