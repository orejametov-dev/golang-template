// database/mysql_connection.go
package database

import (
	"database/sql"
	"experiment/internal/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnection struct {
	DB *sql.DB
}

func NewMySQLConnection(cfg *config.DBConfig) (*MySQLConnection, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", cfg.Username, cfg.Password, cfg.Host, cfg.Port, cfg.Database)
	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("failed to open MySQL connection: %w", err)
	}

	if err := db.Ping(); err != nil {
		db.Close()
		return nil, fmt.Errorf("failed to ping MySQL: %w", err)
	}

	return &MySQLConnection{DB: db}, nil
}

// Query выполнение SQL-запроса в MySQL
func (conn *MySQLConnection) Query(query string) ([]string, error) {
	rows, err := conn.DB.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to execute MySQL query: %w", err)
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var result string
		if err := rows.Scan(&result); err != nil {
			return nil, fmt.Errorf("failed to scan MySQL row: %w", err)
		}
		results = append(results, result)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("error during MySQL rows iteration: %w", err)
	}

	return results, nil
}
