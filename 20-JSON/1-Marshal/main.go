package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome string `json:"nome"`
	Racas string `json:"raca"`
	Idade uint  `json:"idade"`
}

func main() {
	c := cachorro{"Rex", "Dalmata", 3}
	fmt.Println(c)

	cachorroEmJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	// retorna um slic de byte ou []uint8
	fmt.Println(cachorroEmJSON)
	// esse retorna um json com esta biblioteca bytes.NewBuffer()
	fmt.Println(bytes.NewBuffer(cachorroEmJSON))

	c2 := map[string] string {
		"nome": "Tody",
		"raca": "Poodle",
	}

	cachorro2EmJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EmJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EmJSON))
}