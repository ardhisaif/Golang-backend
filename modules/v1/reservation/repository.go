package reservation

import (

	"github.com/ardhisaif/golang_backend/database/orm/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type repository struct {
	db *gorm.DB
}

func NewRepo(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) FindAll(userID string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("user_id = ?", userID).Where("is_booked = false").Preload(clause.Associations).Find(&reservation)
	for i := 0; i < len(reservation); i++ {
		reservation[i].User.Password = ""
	}
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) FindByUserID(id string, userID string) (*model.Reservations, error) {
	var reservation model.Reservations
	r.db.Preload("VehicleID").Preload("UserID")
	data := r.db.Where("user_id = ?", userID).Where("is_booked = false").Find(&reservation)
	for i := 0; i < len(reservation); i++ {
		reservation[i].User.Password = ""
	}
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) FindByID(id string, userID string) (*model.Reservation, error) {
	var reservation model.Reservation
	data := r.db.Where("user_id = ?", userID).Where("reservation_id = ?", id).Preload(clause.Associations).First(&reservation)
	reservation.User.Password = ""
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) History(userID string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("user_id = ?", userID).Where("is_booked = true").Find(&reservation)
	for i := 0; i < len(reservation); i++ {
		reservation[i].User.Password = ""
	}
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Search(name string, userID string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("user_id = ?", userID).Where("vehicle.name like ?", "%"+name+"%").Limit(3).Find(&reservation)
	for i := 0; i < len(reservation); i++ {
		reservation[i].User.Password = ""
	}
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Sort(name string, userID string) (*model.Reservations, error) {
	var reservation model.Reservations
	data := r.db.Where("user_id = ?", userID).Where("is_booked = true").Order(name).Find(&reservation)
	for i := 0; i < len(reservation); i++ {
		reservation[i].User.Password = ""
	}
	if data.Error != nil {
		return nil, data.Error
	}

	return &reservation, nil
}

func (r *repository) Create(reservation *model.Reservation) (*model.Reservation, error) {
	var user model.User
	var vehicle model.Vehicle
	r.db.Where("user_id = ?", reservation.UserID).Find(&user)
	r.db.Where("vehicle_id = ?", reservation.VehicleID).Find(&vehicle)
	reservation.TotalPayment = reservation.Quantity * int(vehicle.Price)

	data := r.db.Create(&reservation)
	if data.Error != nil { 
		return nil, data.Error
	}
	r.db.Where("reservation_id = ?", reservation.ReservationID).Preload(clause.Associations).First(&reservation)
	reservation.User.Password = ""

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
	tx := r.db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	vehicle.Point = reservation.Vehicle.Point + 1
	if err := tx.Model(&vehicle).Where("vehicle_id = ?", reservation.VehicleID).Updates(&vehicle).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	reservation.IsBooked = true
	if err := tx.Model(&reservation).Where("reservation_id = ?", id).Updates(&reservation).Error; err != nil {
		return nil, err
	}
	
	return reservation, nil
}

func (r *repository) Delete(reservation *model.Reservation, id string) (*model.Reservation, error) {
	err := r.db.Where("reservation_id = ?", id).Delete(&reservation).Error
	if err != nil {
		return nil, err
	}

	return reservation, nil
}
