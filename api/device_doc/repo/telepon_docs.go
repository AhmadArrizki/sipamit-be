package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type TeleponDoc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Departemen string         `json:"departemen" bson:"departemen"`
	User       string         `json:"user" bson:"user"`
	Ext        string         `json:"ext" bson:"ext"`
	Merk       string         `json:"merk" bson:"merk"`
	Tipe       string         `json:"tipe" bson:"tipe"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type TeleponDocCollRepository struct {
	coll *mongo.Collection
}

func NewTeleponDocRepository(db *mongo.Database) *TeleponDocCollRepository {
	return &TeleponDocCollRepository{
		coll: db.Collection("telepon_docs"),
	}
}

func (r *TeleponDocCollRepository) FindAll() (*[]TeleponDoc, error) {
	var teleponDocs []TeleponDoc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &teleponDocs)
	if err != nil {
		return nil, err
	}
	return &teleponDocs, nil
}

func (r *TeleponDocCollRepository) FindOneByID(id bson.ObjectID) (*TeleponDoc, error) {
	var teleponDoc TeleponDoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&teleponDoc)
	if err != nil {
		return nil, err
	}
	return &teleponDoc, nil
}

func (r *TeleponDocCollRepository) InsertOne(teleponDoc *TeleponDoc) error {
	_, err := r.coll.InsertOne(context.TODO(), teleponDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeleponDocCollRepository) UpdateOneByID(id bson.ObjectID, teleponDoc *TeleponDoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": teleponDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeleponDocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
