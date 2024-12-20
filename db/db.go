package db

import (
	"log"

	"github.com/gocql/gocql"
)

var Session *gocql.Session

// Initialize the connection to ScyllaDB
func InitScyllaDB(clusterIPs []string, keyspace string) {
	cluster := gocql.NewCluster(clusterIPs...)
	cluster.Keyspace = keyspace
	cluster.Consistency = gocql.Quorum

	var err error
	Session, err = cluster.CreateSession()
	if err != nil {
		log.Fatal("Error connectiong to ScyllaDB:", err)
	}

	log.Println("Connected to ScyllaDB succssfully.")
}

// Close the session
func CloseScyllaDB() {
	if Session != nil {
		Session.Close()
		log.Println("ScyllaDB session closed.")
	}
}

// Create Enemy Table
func CreateEnemyTable() {
	query := `
	CREATE TABLE IF NOT EXISTS enemies (
	id UUID PRIMARY KEY,
	name TEXT,
	health INT,
	strength INT,
	dexterity INT
	constitution INT
	)`
	err := Session.Query(query).Exec()
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
	log.Println("Table 'users' created successfully.")
}
