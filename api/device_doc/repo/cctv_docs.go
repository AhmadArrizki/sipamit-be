package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type CCTVDoc struct {
	ID         bson.ObjectID  `json:"_id" bson:"_id"`
	Nama       string         `json:"nama" bson:"nama"`
	Lokasi     string         `json:"lokasi" bson:"lokasi"`
	Kode       string         `json:"kode" bson:"kode"`
	Checkpoint []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted   doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool           `json:"-" bson:"is_deleted"`
}

type CCTVDocCollRepository struct {
	coll *mongo.Collection
}

func NewCCTVDocRepository(db *mongo.Database) *CCTVDocCollRepository {
	return &CCTVDocCollRepository{
		coll: db.Collection("cctv_docs"),
	}
}

func (r *CCTVDocCollRepository) FindAll(cq *util.CommonQuery) (*[]CCTVDoc, error) {
	var cctvDocs []CCTVDoc
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

	err = cur.All(context.TODO(), &cctvDocs)
	if err != nil {
		return nil, err
	}
	if cctvDocs == nil {
		return &[]CCTVDoc{}, nil
	}
	return &cctvDocs, nil
}

func (r *CCTVDocCollRepository) FindOneByID(id bson.ObjectID) (*CCTVDoc, error) {
	var cctvDoc CCTVDoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&cctvDoc)
	if err != nil {
		return nil, err
	}
	return &cctvDoc, nil
}

func (r *CCTVDocCollRepository) InsertOne(cctvDoc *CCTVDoc) error {
	_, err := r.coll.InsertOne(context.TODO(), cctvDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *CCTVDocCollRepository) UpdateOneByID(id bson.ObjectID, cctvDoc *CCTVDoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": cctvDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *CCTVDocCollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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

func (r *CCTVDocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
