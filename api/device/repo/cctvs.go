package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type CCTV struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	Nama      string        `json:"nama" bson:"nama"`
	Lokasi    string        `json:"lokasi" bson:"lokasi"`
	Kode      string        `json:"kode" bson:"kode"`
	Inserted  doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated   *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted bool          `json:"-" bson:"is_deleted"`
}

type CCTVCollRepository struct {
	coll *mongo.Collection
}

func NewCCTVRepository(db *mongo.Database) *CCTVCollRepository {
	return &CCTVCollRepository{
		coll: db.Collection("cctvs"),
	}
}

func (r *CCTVCollRepository) FindAll(cq *util.CommonQuery) (*[]CCTV, error) {
	var cctvs []CCTV
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

	err = cur.All(context.TODO(), &cctvs)
	if err != nil {
		return nil, err
	}
	if cctvs == nil {
		return &[]CCTV{}, nil
	}
	return &cctvs, nil
}

func (r *CCTVCollRepository) FindOneByID(id bson.ObjectID) (*CCTV, error) {
	var cctv CCTV
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&cctv)
	if err != nil {
		return nil, err
	}
	return &cctv, nil
}

func (r *CCTVCollRepository) InsertOne(cctv *CCTV) error {
	_, err := r.coll.InsertOne(context.TODO(), cctv)
	if err != nil {
		return err
	}
	return nil
}

func (r *CCTVCollRepository) InsertMany(cctvs []CCTV) error {
	_, err := r.coll.InsertMany(context.TODO(), cctvs)
	if err != nil {
		return err
	}
	return nil
}

func (r *CCTVCollRepository) UpdateOneByID(id bson.ObjectID, cctv *CCTV) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": cctv,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *CCTVCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *CCTVCollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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

func (r *CCTVCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
