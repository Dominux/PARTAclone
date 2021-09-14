package repository

import (
	"gorm.io/gorm"

	"github.com/SPARTAclone/SPARTA/internal/homeworks/entity"
)


type HomeworkRepository struct {
	gorm.Model
	entity.Homework
}

func CreateHomework(h *entity.Homework, db gorm.DB) (*entity.Homework, error) {

	// 1. Creating HomeworkRepository
	hr := new(HomeworkRepository)

	db.Create(hr)
}
