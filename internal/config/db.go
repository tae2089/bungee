package config

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"
)

func NewDbConfig() (*sql.DB, error) {
	// 사용자의 홈 디렉토리 경로 가져오기
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return nil, fmt.Errorf("홈 디렉토리 경로를 가져오는 데 실패했습니다. %s", err.Error())
	}

	// SQLite 데이터베이스 파일 경로 생성
	dbPath := filepath.Join(homeDir, "bungee.db")
	// open sqlite
	db, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		return nil, err
	}
	defer db.SetMaxIdleConns(0)

	//create table if not exists for sshregisterdto
	_, err = db.Exec("CREATE TABLE IF NOT EXISTS ssh_info (id INTEGER PRIMARY KEY AUTOINCREMENT, name VARCHAR(255) NOT NULL UNIQUE, user VARCHAR(255) NOT NULL UNIQUE, host VARCHAR(255) NOT NULL UNIQUE, port INTEGER NOT NULL DEFAULT 22, key TEXT, password VARCHAR(255))")
	if err != nil {
		return nil, err
	}
	return db, nil
}
