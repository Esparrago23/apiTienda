package infraestructure

import (
	"mi-tienda-online/src/core"
	"mi-tienda-online/src/products/domain/entities"
	"fmt"
	"log"
	"time"
	"database/sql"
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

func (mysql *MySQL) FindById(product_id int) (entities.Product, error) {
	product := entities.Product{}
	query := fmt.Sprintf("SELECT * FROM products WHERE product_id = %d", product_id)
	row := mysql.conn.DB.QueryRow(query)

	var createdAtStr []uint8
	err := row.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.Stock, &createdAtStr)

	if err != nil {
        if err == sql.ErrNoRows {
            return product, fmt.Errorf("producto con id %d no encontrado", product_id)
        }
        return product, fmt.Errorf("error al obtener producto: %v", err)
    }
	product.CreatedAt, err = time.Parse("2006-01-02 15:04:05", string(createdAtStr))
    if err != nil {
        return product, fmt.Errorf("error al parsear la fecha: %v", err)
    }

	return product, nil
}

func (mysql *MySQL) FindAll() ([]entities.Product, error) {
	products := []entities.Product{}
	query := "SELECT * FROM products"
	rows, err := mysql.conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error al obtener productos: %v", err)
	}
	defer rows.Close()
	for rows.Next() {
		var createdAtStr string
		product := entities.Product{}
		if err := rows.Scan(&product.ProductID, &product.Name, &product.Description, &product.Price, &product.Stock, &createdAtStr); err != nil {
            return nil, fmt.Errorf("error al escanear productos: %v", err)
        }
        product.CreatedAt, err = time.Parse("2006-01-02 15:04:05", createdAtStr)
        if err != nil {
            return nil, fmt.Errorf("error al parsear la fecha: %v", err)
        }
		products = append(products, product)
	}
	return products, nil
}

func (mysql *MySQL) Save(product *entities.Product) error {
	query := fmt.Sprintf("INSERT INTO products (name, description, price, stock) VALUES ('%s', '%s', %f, %d)", product.Name, product.Description, product.Price, product.Stock)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al guardar producto: %v", err)
	}
	return nil
}

func (mysql *MySQL) Update(product *entities.Product) error {
	query := fmt.Sprintf("UPDATE products SET name = '%s', description = '%s', price = %f, stock = %d WHERE product_id = %d", product.Name, product.Description, product.Price, product.Stock, product.ProductID)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al actualizar producto: %v", err)
	}
	return nil
}

func (mysql *MySQL) Delete(product_id int) error {
	query := fmt.Sprintf("DELETE FROM products WHERE product_id = %d", product_id)
	_, err := mysql.conn.ExecutePreparedQuery(query)
	if err != nil {
		return fmt.Errorf("error al eliminar producto: %v", err)
	}	
	return nil
}