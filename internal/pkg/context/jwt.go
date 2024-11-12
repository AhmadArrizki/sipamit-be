package context

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/v2/bson"
	"net/http"
	"sipamit-be/api/app/repo"
	"sipamit-be/internal/config"
	_db "sipamit-be/internal/db"
	"sipamit-be/internal/pkg/const"
	"sipamit-be/internal/pkg/doc"
	"sipamit-be/internal/pkg/log"
	"sipamit-be/internal/pkg/util"
	"strings"
	"sync"
	"time"
)

var onceUserRepo sync.Once
var userRepo *repo.UserCollRepository

type UserClaims struct {
	jwt.StandardClaims
	ID                 string        `json:"id"`
	IDAsObjectID       bson.ObjectID `json:"-"`
	Username           string        `json:"username"`
	Role               string        `json:"role"`
	ExpiredDateInMilis int64         `json:"expiredDateInMilis"`
}

func (u *UserClaims) IsSuperAdminOrAdmin() bool {
	return u.IsSuperAdmin() || u.IsAdmin()
}

func (u *UserClaims) IsSuperAdmin() bool {
	return u.Role == _const.SuperAdminRole
}

func (u *UserClaims) IsAdmin() bool {
	return u.Role == _const.AdminRole
}

func (u *UserClaims) ByAt() doc.ByAt {
	return doc.ByAt{
		ID: &u.IDAsObjectID,
		At: time.Now(),
	}
}

func (u *UserClaims) ByAtPtr() *doc.ByAt {
	return &doc.ByAt{
		ID: &u.IDAsObjectID,
		At: time.Now(),
	}
}

type Context struct {
	echo.Context
	Claims       *UserClaims
	loggedInUser *repo.User
}

func (c *Context) LoggedInUser() *repo.User {
	if c.loggedInUser == nil {
		onceUserRepo.Do(func() {
			userRepo = repo.NewUserRepository(_db.Client)
		})
		u, err := userRepo.FindByID(c.Claims.IDAsObjectID)
		if err != nil {
			log.Errorc(c, err)
		}
		c.loggedInUser = u
	}
	return c.loggedInUser
}

func Handler(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc, err := MakeContext(c)
		if err != nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		if nc.LoggedInUser() == nil {
			return echo.NewHTTPError(http.StatusUnauthorized, "Unauthorized")
		}
		return next(nc)
	}
}

func MakeContext(c echo.Context) (*Context, error) {
	claims, err := NewUserClaims(c)
	if err != nil {
		return nil, err
	}
	return &Context{c, claims, nil}, nil
}

func NewUserClaims(c echo.Context) (*UserClaims, error) {
	// get identity
	header := c.Request().Header.Get("Authorization")
	bearer := strings.Split(header, " ")
	if len(bearer) != 2 {
		return nil, echo.ErrUnauthorized
	}

	if bearer[0] != "Bearer" {
		return nil, echo.ErrUnauthorized
	}

	return NewUserClaimsFromString(bearer[1])
}

func NewUserClaimsFromString(s string) (*UserClaims, error) {
	cred := &UserClaims{}
	token, err := jwt.ParseWithClaims(s, cred, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(config.JWT.Key), nil
	})
	if err != nil {
		return nil, echo.ErrUnauthorized
	}
	if claims, ok := token.Claims.(*UserClaims); ok && token.Valid {
		//Tidak bisa akses jika jwt sudah kadaluarsa.
		if claims.ExpiredDateInMilis < util.TimeToMilis(time.Now()) {
			return nil, echo.ErrUnauthorized
		}

		IDAsObjectID, err := bson.ObjectIDFromHex(claims.ID)
		if err != nil {
			return nil, err
		}
		claims.IDAsObjectID = IDAsObjectID
		return claims, nil
	}
	return nil, echo.ErrUnauthorized
}

func SuperAdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc, ok := c.(*Context)
		if !ok {
			return echo.ErrUnauthorized
		}
		if nc.Claims.IsSuperAdmin() {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

func AdminOrSuperAdminOnly(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		nc, ok := c.(*Context)
		if !ok {
			return echo.ErrUnauthorized
		}
		if nc.Claims.IsSuperAdminOrAdmin() {
			return next(c)
		}
		return echo.ErrUnauthorized
	}
}

func MakeToken(u *repo.User) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["id"] = u.ID.Hex()
	claims["username"] = u.Username
	claims["role"] = u.Role
	claims["expiredDateInMilis"] = time.Now().AddDate(0, 0, config.JWT.Expire).Unix() * 1000

	accessToken, err := token.SignedString([]byte(config.JWT.Key))
	if err != nil {
		return "", echo.NewHTTPError(http.StatusInternalServerError, "Internal server exception: "+err.Error()).SetInternal(err)
	}
	return accessToken, nil
}
