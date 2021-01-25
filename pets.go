package survey

import (
	"net/http"
)

func NewPets() *Pets {
	return &Pets{}
}

type Pets struct {
}

func (p *Pets) FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams) {
}

func (p *Pets) AddPet(w http.ResponseWriter, r *http.Request) {
}

func (p *Pets) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
}

func (p *Pets) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
}
