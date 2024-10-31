package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type PrinterDoc struct {
	ID          bson.ObjectID  `json:"_id" bson:"_id"`
	Nama        string         `json:"nama" bson:"nama"`
	Departemen  string         `json:"departemen" bson:"departemen"`
	TipePrinter string         `json:"tipe_printer" bson:"tipe_printer"`
	NoSeri      string         `json:"no_seri" bson:"no_seri"`
	Checkpoint  []doc.CPDetail `json:"checkpoint" bson:"checkpoint"`
	Inserted    doc.ByAt       `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated     *doc.ByAt      `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted   bool           `json:"-" bson:"is_deleted"`
}

type PrinterDocCollRepository struct {
	coll *mongo.Collection
}

func NewPrinterDocRepository(db *mongo.Database) *PrinterDocCollRepository {
	return &PrinterDocCollRepository{
		coll: db.Collection("printer_docs"),
	}
}

func (r *PrinterDocCollRepository) FindAll() (*[]PrinterDoc, error) {
	var printerDocs []PrinterDoc
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &printerDocs)
	if err != nil {
		return nil, err
	}
	return &printerDocs, nil
}

func (r *PrinterDocCollRepository) FindOneByID(id bson.ObjectID) (*PrinterDoc, error) {
	var printerDoc PrinterDoc
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&printerDoc)
	if err != nil {
		return nil, err
	}
	return &printerDoc, nil
}

func (r *PrinterDocCollRepository) InsertOne(printerDoc *PrinterDoc) error {
	_, err := r.coll.InsertOne(context.TODO(), printerDoc)
	if err != nil {
		return err
	}
	return nil
}

func (r *PrinterDocCollRepository) UpdateOneByID(id bson.ObjectID, printerDoc *PrinterDoc) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	update := bson.M{
		"$set": printerDoc,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *PrinterDocCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
