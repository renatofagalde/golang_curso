package cachorro

// tipo de dado
type Cachorro struct {
	Nome  string `json:"nome"` // `json:"-"`  para ignorar no json
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}
