package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type TOA struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	Nama      string        `json:"nama" bson:"nama"`
	Lokasi    string        `json:"lokasi" bson:"lokasi"`
	Kode      string        `json:"kode" bson:"kode"`
	Posisi    string        `json:"posisi" bson:"posisi"`
	Inserted  doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated   *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted bool          `json:"-" bson:"is_deleted"`
}

type TOACollRepository struct {
	coll *mongo.Collection
}

func NewTOARepository(db *mongo.Database) *TOACollRepository {
	return &TOACollRepository{
		coll: db.Collection("toas"),
	}
}

func (r *TOACollRepository) FindAll(cq *util.CommonQuery) (*[]TOA, error) {
	var toas []TOA
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

	err = cur.All(context.TODO(), &toas)
	if err != nil {
		return nil, err
	}
	if toas == nil {
		return &[]TOA{}, nil
	}
	return &toas, nil
}

func (r *TOACollRepository) FindOneByID(id bson.ObjectID) (*TOA, error) {
	var toa TOA
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&toa)
	if err != nil {
		return nil, err
	}
	return &toa, nil
}

func (r *TOACollRepository) InsertOne(toa *TOA) error {
	_, err := r.coll.InsertOne(context.TODO(), toa)
	if err != nil {
		return err
	}
	return nil
}

func (r *TOACollRepository) InsertMany(toas []TOA) error {
	_, err := r.coll.InsertMany(context.TODO(), toas)
	if err != nil {
		return err
	}
	return nil
}

func (r *TOACollRepository) UpdateOneByID(id bson.ObjectID, toa *TOA) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": toa,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *TOACollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *TOACollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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

func (r *TOACollRepository) DeleteOneByID(id bson.ObjectID) error {
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
