package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type Checkpoint struct {
	ID         bson.ObjectID `json:"_id" bson:"_id"`
	Device     string        `json:"device" bson:"device"`
	Checkpoint []string      `json:"checkpoint" bson:"checkpoint"`
	Updated    *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
}

type CheckpointCollRepository struct {
	coll *mongo.Collection
}

func NewCheckpointRepository(db *mongo.Database) *CheckpointCollRepository {
	return &CheckpointCollRepository{
		coll: db.Collection("checkpoint"),
	}
}

func (r *CheckpointCollRepository) InsertMany(checkpoints []Checkpoint) error {
	_, err := r.coll.InsertMany(context.TODO(), checkpoints)
	if err != nil {
		return err
	}
	return nil
}

func (r *CheckpointCollRepository) FindByDevice(device string) (*Checkpoint, error) {
	var checkpoint Checkpoint
	filter := bson.M{
		"device": device,
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&checkpoint)
	if err != nil {
		return nil, err
	}
	return &checkpoint, nil
}

func (r *CheckpointCollRepository) UpdateByDevice(device string, checkpoint *Checkpoint) error {
	filter := bson.M{
		"device": device,
	}
	update := bson.M{
		"$set": checkpoint,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *CheckpointCollRepository) Count() (int64, error) {
	filter := bson.M{}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
