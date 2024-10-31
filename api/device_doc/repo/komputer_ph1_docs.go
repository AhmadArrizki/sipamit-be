package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type KomputerPH1Doc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Nama       string         `json:"nama" bson:"nama"`
	Merk       string         `json:"merk" bson:"merk"`
	PC         string         `json:"pc" bson:"pc"`
	Monitor    string         `json:"monitor" bson:"monitor"`
	CPU        string         `json:"cpu" bson:"cpu"`
	RAM        string         `json:"ram" bson:"ram"`
	Internal   string         `json:"internal" bson:"internal"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type KomputerPH1DocCollRepository struct {
	coll *mongo.Collection
}

func NewKomputerPH1DocRepository(db *mongo.Database) *KomputerPH1DocCollRepository {
	return &KomputerPH1DocCollRepository{
		coll: db.Collection("komputer_ph1_docs"),
	}
}

func (r *KomputerPH1DocCollRepository) FindAll() (*[]KomputerPH1Doc, error) {
	var kph1Doc []KomputerPH1Doc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &kph1Doc)
	if err != nil {
		return nil, err
	}
	return &kph1Doc, nil
}

func (r *KomputerPH1DocCollRepository) FindOneByID(id bson.ObjectID) (*KomputerPH1Doc, error) {
	var kph1Doc KomputerPH1Doc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&kph1Doc)
	if err != nil {
		return nil, err
	}
	return &kph1Doc, nil
}

func (r *KomputerPH1DocCollRepository) InsertOne(kph1Doc *KomputerPH1Doc) error {
	_, err := r.coll.InsertOne(context.TODO(), kph1Doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH1DocCollRepository) UpdateOneByID(id bson.ObjectID, kph1Doc *KomputerPH1Doc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": kph1Doc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH1DocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
