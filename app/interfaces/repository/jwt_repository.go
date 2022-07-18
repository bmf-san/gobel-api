package repository

import (
	"context"
	"strconv"
	"time"

	"github.com/bmf-san/gobel-api/app/domain"
	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
)

// A JWTRepository is a repository for jwt.
type JWTRepository struct {
	ConnRedis *redis.Client
}

// FindIDByAccessUUID returns the entity identified by the given access uuid.
func (jr *JWTRepository) FindIDByAccessUUID(au string) (int, error) {
	ctx := context.Background()
	idStr, err := jr.ConnRedis.Get(ctx, au).Result()
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
func (jr *JWTRepository) FindIDByRefreshUUID(ru string) (int, error) {
	ctx := context.Background()
	idStr, err := jr.ConnRedis.Get(ctx, ru).Result()
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
	auuid, err := uuid.NewRandom()
	if err != nil {
		return domain.JWT{}, err
	}
	ruuid, err := uuid.NewRandom()
	if err != nil {
		return domain.JWT{}, err
	}
	j := domain.JWT{
		AccessUUID:  auuid.String(),
		RefreshUUID: ruuid.String(),
		AtExpires:   time.Now().Add(time.Minute * 15).Unix(),
		RtExpires:   time.Now().Add(time.Hour * 24 * 7).Unix(),
	}

	if j.AccessToken, err = j.CreateAccessToken(id); err != nil {
		return j, err
	}

	if j.RefreshToken, err = j.CreateRefreshToken(id); err != nil {
		return j, err
	}

	now := time.Now()

	ctx := context.Background()
	if err = jr.ConnRedis.Set(ctx, j.AccessUUID, strconv.Itoa(id), time.Unix(j.AtExpires, 0).Sub(now)).Err(); err != nil {
		return j, err
	}

	if err = jr.ConnRedis.Set(ctx, j.RefreshUUID, strconv.Itoa(id), time.Unix(j.RtExpires, 0).Sub(now)).Err(); err != nil {
		return j, err
	}

	return j, err
}

// DeleteByAccessUUID deletes the entity by access uuid.
func (jr *JWTRepository) DeleteByAccessUUID(au string) (int64, error) {
	ctx := context.Background()
	deleted, err := jr.ConnRedis.Del(ctx, au).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}

// DeleteByRefreshUUID deletes the entity by access uuid.
func (jr *JWTRepository) DeleteByRefreshUUID(au string) (int64, error) {
	ctx := context.Background()
	deleted, err := jr.ConnRedis.Del(ctx, au).Result()
	if err != nil {
		return 0, err
	}

	return deleted, nil
}
