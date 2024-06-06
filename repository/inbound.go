package repository

import (
	"a21hc3NpZ25tZW50/models"

	"gorm.io/gorm"
)

type InboundRepository struct {
	db *gorm.DB
}

func NewSchoolRepo(db *gorm.DB) InboundRepository {
	return InboundRepository{db}
}

func (u *InboundRepository) AddInboundRepository(inbound models.Inbound) error {
	dataAdd := u.db.Create(&inbound)
	Alert := dataAdd.Error
	if Alert != nil {
		return Alert
	}

	return nil
}

func (u *InboundRepository) ReadInboundRepository() ([]models.Inbound, error) {
	var baca []models.Inbound
	dataRead := u.db.Find(&baca)
	Alert := dataRead.Error
	if Alert != nil {
		return []models.Inbound{}, Alert
	}

	return baca, nil
}

func (u *InboundRepository) FindAllInboundRepository(db *gorm.DB, id int) ([]models.Inbound, error) {
	var findData []models.Inbound
	read := u.db.Debug().Table("inbounds").Where("id = ?", id).First(&findData)
	Alert := read.Error

	if Alert != nil {
		return findData, Alert
	}

	return findData, nil
}

func (u *InboundRepository) UpdateDone(id int) error {
	DataAkhir := u.db.Model(&models.Inbound{}).
		Where("id = ?", id)
	Alert := DataAkhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}

func (u *InboundRepository) DeleteInboundRepository(id uint) error {

	DataAkhir := u.db.Where("id = ?", id).
		Delete(&models.Inbound{})
	Alert := DataAkhir.Error
	if Alert != nil {
		return Alert
	}
	return nil // TODO: replace this
}
