package store

import (
	"database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
	"github.com/oluwatunmise/janus-proxy/services/controlplane/internal/config"
)

func NewDB(cfg config.Config) (*sql.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s",
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBName,
	)

	db, err := sql.Open("postgres", dsn)

	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)

	db.SetMaxIdleConns(10)

	db.SetConnMaxLifetime(5 * time.Minute)

	if err := db.Ping(); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}
