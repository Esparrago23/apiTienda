package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"mi-tienda-online/src/core"
	"mi-tienda-online/src/order_items/domain/entities"
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

func (mysql *MySQL) FindById(order_item_id int) (entities.OrderItem, error) {
	orderItem := entities.OrderItem{}
	query := fmt.Sprintf("SELECT * FROM order_items WHERE order_item_id = %d", order_item_id)
	row := mysql.conn.DB.QueryRow(query)

	err := row.Scan(&orderItem.OrderItemID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price, &orderItem.TotalPrice)
	if err != nil {
		if err == sql.ErrNoRows {
			return orderItem, fmt.Errorf("order item con id %d no encontrado", order_item_id)
		}
		return orderItem, fmt.Errorf("error al obtener order item: %v", err)
	}

	return orderItem, nil
}

func (mysql *MySQL) FindAll() ([]entities.OrderItem, error) {
	orderItems := []entities.OrderItem{}
	query := "SELECT * FROM order_items"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener order items: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		orderItem := entities.OrderItem{}
		if err := rows.Scan(&orderItem.OrderItemID, &orderItem.OrderID, &orderItem.ProductID, &orderItem.Quantity, &orderItem.Price, &orderItem.TotalPrice); err != nil {
			return nil, fmt.Errorf("error al escanear order items: %v", err)
		}
		orderItems = append(orderItems, orderItem)
	}
	return orderItems, nil
}

func (mysql *MySQL) Save(orderItem *entities.OrderItem) error {
	query := fmt.Sprintf("INSERT INTO order_items (order_id, product_id, quantity, price, total_price) VALUES (%d, %d, %d, %f, %f)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price, orderItem.TotalPrice)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al guardar order item: %v", err)
	}
	return nil
}

func (mysql *MySQL) Update(orderItem *entities.OrderItem) error {
	query := fmt.Sprintf("UPDATE order_items SET order_id = %d, product_id = %d, quantity = %d, price = %f, total_price = %f WHERE order_item_id = %d", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price, orderItem.TotalPrice, orderItem.OrderItemID)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al actualizar order item: %v", err)
	}
	return nil
}

func (mysql *MySQL) Delete(order_item_id int) error {
	query := fmt.Sprintf("DELETE FROM order_items WHERE order_item_id = %d", order_item_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al eliminar order item: %v", err)
	}
	return nil
}
