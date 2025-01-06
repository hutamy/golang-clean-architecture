package postgres

import (
	"log"

	"golang.org/x/exp/rand"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DBManager struct {
	Master   *gorm.DB
	Replicas []*gorm.DB
}

func NewDBManager(masterDSN string, replicaDSNs []string) *DBManager {
	// Connect to the master database
	master, err := gorm.Open(postgres.Open(masterDSN), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to master database: %v", err)
	}

	// Connect to replicas
	var replicas []*gorm.DB
	for _, dsn := range replicaDSNs {
		replica, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to connect to replica: %v", err)
			continue
		}
		replicas = append(replicas, replica)
	}

	if len(replicas) == 0 {
		log.Fatal("No replicas available")
	}

	return &DBManager{
		Master:   master,
		Replicas: replicas,
	}
}

// GetReplica returns a random replica connection (for load balancing)
func (dbm *DBManager) GetReplica() *gorm.DB {
	if len(dbm.Replicas) == 0 {
		return dbm.Master // Fallback to master if no replicas are available
	}
	index := rand.Intn(len(dbm.Replicas))
	return dbm.Replicas[index]
}
