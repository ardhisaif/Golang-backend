package seeders

import "github.com/ardhisaif/golang_backend/database/orm/model"

var VehicleSeed = model.Vehicles{
	model.Vehicle{
		Type: "car",
		Name: "tesla",
		Image: "",
		Location: "Jakarta",
		Capacity: 2,
		Price: 200000,
		Point: 3,
	},
}