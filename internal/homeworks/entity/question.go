package entity

import (
	"./answer"
	"./points"
)


type Question struct {
	Number int
	Text string
	Options []string
	RightAnswers []*answer.Answer
	Points points.Points
}
