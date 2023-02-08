package vehicle

import (
	"github.com/ardhisaif/golang_backend/database/orm/model"
	"gorm.io/gorm"
)

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll() (*model.Vehicles, error) {
	var vehicles model.Vehicles
	data := r.db.Find(&vehicles)
	if data.Error != nil {
		return nil, data.Error
	}

	return &vehicles, nil
}

func (r *repository) Popular() (*model.Vehicles, error) {
	var vehicles model.Vehicles
	data := r.db.Order("point desc").Limit(3).Find(&vehicles)
	if data.Error != nil {
		return nil, data.Error
	}

	return &vehicles, nil
}

func (r *repository) Search(name string) (*model.Vehicles, error) {
	var vehicles model.Vehicles
	data := r.db.Where("name like ?", "%"+name+"%").Limit(3).Find(&vehicles)
	if data.Error != nil {
		return nil, data.Error
	}

	return &vehicles, nil
}

func (r *repository) Sort(name string) (*model.Vehicles, error) {
	var vehicles model.Vehicles
	data := r.db.Order(name).Find(&vehicles)
	if data.Error != nil {
		return nil, data.Error
	}

	return &vehicles, nil
}

func (r *repository) Create(vehicle *model.Vehicle) (*model.Vehicle, error) {
	data := r.db.Create(vehicle)
	if data.Error != nil {
		return nil, data.Error
	}

	return vehicle, nil
}

func (r *repository) Update(vehicle *model.Vehicle, id string) (*model.Vehicle, error) {
	data := r.db.Model(&vehicle).Where("vehicle_id = ?",id).Updates(&vehicle)
	if data.Error != nil {
		return nil, data.Error
	}

	return vehicle, nil
}

func (r *repository) Delete(vehicle *model.Vehicle, id string) (*model.Vehicle, error) {
	err := r.db.Where("vehicle_id = ?", id).Delete(&vehicle).Error
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}
