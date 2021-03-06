package model

type User struct {
	DNI            int64  `json:"dni"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	QuantityMovies int8   `json:"quantity_movies"`
}

func (user *User) CreateUser(dni int64, name string, lastName string, quantityMovies int8) (User, error) {
	return User{
		DNI:            dni,
		Name:           name,
		LastName:       lastName,
		QuantityMovies: quantityMovies,
	}, nil
}
