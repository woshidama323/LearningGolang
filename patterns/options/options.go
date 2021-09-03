package options

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Country string `json:"country"`
	Gender  string `json:"gender"`
	Height  int    `json:"height"`
	Address string `json:"address"`
}

type per func(*Person)

func Country(country string) per {
	return func(person *Person) {
		person.Country = country
	}
}

func Gender(gender string) per {
	return func(person *Person) {
		person.Gender = gender
	}
}

func Address(address string) per {
	return func(person *Person) {
		person.Address = address
	}
}

func NewPerson(name string, settings ...per) *Person {
	person := &Person{
		Name:    name,
		Age:     -1,
		Gender:  "Male",
		Country: "china",
		Height:  0,
		Address: "unknow",
	}

	for _, setting := range settings {
		setting(person)
	}
	return person
}
