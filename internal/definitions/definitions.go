package definitions

const (
	BlueT = "azul"
	RedT  = "rojo"
)

var (
	BlueTeam []User
	RedTeam  []User
)

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Team    string `json:"team"`
	Age     int    `json:"age"`
	Word    string `json:"word"`
}

var Users = []User{
	{Id: "AAA", Name: "Javier", Surname: "Miguez", Team: RedT, Age: 44, Word: "riachuelo"},
	{Id: "AAB", Name: "Luis", Surname: "Alonso", Team: BlueT, Age: 32, Word: "murcielago"},
	{Id: "ABA", Name: "Cristina", Surname: "Perez", Team: BlueT, Age: 21, Word: "solos"},
	{Id: "ABB", Name: "Mariano", Surname: "Gomez", Team: RedT, Age: 18, Word: "reconocer"},
	{Id: "BAA", Name: "Alicia", Surname: "Domenech", Team: RedT, Age: 32, Word: "oso"},
	{Id: "BBA", Name: "Susana", Surname: "Gomez", Team: BlueT, Age: 19, Word: "elefante"},
}
