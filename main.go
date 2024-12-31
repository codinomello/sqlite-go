package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func createTable(db *sql.DB) {
	createTableSQL := `CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		age INTEGER
	);`
	_, err := db.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

func insertUser(db *sql.DB, name string, age int) {
	insertSQL := `INSERT INTO users (name, age) VALUES (?, ?)`
	_, err := db.Exec(insertSQL, name, age)
	if err != nil {
		log.Fatal(err)
	}
}

func readUsers(db *sql.DB) {
	rows, err := db.Query("SELECT id, name, age FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var name string
		var age int
		if err := rows.Scan(&id, &name, &age); err != nil {
			log.Fatal(err)
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}
}

func updateUser(db *sql.DB, id int, name string, age int) {
	updateSQL := `UPDATE users SET name = ?, age = ? WHERE id = ?`
	_, err := db.Exec(updateSQL, name, age, id)
	if err != nil {
		log.Fatal(err)
	}
}

func deleteUser(db *sql.DB, id int) {
	deleteSQL := `DELETE FROM users WHERE id = ?`
	_, err := db.Exec(deleteSQL, id)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	// Conectar ao banco de dados SQLite
	db, err := sql.Open("sqlite3", "./users.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Criar a tabela de usuários
	createTable(db)

	// Inserir dados
	insertUser(db, "John Doe", 30)
	insertUser(db, "Jane Doe", 28)

	// Ler dados
	fmt.Println("Users in database:")
	readUsers(db)

	// Atualizar dados
	updateUser(db, 1, "Johnathan Doe", 31)

	// Excluir dados
	deleteUser(db, 2)

	// Ler novamente após atualização e exclusão
	fmt.Println("\nUsers after update and deletion:")
	readUsers(db)
}
