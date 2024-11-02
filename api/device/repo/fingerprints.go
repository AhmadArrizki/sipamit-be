package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type FingerPrint struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	Nama      string        `json:"nama" bson:"nama"`
	Lokasi    string        `json:"lokasi" bson:"lokasi"`
	Kode      string        `json:"kode" bson:"kode"`
	Inserted  doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated   *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted bool          `json:"-" bson:"is_deleted"`
}

type FingerPrintCollRepository struct {
	coll *mongo.Collection
}

func NewFingerPrintRepository(db *mongo.Database) *FingerPrintCollRepository {
	return &FingerPrintCollRepository{
		coll: db.Collection("fingerprints"),
	}
}

func (r *FingerPrintCollRepository) FindAll(cq *util.CommonQuery) (*[]FingerPrint, error) {
	var fps []FingerPrint
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

	err = cur.All(context.TODO(), &fps)
	if err != nil {
		return nil, err
	}
	if fps == nil {
		return &[]FingerPrint{}, nil
	}
	return &fps, nil
}

func (r *FingerPrintCollRepository) FindOneByID(id bson.ObjectID) (*FingerPrint, error) {
	var fps FingerPrint
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&fps)
	if err != nil {
		return nil, err
	}
	return &fps, nil
}

func (r *FingerPrintCollRepository) InsertOne(fp *FingerPrint) error {
	_, err := r.coll.InsertOne(context.TODO(), fp)
	if err != nil {
		return err
	}
	return nil
}

func (r *FingerPrintCollRepository) InsertMany(fps []FingerPrint) error {
	_, err := r.coll.InsertMany(context.TODO(), fps)
	if err != nil {
		return err
	}
	return nil
}

func (r *FingerPrintCollRepository) UpdateOneByID(id bson.ObjectID, fp *FingerPrint) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": fp,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *FingerPrintCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *FingerPrintCollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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

func (r *FingerPrintCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
