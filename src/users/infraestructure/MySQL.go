package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"mi-tienda-online/src/core"
	"mi-tienda-online/src/users/domain/entities"
	"time"
)

type MySQL struct {
	conn *core.Conn_MySQL
}

func NewMySQL() *MySQL {
	conn := core.GetDBPool()
	if conn.Err != "" {
		log.Fatalf("Error al configurar el pool de conexiones: %v", conn.Err)
	}
	return &MySQL{conn: conn}
}

func (mysql *MySQL) FindById(user_id int) (entities.User, error) {
	user := entities.User{}
	query := fmt.Sprintf("SELECT * FROM users WHERE user_id = %d", user_id)
	row := mysql.conn.DB.QueryRow(query)

	var createdAtStr []uint8
	err := row.Scan(&user.User_id, &user.Name, &user.Email, &user.Address, &createdAtStr)

	if err != nil {
		if err == sql.ErrNoRows {
			return user, fmt.Errorf("usuario con id %d no encontrado", user_id)
		}
		return user, fmt.Errorf("error al obtener usuario: %v", err)
	}
	user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAtStr))
	if err != nil {
		return user, fmt.Errorf("error al parsear la fecha: %v", err)
	}

	return user, nil
}

func (mysql *MySQL) FindAll() ([]entities.User, error) {
	users := []entities.User{}
	query := "SELECT * FROM users"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener usuarios: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var createdAtStr string
		user := entities.User{}
		if err := rows.Scan(&user.User_id, &user.Name, &user.Email, &user.Address, &createdAtStr); err != nil {
			return nil, fmt.Errorf("error al escanear usuarios: %v", err)
		}
		user.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("error al parsear la fecha: %v", err)
		}
		users = append(users, user)
	}
	return users, nil
}

func (mysql *MySQL) Save(user *entities.User) error {
	query := fmt.Sprintf("INSERT INTO users (name, email, address) VALUES ('%s', '%s', '%s')", user.Name, user.Email, user.Address)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al guardar usuario: %v", err)
	}
	return nil
}

func (mysql *MySQL) Update(user *entities.User) error {
	query := fmt.Sprintf("UPDATE users SET name = '%s', email = '%s', address = '%s' WHERE user_id = %d", user.Name, user.Email, user.Address, user.User_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al actualizar usuario: %v", err)
	}
	return nil
}

func (mysql *MySQL) Delete(user_id int) error {
	query := fmt.Sprintf("DELETE FROM users WHERE user_id = %d", user_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al eliminar usuario: %v", err)
	}
	return nil
}
