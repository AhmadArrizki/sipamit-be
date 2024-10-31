package repo

import (
	"context"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/internal/pkg/doc"
)

type User struct {
	ID        bson.ObjectID `json:"_id" bson:"_id"`
	FullName  string        `json:"full_name" bson:"full_name"`
	Username  string        `json:"username" bson:"username"`
	Password  string        `json:"-" bson:"password"`
	Role      string        `json:"role" bson:"role"`
	Inserted  doc.ByAt      `json:"inserted,omitempty" bson:"inserted,omitempty"`
	Updated   *doc.ByAt     `json:"updated,omitempty" bson:"updated,omitempty"`
	IsDeleted bool          `json:"-" bson:"is_deleted"`
}

type UserCollRepository struct {
	coll *mongo.Collection
}

func NewUserRepository(db *mongo.Database) *UserCollRepository {
	return &UserCollRepository{
		coll: db.Collection("users"),
	}
}

func (r *UserCollRepository) FindAll() (*[]User, error) {
	var users []User
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	cur, err := r.coll.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cur.Close(context.Background())

	err = cur.All(context.Background(), &users)
	if err != nil {
		return nil, err
	}
	return &users, nil
}

func (r *UserCollRepository) FindByID(_id bson.ObjectID) (*User, error) {
	var user *User
	filter := bson.M{
		"_id":        _id,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserCollRepository) FindByUsername(username string) (*User, error) {
	var user *User
	filter := bson.M{
		"username":   username,
		"is_deleted": bson.M{"$ne": true},
	}

	err := r.coll.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (r *UserCollRepository) InsertOne(user *User) error {
	_, err := r.coll.InsertOne(context.TODO(), user)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserCollRepository) UpdateOne(user *User) error {
	filter := bson.M{
		"_id":        user.ID,
		"is_deleted": bson.M{"$ne": true},
	}
	update := bson.M{
		"$set": user,
	}

	_, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	return nil
}

func (r *UserCollRepository) Count() (int64, error) {
	filter := bson.M{
		"is_deleted": bson.M{"$ne": true},
	}

	count, err := r.coll.CountDocuments(context.TODO(), filter)
	if err != nil {
		return 0, err
	}
	return count, nil
}
