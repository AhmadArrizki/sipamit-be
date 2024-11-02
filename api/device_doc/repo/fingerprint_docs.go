package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
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

func (r *FingerprintDocCollRepository) FindAll(cq *util.CommonQuery) (*[]FingerprintDoc, error) {
	var fpDocs []FingerprintDoc
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

	err = cur.All(context.TODO(), &fpDocs)
	if err != nil {
		return nil, err
	}
	if fpDocs == nil {
		return &[]FingerprintDoc{}, nil
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

func (r *FingerprintDocCollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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
