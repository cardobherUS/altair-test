package definitions

const (
	BLUE_T = "azul"
	RED_T  = "rojo"
)

var RedTeam, BlueTeam []User

type User struct {
	Id      string `json:"id"`
	Name    string `json:"name"`
	Surname string `json:"surname"`
	Team    string `json:"team"`
	Age     int    `json:"age"`
	Word    string `json:"word"`
}

var Users = []User{
	{Id: "AAA", Name: "Javier", Surname: "Miguez", Team: RED_T, Age: 44, Word: "riachuelo"},
	{Id: "AAB", Name: "Luis", Surname: "Alonso", Team: BLUE_T, Age: 32, Word: "murcielago"},
	{Id: "ABA", Name: "Cristina", Surname: "Perez", Team: BLUE_T, Age: 21, Word: "solos"},
	{Id: "ABB", Name: "Mariano", Surname: "Gomez", Team: RED_T, Age: 18, Word: "reconocer"},
	{Id: "BAA", Name: "Alicia", Surname: "Domenech", Team: RED_T, Age: 32, Word: "oso"},
	{Id: "BBA", Name: "Susana", Surname: "Gomez", Team: BLUE_T, Age: 19, Word: "elefante"},
}
