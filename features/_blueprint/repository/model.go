package repository

import (
	"blueprint_golang/features/_blueprint"

	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type model struct {
	db *gorm.DB
}

func New(db *gorm.DB) _blueprint.Repository {
	return &model{
		db: db,
	}
}

func (mdl *model) Paginate(page, size int) ([]_blueprint.Placeholder, error) {
	var placeholders []_blueprint.Placeholder

	offset := (page - 1) * size

	err := mdl.db.Offset(offset).Limit(size).Find(&placeholders).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return placeholders, nil
}

func (mdl *model) Insert(newPlaceholder _blueprint.Placeholder) error {
	err := mdl.db.Create(&newPlaceholder).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}

func (mdl *model) SelectByID(placeholderID uint) (*_blueprint.Placeholder, error) {
	var placeholder _blueprint.Placeholder
	err := mdl.db.First(&placeholder, placeholderID).Error

	if err != nil {
		log.Error(err)
		return nil, err
	}

	return &placeholder, nil
}

func (mdl *model) Update(placeholder _blueprint.Placeholder) error {
	err := mdl.db.Updates(&placeholder).Error

	if err != nil {
		log.Error(err)
	}

	return err
}

func (mdl *model) DeleteByID(placeholderID uint) error {
	err := mdl.db.Delete(&_blueprint.Placeholder{}, placeholderID).Error

	if err != nil {
		log.Error(err)
		return err
	}

	return nil
}
