package reservation

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

func (r *repository) FindAll() (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("is_booked = false").Find(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) FindByUserID(id string) (*model.Reservations, error) {
	var reservation model.Reservations
	r.db.Preload("VehicleID").Preload("UserID")
	data := r.db.Where("is_booked = false").Find(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}


func (r *repository) History() (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("is_booked = true").Find(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Search(name string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("vehicle.name like ?", "%"+name+"%").Limit(3).Find(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Sort(name string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("is_booked = true").Order(name).Find(&reservation)
	
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Create(reservation *model.Reservation) (*model.Reservation, error) {
	data := r.db.Create(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return reservation, nil
}

func (r *repository) Update(reservation *model.Reservation, id string) (*model.Reservation, error) {
	data := r.db.Model(&reservation).Where("reservation_id = ?", id).Updates(&reservation)
	if data.Error != nil {
		return nil, data.Error
	}

	return reservation, nil
}

func (r *repository) Pay(reservation *model.Reservation, vehicle *model.Vehicle, id string) (*model.Reservation, error) {
	err := r.db.Where("reservation_id = ?", id).First(&reservation).Error
	if err != nil {
		return nil, err
	}
	reservation.IsBooked = true
	r.db.Save(&reservation)

	err = r.db.Where("vehicle_id = ?", reservation.VehicleID).First(&vehicle).Error
	if err != nil {
		return nil, err
	}

	vehicle.Point = vehicle.Point + 1
	r.db.Save(vehicle)

	return reservation, nil
}

func (r *repository) Delete(reservation *model.Reservation, id string) (*model.Reservation, error) {
	err := r.db.Where("reservation_id = ?", id).Delete(&reservation).Error
	if err != nil {
		return nil, err
	}

	return reservation, nil
}
