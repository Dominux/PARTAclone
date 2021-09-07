package question

import (
	"./answer"
	"./points"
)


type Question struct {
	number int
	text string
	options []string
	rightAnswers []*answer.Answer
	points points.Points
}
