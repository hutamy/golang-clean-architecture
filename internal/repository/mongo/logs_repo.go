package mongo

import (
	"context"
	"time"

	"github.com/hutamy/golang-clean-architecture/internal/domain"
	"github.com/hutamy/golang-clean-architecture/internal/domain/entity"
	"go.mongodb.org/mongo-driver/mongo"
)

type LogsRepository struct {
	collection *mongo.Collection
}

func NewLogsRepository(dbManager *DBManager) domain.LogsRepository {
	return &LogsRepository{
		collection: dbManager.GetCollection("logs"),
	}
}

func (r *LogsRepository) InsertLog(ctx context.Context, log *entity.Log) error {
	log.CreatedAt = time.Now()
	_, err := r.collection.InsertOne(ctx, log)
	return err
}

func (r *LogsRepository) FindLogs(ctx context.Context, filter map[string]interface{}) ([]*entity.Log, error) {
	var logs []*entity.Log
	cursor, err := r.collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var log entity.Log
		if err := cursor.Decode(&log); err != nil {
			return nil, err
		}
		logs = append(logs, &log)
	}

	return logs, nil
}
