package models

import _ "github.com/google/uuid"

func CreateUserTable() error {
	_, err := Db.Exec(`
        CREATE TABLE IF NOT EXISTS userg (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            email TEXT UNIQUE NOT NULL,
            role TEXT DEFAULT 'user' CHECK (role IN ('admin', 'user')),
            gender TEXT CHECK (gender IN ('male', 'female')) NOT NULL,
            dob DATE NOT NULL,
            created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
            updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
        )
    `)
	if err != nil {
		return err
	}
	return nil
}

