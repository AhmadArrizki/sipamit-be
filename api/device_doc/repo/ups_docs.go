package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type UPSDoc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Nama       string         `json:"nama" bson:"nama"`
	Departemen string         `json:"departemen" bson:"departemen"`
	Tipe       string         `json:"tipe" bson:"tipe"`
	NoSeri     string         `json:"no_seri" bson:"no_seri"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type UPSDocCollRepository struct {
	coll *mongo.Collection
}

func NewUPSDocRepository(db *mongo.Database) *UPSDocCollRepository {
	return &UPSDocCollRepository{
		coll: db.Collection("ups_docs"),
	}
}

func (r *UPSDocCollRepository) FindAll() (*[]UPSDoc, error) {
	var upsDocs []UPSDoc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &upsDocs)
	if err != nil {
		return nil, err
	}
	return &upsDocs, nil
}

func (r *UPSDocCollRepository) FindOneByID(id bson.ObjectID) (*UPSDoc, error) {
	var upsDoc UPSDoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&upsDoc)
	if err != nil {
		return nil, err
	}
	return &upsDoc, nil
}

func (r *UPSDocCollRepository) InsertOne(upsDoc *UPSDoc) error {
	_, err := r.coll.InsertOne(context.TODO(), upsDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *UPSDocCollRepository) UpdateOneByID(id bson.ObjectID, upsDoc *UPSDoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": upsDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *UPSDocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
