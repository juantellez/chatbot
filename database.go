// database.go

package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

// Database representa la conexión a la base de datos
type Database struct {
	db *sql.DB
}

// NewDatabase crea una nueva instancia de Database y establece la conexión a la base de datos
func NewDatabase(databasePath string) (*Database, error) {
	db, err := sql.Open("sqlite3", databasePath)
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %v", err)
	}

	// Verificar la conexión a la base de datos
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("failed to ping database: %v", err)
	}

	return &Database{
		db: db,
	}, nil
}

// Close cierra la conexión a la base de datos
func (db *Database) Close() error {
	return db.db.Close()
}

// SaveUser guarda los datos de un usuario en la base de datos
func (db *Database) SaveUser(user *User) error {
	// Implementa aquí la lógica para guardar los datos del usuario en la base de datos
	// Por ejemplo, puedes crear una tabla "users" y guardar los datos del usuario en ella
	// Utiliza sentencias SQL para realizar las operaciones de inserción y actualización

	return nil
}

// LoadUser carga los datos de un usuario desde la base de datos
func (db *Database) LoadUser(userID string) (*User, error) {
	// Implementa aquí la lógica para cargar los datos del usuario desde la base de datos
	// Por ejemplo, puedes consultar la tabla "users" para obtener los datos del usuario
	// Utiliza sentencias SQL para realizar la consulta y mapea los resultados a la estructura User

	return nil, nil
}
