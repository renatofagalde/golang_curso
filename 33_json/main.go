package main

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	//json.Marshal()  map -> json
	//json.Unmarshal() json -> map
}
