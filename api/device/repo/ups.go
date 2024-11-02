package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
)

type UPS struct {
	ID         bson.ObjectID `json:"_id" bson:"_id"`
	Nama       string        `json:"nama" bson:"nama"`
	Departemen string        `json:"departemen" bson:"departemen"`
	Tipe       string        `json:"tipe" bson:"tipe"`
	NoSeri     string        `json:"no_seri" bson:"no_seri"`
	Lokasi     string        `json:"lokasi" bson:"lokasi"`
	Inserted   doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated    *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted  bool          `json:"-" bson:"is_deleted"`
}

type UPSCollRepository struct {
	coll *mongo.Collection
}

func NewUPSRepository(db *mongo.Database) *UPSCollRepository {
	return &UPSCollRepository{
		coll: db.Collection("ups"),
	}
}

func (r *UPSCollRepository) FindAll(cq *util.CommonQuery) (*[]UPS, error) {
	var ups []UPS
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

	err = cur.All(context.TODO(), &ups)
	if err != nil {
		return nil, err
	}
	if ups == nil {
		return &[]UPS{}, nil
	}
	return &ups, nil
}

func (r *UPSCollRepository) FindOneByID(id bson.ObjectID) (*UPS, error) {
	var ups UPS
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&ups)
	if err != nil {
		return nil, err
	}
	return &ups, nil
}

func (r *UPSCollRepository) InsertOne(ups *UPS) error {
	_, err := r.coll.InsertOne(context.TODO(), ups)
	if err != nil {
		return err
	}
	return nil
}

func (r *UPSCollRepository) InsertMany(ups []UPS) error {
	_, err := r.coll.InsertMany(context.TODO(), ups)
	if err != nil {
		return err
	}
	return nil
}

func (r *UPSCollRepository) UpdateOneByID(id bson.ObjectID, ups *UPS) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": ups,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *UPSCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *UPSCollRepository) CountQuery(cq *util.CommonQuery) (int64, error) {
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

func (r *UPSCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
