package db

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	"astmn/internal/log"

	_ "modernc.org/sqlite"
)

var Pool *sql.DB

type InstalledPackage struct {
	ID        string
	Version   string
	Preset    string
	CreatedAt string
}

type PackageFile struct {
	ID           int64
	PackageID    string
	RelativePath string
	SHA256       string
	FileSize     int64
}

func OpenPool(dbPath string) error {
	dir := filepath.Dir(dbPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Errorf("failed to create db directory (%v): %v", dir, err)
		return err
	}

	dsn := fmt.Sprintf("%s?_pragma=foreign_keys(1)&_pragma=journal_mode(WAL)", dbPath)

	db, err := sql.Open("sqlite", dsn)
	if err != nil {
		log.Errorf("Unable to open db pool: %v", err)
		return err
	}

	Pool = db
	Pool.SetConnMaxLifetime(0)
	Pool.SetMaxIdleConns(2)
	Pool.SetMaxOpenConns(5)

	if err := PingDB(); err != nil {
		return err
	}

	err = initSchema()
	if err != nil {
		log.Errorf("failed to init db schema: %v", err)
	}

	return nil
}

func ClosePool() error {
	if Pool == nil {
		return nil
	}

	err := Pool.Close()
	if err != nil {
		log.Errorf("Unable to close db pool: %v", err)
		return err
	}

	return nil
}

func PingDB() error {
	if err := Pool.Ping(); err != nil {
		log.Errorf("Unable to ping db: %v", err)
		return err
	}

	return nil
}

func initSchema() error {
	query := `
	CREATE TABLE IF NOT EXISTS Installed_Packages (
		id INTEGER PRIMARY KEY,
		version TEXT NOT NULL,
		preset TEXT NOT NULL,
		installed_at DATETIME DEFAULT CURRENT_TIMESTAMP
	) STRICT;

	CREATE TABLE IF NOT EXISTS Package_Files (
		id INTEGER PRIAMRY KEY,
		package_id TEXT NOT NULL, 
		relative_path TEXT NOT NULL, 
		sha256 BLOB NOT NULL,
		file_size INTEGER NOT NULL,

		FOREIGN KEY(package_id) REFERENCES Installed_Packages(id) ON DELETE CASCADE
	) STRICT;
	`

	_, err := Pool.Exec(query)
	if err != nil {
		log.Errorf("failed to run query(init schema): %v", err)
		return err
	}

	return nil
}
