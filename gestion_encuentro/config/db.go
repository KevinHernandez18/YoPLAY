package config

import (
	"database/sql" // Para conexiones con SQL
	"fmt"
	"log" // Para imprimir y manejar logs

	_ "github.com/lib/pq" // Driver PostgreSQL
)

var DB *sql.DB // Instancia global de la base de datos

// connectDB establece conexión con postgreSQL

func ConnectDB() {
	// Variables para la conexión
	host := "localhost"
	port := 5432
	user := "postgres"
	password := "postgres"
	dbname := "yoplay"
	schema := "gestion_encuentro"

	// Cadena de conexión
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s sslmode=disable",
		host, port, user, password, dbname, schema,
	)

	// Abrir conexión DB
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error al conectar:", err)
	}

	// verificación de conexión
	err = db.Ping()
	if err != nil {
		log.Fatal("No se puede conectar:", err)
	}

	log.Println("Conexión a base de datos exitosa.")
	fmt.Println("Conectado a la db:", dbname, "Y esquema:", schema)

	DB = db // Asignar a conexión global
}
