package domain

type IEvento interface {
	Guardar(evento *Evento) error
	ObtenerTodos() ([]Evento, error)
}
