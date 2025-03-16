package mysql

import (
	"database/sql"
	"Eventos/src/core"
	"fmt"
	"Eventos/src/Eventos/domain"
)

// RepositorioEventos es la estructura que implementa la interfaz IEvento para interactuar con MySQL
type RepositorioEventos struct {
	conn *core.Conn_MySQL
}

// NuevoRepositorioEventos crea un nuevo repositorio de eventos
func NuevoRepositorioEventos(db *sql.DB) *RepositorioEventos {
	conn := core.GetDBPool()
	return &RepositorioEventos{conn:conn}
}

// Guardar guarda un evento en la base de datos
func (r *RepositorioEventos) Guardar(evento *domain.Evento) error {
	// Insertar un nuevo evento en la base de datos
	query := `INSERT INTO eventos (tipo_sensor, valor, timestamp, criado_en) VALUES (?, ?, ?, ?)`
	_, err := r.conn.DB.Exec(query, evento.TipoSensor, evento.Valor, evento.Timestamp, evento.CriadoEn)
	if err != nil {
		return fmt.Errorf("error al guardar evento: %v", err)
	}
	return nil
}

// ObtenerTodos obtiene todos los eventos desde la base de datos
func (r *RepositorioEventos) ObtenerTodos() ([]domain.Evento, error) {
	var eventos []domain.Evento

	// Obtener todos los eventos
	query := `SELECT id, tipo_sensor, valor, timestamp, criado_en FROM eventos`
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener eventos: %v", err)
	}
	defer rows.Close()

	// Iterar sobre los resultados y mapearlos a la estructura Evento
	for rows.Next() {
		var evento domain.Evento
		if err := rows.Scan(&evento.ID, &evento.TipoSensor, &evento.Valor, &evento.Timestamp, &evento.CriadoEn); err != nil {
			return nil, fmt.Errorf("error al escanear evento: %v", err)
		}
		eventos = append(eventos, evento)
	}

	return eventos, nil
}

// Eliminar elimina un evento por su ID
func (r *RepositorioEventos) Eliminar(id int) error {
	// Eliminar un evento
	query := `DELETE FROM eventos WHERE id = ?`
	_, err := r.conn.DB.Exec(query, id)
	if err != nil {
		return fmt.Errorf("error al eliminar evento: %v", err)
	}
	return nil
}
