package interfaces

import (
	"strconv"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/go-redis/redis/v7"
	uuid "github.com/satori/go.uuid"
)

// A JWTRepository is a repository for jwt.
type JWTRepository struct {
	ConnRedis *redis.Client
}

// FindIDByAccessUUID returns the entity identified by the given access uuid.
func (jr *JWTRepository) FindIDByAccessUUID(au string) (int, error) {
	idStr, err := jr.ConnRedis.Get(au).Result()
	if err != nil {
		return 0, err
	}

	var id int
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// FindIDByRefreshUUID returns the entity identified by the given refresh uuid.
func (jr *JWTRepository) FindIDByRefreshUUID(au string) (int, error) {
	idStr, err := jr.ConnRedis.Get(au).Result()
	if err != nil {
		return 0, err
	}

	var id int
	id, err = strconv.Atoi(idStr)
	if err != nil {
		return 0, err
	}
	return id, nil
}

// SaveID saves the given entity by id
func (jr *JWTRepository) SaveID(id int) (domain.JWT, error) {
	j := domain.JWT{
		AccessUUID:  uuid.NewV4().String(),
		RefreshUUID: uuid.NewV4().String(),
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	var err error
	if j.AccessToken, err = j.CreateAccessToken(id); err != nil {
		return j, err
	}

	if j.RefreshToken, err = j.CreateRefreshToken(id); err != nil {
		return j, err
	}

	now := time.Now()

	if err = jr.ConnRedis.Set(j.AccessUUID, strconv.Itoa(id), time.Unix(j.AtExpires, 0).Sub(now)).Err(); err != nil {
		return j, err
	}

	if err = jr.ConnRedis.Set(j.RefreshUUID, strconv.Itoa(id), time.Unix(j.RtExpires, 0).Sub(now)).Err(); err != nil {
		return j, err
	}

	return j, err
}

// DeleteByAccessUUID deletes the entity by access uuid.
func (jr *JWTRepository) DeleteByAccessUUID(au string) (int64, error) {
	deleted, err := jr.ConnRedis.Del(au).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

// DeleteByRefreshUUID deletes the entity by access uuid.
func (jr *JWTRepository) DeleteByRefreshUUID(au string) (int64, error) {
	deleted, err := jr.ConnRedis.Del(au).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
