package models

import (
	"database/sql"

	_ "github.com/lib/pq"
)

type Server struct {
	ID   string
	Hostname  string
	IP string
	Version  float32
}

func AllServers() ([]Server, error) {
	// This now uses the unexported global variable.
	rows, err := db.Query("SELECT * FROM servers")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var serv []Server

	for rows.Next() {
		var servs Server

		err := rows.Scan(&servs.ID, &servs.Hostname, &servs.IP, &servs.Version)
		if err != nil {
			return nil, err
		}

		serv = append(serv, servs)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return serv, nil
}
