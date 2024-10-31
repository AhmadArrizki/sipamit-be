package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type Telepon struct {
	ID         bson.ObjectID `json:"_id" bson:"_id"`
	Lokasi     string        `json:"lokasi" bson:"lokasi"`
	Departemen string        `json:"departemen" bson:"departemen"`
	User       string        `json:"user" bson:"user"`
	Ext        string        `json:"ext" bson:"ext"`
	Merk       string        `json:"merk" bson:"merk"`
	Tipe       string        `json:"tipe" bson:"tipe"`
	Inserted   doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool          `json:"-" bson:"is_deleted"`
}

type TeleponCollRepository struct {
	coll *mongo.Collection
}

func NewTeleponRepository(db *mongo.Database) *TeleponCollRepository {
	return &TeleponCollRepository{
		coll: db.Collection("telepons"),
	}
}

func (r *TeleponCollRepository) FindAll() (*[]Telepon, error) {
	var telepons []Telepon
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &telepons)
	if err != nil {
		return nil, err
	}
	return &telepons, nil
}

func (r *TeleponCollRepository) FindOneByID(id bson.ObjectID) (*Telepon, error) {
	var telepon Telepon
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&telepon)
	if err != nil {
		return nil, err
	}
	return &telepon, nil
}

func (r *TeleponCollRepository) InsertOne(telepon *Telepon) error {
	_, err := r.coll.InsertOne(context.TODO(), telepon)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeleponCollRepository) InsertMany(telepons []Telepon) error {
	_, err := r.coll.InsertMany(context.TODO(), telepons)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeleponCollRepository) UpdateOneByID(id bson.ObjectID, telepon *Telepon) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": telepon,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TeleponCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TeleponCollRepository) DeleteOneByID(id bson.ObjectID) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
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
