package seed

import (
	"github.com/labstack/gommon/log"
	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"sipamit-be/api/app/repo"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/util"
	"time"
)

func SuperAdmin(db *mongo.Database) {
	userRepo := repo.NewUserRepository(db)

	count, _ := userRepo.Count()
	if count > 0 {
		log.Info("User already seeded")
		return
	}

	superadmin := &repo.User{
		ID:       bson.NewObjectID(),
		FullName: "Super Admin",
		Username: "superadmin",
		Password: util.CryptPassword("superadmin"),
		Role:     doc.SuperAdminRole,
		Inserted: doc.ByAt{
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
