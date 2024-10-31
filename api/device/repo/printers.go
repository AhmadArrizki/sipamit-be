package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type Printer struct {
	ID          bson.ObjectID `json:"_id" bson:"_id"`
	Nama        string        `json:"nama" bson:"nama"`
	Departemen  string        `json:"departemen" bson:"departemen"`
	TipePrinter string        `json:"tipe_printer" bson:"tipe_printer"`
	NoSeri      string        `json:"no_seri" bson:"no_seri"`
	Inserted    doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated     *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted   bool          `json:"-" bson:"is_deleted"`
}

type PrinterCollRepository struct {
	coll *mongo.Collection
}

func NewPrinterRepository(db *mongo.Database) *PrinterCollRepository {
	return &PrinterCollRepository{
		coll: db.Collection("printers"),
	}
}

func (r *PrinterCollRepository) FindAll() (*[]Printer, error) {
	var printers []Printer
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.TODO())

	err = cur.All(context.TODO(), &printers)
	if err != nil {
		return nil, err
	}
	return &printers, nil
}

func (r *PrinterCollRepository) FindOneByID(id bson.ObjectID) (*Printer, error) {
	var printer Printer
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&printer)
	if err != nil {
		return nil, err
	}
	return &printer, nil
}

func (r *PrinterCollRepository) InsertOne(printer *Printer) error {
	_, err := r.coll.InsertOne(context.TODO(), printer)
	if err != nil {
		return err
	}
	return nil
}

func (r *PrinterCollRepository) InsertMany(printers []Printer) error {
	_, err := r.coll.InsertMany(context.TODO(), printers)
	if err != nil {
		return err
	}
	return nil
}

func (r *PrinterCollRepository) UpdateOneByID(id bson.ObjectID, printer *Printer) error {
	filter := bson.M{
		"_id":        id,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": printer,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *PrinterCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}
	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}

func (r *PrinterCollRepository) DeleteOneByID(id bson.ObjectID) error {
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
