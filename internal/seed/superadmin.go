package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/app/repo"
	"sipamit-be/internal/pkg/const"
	"sipamit-be/internal/pkg/util"
	"time"
)

var SuperAdminID bson.ObjectID

func SuperAdmin(db *mongo.Database) {
	userRepo := repo.NewUserRepository(db)

	count, _ := userRepo.Count()
	if count > 0 {
		log.Info("User already seeded")

		superadmin, err := userRepo.FindByUsername("superadmin")
		if err == nil {
			SuperAdminID = superadmin.ID
		}

		return
	}

	SuperAdminID = bson.NewObjectID()

	superadmin := &repo.User{
		ID:       SuperAdminID,
		FullName: "Super Admin",
		Username: "superadmin",
		Password: util.CryptPassword("superadmin"),
		Role:     _const.SuperAdminRole,
		Inserted: repo.ByAt{
			ID: &SuperAdminID,
			At: time.Now(),
		},
		IsDeleted: false,
	}

	err := userRepo.InsertOne(superadmin)
	if err != nil {
		log.Errorf("Failed to seed superadmin: %v", err)
	}
	log.Info("SuperAdmin seeded")
}
