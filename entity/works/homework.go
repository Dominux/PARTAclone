package homework

import (
	"./answer"
	"./exercise"
	"./points"
	"./users/user"
)


type Homework struct {
	student *user.User
	exercise *exercise.Exercise
	answers []*answer.Answer
	points points.Points
	isGraded bool
}

// Auto check homework's questions by comparing 
// it's right answers with given by student 
func (h *Homework) autoCheck() {
	for _, answer := range h.answers {

		isAnswerRight, err := answer.question.isAnswerRight(answer)

		if err != nil {
			continue
		}

		if isAnswerRight {
			answer.isRight = true
			h.points += answer.points
		} else {
			answer.isRight = false
		}
	}
}
