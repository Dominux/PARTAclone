package entity

import (
	"errors"

	"github.com/SPARTAclone/SPARTA/internal/users/entity"
)


type Answer struct {
	User *entity.User
	Question *Question
	Answer string
	IsRight bool
}

// Check whether answer is right or not
func (a Answer) isAnswerRight() (bool, error) {
	if len(a.Question.RightAnswers) == 0 {
		return false, errors.New("There is no one right answer")
	}

	for _, rightAnswer := range a.Question.RightAnswers {
		if a.Answer == rightAnswer.Answer {
			return true, nil
		}
	}
	return false, nil
}
