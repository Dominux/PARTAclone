package repository

import (
	"gorm.io/gorm"

	"github.com/SPARTAclone/SPARTA/internal/homeworks/entity"
)


type ExerciseRepository struct {
	gorm.Model
	entity.Exercise
}

func SaveExercise(e *entity.Exercise, db gorm.DB) (*ExerciseRepository, error) {
	er := ExerciseRepository{
		Authors: e.Authors,
		Questions: e.Questions,
	}

	if err := db.Create(&er).Error; err != nil {
		return nil, err
	}

	return er, nil
}
