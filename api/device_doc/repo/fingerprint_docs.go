package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type FingerprintDoc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Nama       string         `json:"nama" bson:"nama"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Kode       string         `json:"kode" bson:"kode"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type FingerprintDocCollRepository struct {
	coll *mongo.Collection
}

func NewFingerprintDocRepository(db *mongo.Database) *FingerprintDocCollRepository {
	return &FingerprintDocCollRepository{
		coll: db.Collection("fingerprint_docs"),
	}
}

func (r *FingerprintDocCollRepository) FindAll() (*[]FingerprintDoc, error) {
	var fpDocs []FingerprintDoc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &fpDocs)
	if err != nil {
		return nil, err
	}
	return &fpDocs, nil
}

func (r *FingerprintDocCollRepository) FindOneByID(id bson.ObjectID) (*FingerprintDoc, error) {
	var fpDoc FingerprintDoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&fpDoc)
	if err != nil {
		return nil, err
	}
	return &fpDoc, nil
}

func (r *FingerprintDocCollRepository) InsertOne(fpDoc *FingerprintDoc) error {
	_, err := r.coll.InsertOne(context.TODO(), fpDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *FingerprintDocCollRepository) UpdateOneByID(id bson.ObjectID, fpDoc *FingerprintDoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": fpDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *FingerprintDocCollRepository) DeleteOneByID(id bson.ObjectID) error {
	filter := bson.M{
		"_id": id,
	}

	update := bson.M{
		"$set": bson.M{"is_deleted": true},
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}
