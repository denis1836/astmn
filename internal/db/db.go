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
		name TEXT NOT NULL,
		version TEXT NOT NULL,
		preset TEXT NOT NULL,
		installed_at TEXT DEFAULT CURRENT_TIMESTAMP
	) STRICT;

	CREATE TABLE IF NOT EXISTS Package_Files (
		id INTEGER PRIMARY KEY,
		package_id INTEGER NOT NULL,
		name TEXT NOT NULL,
		relative_path TEXT NOT NULL, 
		sha256 TEXT NOT NULL,
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

func InsertPackage(name, version, preset string) (int64, error) {
	insertPackageQuery := `
	INSERT INTO Installed_Packages(name, version, preset)
	VALUES (?, ?, ?);
	`
	res, err := Pool.Exec(insertPackageQuery, name, version, preset)
	if err != nil {
		return 0, fmt.Errorf("failed to insert package: %w", err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("failed to get last insert id: %w", err)
	}

	return id, nil
}

func InsertFile(packageId int64, name, relPath, hash string, size int64) error {
	insertPackageFileQuery := `
	INSERT INTO Package_Files(package_id, name, relative_path, sha256, file_size)
	VALUES (?, ?, ?, ?, ?);
	`

	_, err := Pool.Exec(insertPackageFileQuery, packageId, name, relPath, hash, size)
	if err != nil {
		log.Errorf("unable to insert file into db: %v", err)
		return err
	}

	return nil
}

func GetPackageNameById(packageId int64) (string, error) {
	getPackageNameQuery := `
	SELECT name FROM Installed_Packages
	WHERE id = ?
	`

	var name string
	err := Pool.QueryRow(getPackageNameQuery, packageId).Scan(&name)

	if err != nil {
		if err == sql.ErrNoRows {
			return "", fmt.Errorf("package with id %d not found", packageId)
		}
		log.Errorf("unable to get package name: %v", err)
		return "", err
	}

	return name, nil
}
