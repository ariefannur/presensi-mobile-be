package database

import (
	"presensi-mobile/queries"
)

// Queries struct for collect all app queries.
type Queries struct {
	*queries.UserQueries
	*queries.SessionQueries
	*queries.PresensiQueries
}

// OpenDBConnection func for opening database connection.
func OpenDBConnection() (*Queries, error) {
	// Define a new PostgreSQL connection.
	db, err := PostgreSQLConnection()
	if err != nil {
		return nil, err
	}

	return &Queries{
		// Set queries from models:
		UserQueries:     &queries.UserQueries{DB: db},    // from User model
		SessionQueries:  &queries.SessionQueries{DB: db}, // from Session model
		PresensiQueries: &queries.PresensiQueries{DB: db},
	}, nil
}
