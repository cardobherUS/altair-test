package service

import (
	"altair-test/internal/definitions"
	"altair-test/pkg/utils"
	"fmt"
	"log"
	"strings"
)

type ServiceI interface {
	CreateTeams()
	ReturnIDs(users []definitions.User) (ids []string)
	GetTeamUserByID(id string, users []definitions.User) (user definitions.User)

	Exec() (err error)
}

type Service struct {
	Tools utils.ToolsI
}

func NewService(tools utils.ToolsI) ServiceI {
	s := new(Service)
	s.Tools = tools
	return s
}

//CreateTeams Adds each user to their corresponding team
func (s *Service) CreateTeams() {
	for _, us := range definitions.Users {
		switch us.Team {
		case definitions.RedT:
			definitions.RedTeam = append(definitions.RedTeam, us)
		case definitions.BlueT:
			definitions.BlueTeam = append(definitions.BlueTeam, us)
		}
	}
}

//ReturnIDs Returns all the IDs of the team given
func (s *Service) ReturnIDs(users []definitions.User) (ids []string) {
	for _, u := range users {
		ids = append(ids, u.Id)
	}
	return
}

//GetTeamUserByID Returns a user, if it exists in the given list based on the given ID
func (s *Service) GetTeamUserByID(id string, users []definitions.User) (user definitions.User) {
	for _, u := range users {
		if strings.EqualFold(id, u.Id) {
			user = u
			break
		}
	}
	return
}
func (s *Service) Exec() (err error) {
	var id string
	s.CreateTeams() //First we create the different teams

	//We start looking for red team users
	redIDs := s.ReturnIDs(definitions.RedTeam)
	log.Print("Write one of the following IDs for RED TEAM seeking -> ", redIDs)
	_, err = fmt.Scan(&id)
	if err != nil {
		return
	}

	redUser := s.GetTeamUserByID(id, definitions.RedTeam)
	if strings.EqualFold(redUser.Team, "") {
		log.Print("There is none on RED TEAM with the ID:", id)
	} else {
		word := strings.ToLower(redUser.Word)
		wordR := s.Tools.Reverse(word) //Here we check if the word of the user is palindrome
		if strings.EqualFold(word, wordR) {
			log.Printf("The user '%s %s' with ID '%s' has a palindrome word: %s", redUser.Name, redUser.Surname, redUser.Id, word)
		} else {
			log.Printf("The user '%s %s' with ID '%s' hasn't a palindrome word: %s", redUser.Name, redUser.Surname, redUser.Id, word)
		}
	}

	//Now we start looking for blue team users
	blueIDs := s.ReturnIDs(definitions.BlueTeam)
	log.Print("Write one of the following IDs for BLUE TEAM seeking -> ", blueIDs)
	_, err = fmt.Scan(&id)
	if err != nil {
		return
	}

	blueUser := s.GetTeamUserByID(id, definitions.BlueTeam)
	if strings.EqualFold(blueUser.Team, "") {
		log.Print("There is none on BLUE TEAM with the ID:", id)
	} else {
		count := s.Tools.CountVowels(strings.ToLower(blueUser.Word)) //Here we check if the word has at least 5 vowels
		if count >= 5 {
			log.Printf("The user '%s %s' with ID '%s' has a word with at least 5 vowels: %s", blueUser.Name, blueUser.Surname, blueUser.Id, blueUser.Word)
		} else {
			log.Printf("The user '%s %s' with ID '%s' hasn't a word with at least 5 vowels: %s", blueUser.Name, blueUser.Surname, blueUser.Id, blueUser.Word)
		}
	}

	return
}
