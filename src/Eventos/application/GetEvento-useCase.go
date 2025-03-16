package application

import "Eventos/src/Eventos/domain"

type GetEventos struct {
	repo domain.IEvento
}

func NewGetEventos(repo domain.IEvento) *GetEventos {
	return &GetEventos{repo: repo}
}

func (ge *GetEventos) Execute() ([]domain.Evento, error) {
	return ge.repo.ObtenerTodos()
}
