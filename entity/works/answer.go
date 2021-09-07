package answer

import (
	"errors"

	"./question"
	"./users/user"
)


type Answer struct {
	user *user.User
	question *question.Question
	answer string
	isRight bool
}

// Check whether answer is right or not
func (a Answer) isAnswerRight() (bool, error) {
	if len(a.question.rightAnswers) == 0 {
		return false, errors.New("There is no one right answer")
	}

	for _, rightAnswer := range a.question.rightAnswers {
		if a.answer == rightAnswer.answer {
			return true, nil
		}
	}
	return false, nil
}
