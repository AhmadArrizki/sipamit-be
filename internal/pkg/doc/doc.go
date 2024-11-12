package doc

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
	"sipamit-be/api/app/repo"
	_db "sipamit-be/internal/db"
	"sipamit-be/internal/pkg/log"
	"time"
)

type ByAt struct {
	ID *bson.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	At time.Time      `json:"at" bson:"at"`
}

func (u *ByAt) MarshalJSON() ([]byte, error) {
	type Alias ByAt
	var fullName string

	if u.ID != nil {
		userRepo := repo.NewUserRepository(_db.Client)

		user, err := userRepo.FindByID(*u.ID)
		if err == nil {
			fullName = user.FullName
		}
	}

	return json.Marshal(&struct {
		*Alias
		FullName string `json:"full_name" bson:"full_name"`
	}{
		Alias:    (*Alias)(u),
		FullName: fullName,
	})
}

type CPDetail struct {
	Name       string `json:"name" bson:"name"`
	OK         bool   `json:"ok" bson:"ok"`
	Keterangan string `json:"keterangan" bson:"keterangan"`
}

type DeviceDocForm struct {
	DeviceID   string     `form:"device_id" json:"device_id"`
	Checkpoint []CPDetail `form:"checkpoint" json:"checkpoint"`

	DeviceOID bson.ObjectID `form:"-" json:"-"`
}

func NewDeviceDocForm(c echo.Context) (*DeviceDocForm, error) {
	f := new(DeviceDocForm)
	err := c.Bind(f)
	if err != nil {
		log.Errorf("Failed to bind doc form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if f.DeviceID == "" {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}

	f.DeviceOID, err = bson.ObjectIDFromHex(f.DeviceID)
	if err != nil {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid device id")
	}

	if len(f.Checkpoint) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}
	return f, nil
}

type UpdateDeviceDocForm struct {
	Checkpoint []CPDetail `form:"checkpoint" json:"checkpoint"`
}

func NewUpdateDeviceDocForm(c echo.Context) (*UpdateDeviceDocForm, error) {
	f := new(UpdateDeviceDocForm)
	err := c.Bind(f)
	if err != nil {
		log.Errorf("Failed to bind doc form: %v", err)
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Invalid request")
	}

	if len(f.Checkpoint) == 0 {
		return nil, echo.NewHTTPError(http.StatusBadRequest, "Please fill provided field")
	}
	return f, nil
}
