package core


import (
	"database/sql"
	"fmt"
	"log"

	"github.com/joho/godotenv"
	_ "github.com/go-sql-driver/mysql"
)

type Conn_MySQL struct {
	DB  *sql.DB
	Err string
}

func GetDBPool() *Conn_MySQL {
	_ = godotenv.Load()

	// llmar variables de entorno
	dbHost:="localhost"
dbUser:="root"
dbPass:="adrian0710200512#12#"
dbSchema:= "Sensores"
dsn := fmt.Sprintf("%s:%s@tcp(%s:3306)/%s", dbUser, dbPass, dbHost, dbSchema)
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Error al abrir la base de datos: %v", err)
	}

	db.SetMaxOpenConns(10)

	if err := db.Ping(); err != nil {
		log.Fatalf("Error al verificar conexión: %v", err)
	}

	return &Conn_MySQL{DB: db}
}