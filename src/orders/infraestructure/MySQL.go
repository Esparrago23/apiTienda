package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"mi-tienda-online/src/core"
	"mi-tienda-online/src/orders/domain/entities"
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

func (mysql *MySQL) FindById(order_id int) (entities.Order, error) {
	order := entities.Order{}
	query := fmt.Sprintf("SELECT * FROM orders WHERE order_id = %d", order_id)
	row := mysql.conn.DB.QueryRow(query)

	var createdAtStr, updatedAtStr []uint8
	err := row.Scan(&order.OrderID, &order.UserID, &order.TotalAmount, &order.Status, &createdAtStr, &updatedAtStr)

	if err != nil {
		if err == sql.ErrNoRows {
			return order, fmt.Errorf("orden con id %d no encontrada", order_id)
		}
		return order, fmt.Errorf("error al obtener orden: %v", err)
	}
	order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAtStr))
	if err != nil {
		return order, fmt.Errorf("error al parsear la fecha de creación: %v", err)
	}
	order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", string(updatedAtStr))
	if err != nil {
		return order, fmt.Errorf("error al parsear la fecha de actualización: %v", err)
	}

	return order, nil
}

func (mysql *MySQL) FindAll() ([]entities.Order, error) {
	orders := []entities.Order{}
	query := "SELECT * FROM orders"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener órdenes: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var createdAtStr, updatedAtStr string
		order := entities.Order{}
		if err := rows.Scan(&order.OrderID, &order.UserID, &order.TotalAmount, &order.Status, &createdAtStr, &updatedAtStr); err != nil {
			return nil, fmt.Errorf("error al escanear órdenes: %v", err)
		}
		order.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
		if err != nil {
			return nil, fmt.Errorf("error al parsear la fecha de creación: %v", err)
		}
		order.UpdatedAt, err = time.Parse("2006-01-02 15:04:05", updatedAtStr)
		if err != nil {
			return nil, fmt.Errorf("error al parsear la fecha de actualización: %v", err)
		}
		orders = append(orders, order)
	}
	return orders, nil
}

func (mysql *MySQL) Save(order *entities.Order) error {
	query := fmt.Sprintf("INSERT INTO orders (user_id, total_amount, status) VALUES (%d, %f, '%s')", order.UserID, order.TotalAmount, order.Status)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al guardar orden: %v", err)
	}
	return nil
}

func (mysql *MySQL) Update(order *entities.Order) error {
	query := fmt.Sprintf("UPDATE orders SET user_id = %d, total_amount = %f, status = '%s' WHERE order_id = %d", order.UserID, order.TotalAmount, order.Status, order.OrderID)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al actualizar orden: %v", err)
	}
	return nil
}

func (mysql *MySQL) Delete(order_id int) error {
	query := fmt.Sprintf("DELETE FROM orders WHERE order_id = %d", order_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al eliminar orden: %v", err)
	}
	return nil
}
