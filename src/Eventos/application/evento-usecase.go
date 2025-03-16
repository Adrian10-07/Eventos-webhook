package application

import "Eventos/src/Eventos/domain"

type CreateEvento struct {
	repo domain.IEvento
}

func NewCreateEvento(repo domain.IEvento) *CreateEvento {
	return &CreateEvento{repo: repo}
}

func (ce *CreateEvento) Execute(e domain.Evento) error {
	return ce.repo.Guardar(&e)
}
