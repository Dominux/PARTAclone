package exercise

import (
	"./question"
	"./users/user"
)


type Exercise struct {
	authors []*user.User
	questions []*question.Question
}
