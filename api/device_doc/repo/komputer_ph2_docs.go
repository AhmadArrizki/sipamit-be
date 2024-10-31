package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type KomputerPH2Doc struct {
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

type KomputerPH2DocCollRepository struct {
	coll *mongo.Collection
}

func NewKomputerPH2DocRepository(db *mongo.Database) *KomputerPH2DocCollRepository {
	return &KomputerPH2DocCollRepository{
		coll: db.Collection("komputer_ph2_docs"),
	}
}

func (r *KomputerPH2DocCollRepository) FindAll() (*[]KomputerPH2Doc, error) {
	var kph2Doc []KomputerPH2Doc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &kph2Doc)
	if err != nil {
		return nil, err
	}
	return &kph2Doc, nil
}

func (r *KomputerPH2DocCollRepository) FindOneByID(id bson.ObjectID) (*KomputerPH2Doc, error) {
	var kph2Doc KomputerPH2Doc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&kph2Doc)
	if err != nil {
		return nil, err
	}
	return &kph2Doc, nil
}

func (r *KomputerPH2DocCollRepository) InsertOne(kph2Doc *KomputerPH2Doc) error {
	_, err := r.coll.InsertOne(context.TODO(), kph2Doc)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH2DocCollRepository) UpdateOneByID(id bson.ObjectID, kph2Doc *KomputerPH2Doc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": kph2Doc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH2DocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
