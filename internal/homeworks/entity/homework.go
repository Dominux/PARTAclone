package entity

import (
	"github.com/SPARTAclone/SPARTA/internal/users/entity"
)

type Homework struct {
	student *User
	exercise *Exercise
	answers []*Answer
	points Points
	isGraded bool
}

// Auto check homework's questions by comparing 
// it's right answers with given by student 
func (h *Homework) autoCheck() {
	for _, answer := range h.Answers {

		isAnswerRight, err := answer.Question.IsAnswerRight(answer)

		if err != nil {
			continue
		}

		if isAnswerRight {
			answer.IsRight = true
			h.Points += answer.Points
		} else {
			answer.isRight = false
		}
	}
}


func CreateHomework(
	student *User
	exercise *Exercise
	answers []*Answer
	points Points
	isGraded bool
) (*Homework, error) {

	h := *Homework{
		student,
		exercise,
		answers,
		points,
		isGraded,
	}

	// TODO

}
