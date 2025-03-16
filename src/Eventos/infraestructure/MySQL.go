package infraestructure

import (
	"Eventos/src/Eventos/domain"
	"Eventos/src/core"
	"log"
	"strconv"
)

type MySQLEventoRepository struct {
	conn *core.Conn_MySQL
}

func NewMySQLEventoRepository() *MySQLEventoRepository {
	conn := core.GetDBPool()
	return &MySQLEventoRepository{conn: conn}
}

// Guardar: Inserta un nuevo evento en la base de datos
func (r *MySQLEventoRepository) Guardar(e *domain.Evento) error {
	query := "INSERT INTO eventos (tipo_sensor, valor) VALUES (?, ?)"
	_, err := r.conn.DB.Exec(query, e.TipoSensor, e.Valor)
	if err != nil {
		log.Printf("Error al guardar evento: %v", err)
		return err
	}

	log.Println("Evento guardado exitosamente")
	return nil
}

// ObtenerTodos: Obtiene todos los eventos registrados
func (r *MySQLEventoRepository) ObtenerTodos() ([]domain.Evento, error) {
	query := "SELECT id, tipo_sensor, valor, timestamp, creado_en FROM eventos"
	rows, err := r.conn.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var eventos []domain.Evento
	for rows.Next() {
		var evento domain.Evento
		if err := rows.Scan(&evento.ID, &evento.TipoSensor, &evento.Valor, &evento.Timestamp, &evento.CreadoEn); err != nil {
			return nil, err
		}
		eventos = append(eventos, evento)
	}
	return eventos, nil
}

// ObtenerPorID: Obtiene un evento por su ID
func (r *MySQLEventoRepository) ObtenerPorID(id string) (*domain.Evento, error) {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return nil, err
	}

	query := "SELECT id, tipo_sensor, valor, timestamp, creado_en FROM eventos WHERE id = ?"
	row := r.conn.DB.QueryRow(query, intID)

	var evento domain.Evento
	err = row.Scan(&evento.ID, &evento.TipoSensor, &evento.Valor, &evento.Timestamp, &evento.CreadoEn)
	if err != nil {
		return nil, err
	}
	return &evento, nil
}

// Eliminar: Elimina un evento por ID
func (r *MySQLEventoRepository) Eliminar(id string) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := "DELETE FROM eventos WHERE id = ?"
	_, err = r.conn.DB.Exec(query, intID)
	return err
}

// Actualizar: Modifica un evento existente
func (r *MySQLEventoRepository) Actualizar(id string, e *domain.Evento) error {
	intID, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	query := "UPDATE eventos SET tipo_sensor = ?, valor = ?, timestamp = ?, creado_en = ? WHERE id = ?"
	_, err = r.conn.DB.Exec(query, e.TipoSensor, e.Valor, e.Timestamp, e.CreadoEn, intID)
	return err
}
