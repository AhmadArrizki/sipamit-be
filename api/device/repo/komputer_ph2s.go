package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type KomputerPH2 struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	Nama      string        `json:"nama" bson:"nama"`
	Merk      string        `json:"merk" bson:"merk"`
	PC        string        `json:"pc" bson:"pc"`
	Monitor   string        `json:"monitor" bson:"monitor"`
	CPU       string        `json:"cpu" bson:"cpu"`
	RAM       string        `json:"ram" bson:"ram"`
	Internal  string        `json:"internal" bson:"internal"`
	Lokasi    string        `json:"lokasi" bson:"lokasi"`
	Inserted  doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated   *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted bool          `json:"-" bson:"is_deleted"`
}

type KomputerPH2CollRepository struct {
	coll *mongo.Collection
}

func NewKomputerPH2Repository(db *mongo.Database) *KomputerPH2CollRepository {
	return &KomputerPH2CollRepository{
		coll: db.Collection("komputer_ph2s"),
	}
}

func (r *KomputerPH2CollRepository) FindAll(cq *util.CommonQuery) (*[]KomputerPH2, error) {
	var kph2s []KomputerPH2
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	if len(cq.Q) > 0 {
		var pattern = bson.Regex{Pattern: cq.Q, Options: "i"}
		filter["nama"] = bson.M{"$regex": pattern}
	}

	findOptions, err := util.BuildPaginationAndOrderOptionByField(bson.M{"_id": cq.Sort}, cq.Page, cq.Limit)
	if err != nil {
		return nil, err
	}

	cur, err := r.coll.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &kph2s)
	if err != nil {
		return nil, err
	}
	if kph2s == nil {
		return &[]KomputerPH2{}, nil
	}
	return &kph2s, nil
}

func (r *KomputerPH2CollRepository) FindOneByID(id bson.ObjectID) (*KomputerPH2, error) {
	var kph2 KomputerPH2
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&kph2)
	if err != nil {
		return nil, err
	}
	return &kph2, nil
}

func (r *KomputerPH2CollRepository) InsertOne(kph2 *KomputerPH2) error {
	_, err := r.coll.InsertOne(context.TODO(), kph2)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH2CollRepository) InsertMany(kph2s []KomputerPH2) error {
	_, err := r.coll.InsertMany(context.TODO(), kph2s)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH2CollRepository) UpdateOneByID(id bson.ObjectID, kph2 *KomputerPH2) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": kph2,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *KomputerPH2CollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *KomputerPH2CollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	if len(cq.Q) > 0 {
		var pattern = bson.Regex{Pattern: cq.Q, Options: "i"}
		filter["nama"] = bson.M{"$regex": pattern}
	}

	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *KomputerPH2CollRepository) DeleteOneByID(id bson.ObjectID) error {
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
