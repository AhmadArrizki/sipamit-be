package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type TOADoc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Nama       string         `json:"nama" bson:"nama"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Kode       string         `json:"kode" bson:"kode"`
	Posisi     string         `json:"posisi" bson:"posisi"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type TOADocCollRepository struct {
	coll *mongo.Collection
}

func NewTOADocRepository(db *mongo.Database) *TOADocCollRepository {
	return &TOADocCollRepository{
		coll: db.Collection("toa_docs"),
	}
}

func (r *TOADocCollRepository) FindAll() (*[]TOADoc, error) {
	var toaDocs []TOADoc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &toaDocs)
	if err != nil {
		return nil, err
	}
	return &toaDocs, nil
}

func (r *TOADocCollRepository) FindOneByID(id bson.ObjectID) (*TOADoc, error) {
	var toaDoc TOADoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&toaDoc)
	if err != nil {
		return nil, err
	}
	return &toaDoc, nil
}

func (r *TOADocCollRepository) InsertOne(toaDoc *TOADoc) error {
	_, err := r.coll.InsertOne(context.TODO(), toaDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *TOADocCollRepository) UpdateOneByID(id bson.ObjectID, toaDoc *TOADoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": toaDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TOADocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
