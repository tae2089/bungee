package config

import "database/sql"

func NewDbConfig() (*sql.DB, error) {
	// open sqlite
	db, err := sql.Open("sqlite3", "~/bungee.db")
	if err != nil {
		return nil, err
	}
	defer db.Close()

	//create table if not exists
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ssh_info (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL UNIQUE, user VARCHAR(255) NOT NULL UNIQUE, host VARCHAR(255) NOT NULL UNIQUE, port INTEGER NOT NULL DEFAULT 22)")
	if err != nil {
		return nil, err
	}
	return db, nil
}
