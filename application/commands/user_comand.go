package commands

type UserCommand struct {
	DNI            int64  `json:"dni"`
	Name           string `json:"name"`
	LastName       string `json:"last_name"`
	QuantityMovies int8   `json:"quantity_movies"`
}
