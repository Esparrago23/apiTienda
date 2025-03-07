package infraestructure

import (
	"database/sql"
	"fmt"
	"log"
	"mi-tienda-online/src/core"
	"mi-tienda-online/src/payments/domain/entities"
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

func (mysql *MySQL) FindById(payment_id int) (entities.Payment, error) {
	payment := entities.Payment{}
	query := fmt.Sprintf("SELECT * FROM payments WHERE payment_id = %d", payment_id)
	row := mysql.conn.DB.QueryRow(query)

	var paymentDateStr []uint8
	err := row.Scan(&payment.PaymentID, &payment.OrderID, &payment.PaymentStatus, &paymentDateStr)
	if err != nil {
		if err == sql.ErrNoRows {
			return payment, fmt.Errorf("pago con id %d no encontrado", payment_id)
		}
		return payment, fmt.Errorf("error al obtener pago: %v", err)
	}
	payment.PaymentDate, err = time.Parse("2006-01-02 15:04:05", string(paymentDateStr))
	if err != nil {
		return payment, fmt.Errorf("error al parsear la fecha de pago: %v", err)
	}

	return payment, nil
}

func (mysql *MySQL) FindAll() ([]entities.Payment, error) {
	payments := []entities.Payment{}
	query := "SELECT * FROM payments"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener pagos: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var paymentDateStr string
		payment := entities.Payment{}
		if err := rows.Scan(&payment.PaymentID, &payment.OrderID, &payment.PaymentStatus, &paymentDateStr); err != nil {
			return nil, fmt.Errorf("error al escanear pagos: %v", err)
		}
		payment.PaymentDate, err = time.Parse("2006-01-02 15:04:05", paymentDateStr)
		if err != nil {
			return nil, fmt.Errorf("error al parsear la fecha de pago: %v", err)
		}
		payments = append(payments, payment)
	}
	return payments, nil
}

func (mysql *MySQL) Save(payment *entities.Payment) error {
	query := fmt.Sprintf("INSERT INTO payments (order_id, payment_status) VALUES (%d, '%s')", payment.OrderID, payment.PaymentStatus)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al guardar pago: %v", err)
	}
	return nil
}

func (mysql *MySQL) Update(payment *entities.Payment) error {
	query := fmt.Sprintf("UPDATE payments SET order_id = %d, payment_status = '%s' WHERE payment_id = %d", payment.OrderID, payment.PaymentStatus, payment.PaymentID)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al actualizar pago: %v", err)
	}
	return nil
}

func (mysql *MySQL) Delete(payment_id int) error {
	query := fmt.Sprintf("DELETE FROM payments WHERE payment_id = %d", payment_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al eliminar pago: %v", err)
	}
	return nil
}
