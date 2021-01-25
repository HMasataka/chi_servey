package survey

import (
	"fmt"
	"net/http"
)

func NewPets() *Pets {
	return &Pets{}
}

type Pets struct {
}

func (p *Pets) FindPets(w http.ResponseWriter, r *http.Request, params FindPetsParams) {
	user := r.Context().Value("user").(string)
	w.Write([]byte(fmt.Sprintf("hi %s", user)))
}

func (p *Pets) AddPet(w http.ResponseWriter, r *http.Request) {
	user := r.Context().Value("user").(string)
	w.Write([]byte(fmt.Sprintf("hi %s", user)))
}

func (p *Pets) FindPetById(w http.ResponseWriter, r *http.Request, id int64) {
	user := r.Context().Value("user").(string)
	w.Write([]byte(fmt.Sprintf("hi %s", user)))
}

func (p *Pets) DeletePet(w http.ResponseWriter, r *http.Request, id int64) {
	user := r.Context().Value("user").(string)
	w.Write([]byte(fmt.Sprintf("hi %s", user)))
}
