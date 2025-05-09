package database

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type PostgresConnectorInterface interface {
	GetDB() *gorm.DB
	Close() error
}

type PostgresConnector struct {
	db *gorm.DB
}

func NewPostgresConnector(dsn string) (PostgresConnectorInterface, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("❌ failed to connect to PostgreSQL: %w", err)
	}

	fmt.Println("✅ Connected to PostgreSQL!")
	return &PostgresConnector{db: db}, nil
}

func (p *PostgresConnector) GetDB() *gorm.DB {
	return p.db
}

func (p *PostgresConnector) Close() error {
	sqlDB, err := p.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

func ConnectAllPostgresDBs() (map[string]PostgresConnectorInterface, error) {
	connectors := make(map[string]PostgresConnectorInterface)

	dsns := map[string]string{
		"db1": os.Getenv("POSTGRES_DB1_DSN"),
		"db3": os.Getenv("POSTGRES_DB3_DSN"),
	}

	for name, dsn := range dsns {
		if dsn == "" {
			fmt.Printf("⚠️ DSN for %s is empty. Skipping...\n", name)
			continue
		}

		conn, err := NewPostgresConnector(dsn)
		if err != nil {
			fmt.Printf("❌ Failed to connect to %s: %v\n", name, err)
			continue
		}
		connectors[name] = conn
	}

	return connectors, nil
}
