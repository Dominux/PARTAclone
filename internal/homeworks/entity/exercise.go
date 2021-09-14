package entity

import (
	users_entity "github.com/SPARTAclone/SPARTA/internal/users/entity"
	"github.com/SPARTAclone/SPARTA/internal/homeworks/repository"
)


type Exercise struct {
	Authors []*users_entity.User
	Questions []*Question
}

func CreateExercise(as []*users_entity.User, qs []*Question) (*Exercise, error) {
	e := Exercise{as, qs}

	if _, err := repository.SaveExercise(e); err != nil {
		return nil, err
	}

	return e, nil
}
