package api

import (
	"BM8/lv2/models"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func SetUserInfo(id uint, name string, email string) error {
	user := models.User{
		ID:    id,
		Name:  name,
		Email: email,
	}
	models.Db.Create(&user)

	SetUserCache(&user)
	return nil
}
func GetUserInfo(id int) (*models.User, error) {
	// 尝试从缓存获取用户信息
	user, err := GetUserFromCache(id)
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	// 如果缓存中没有，从数据库获取
	user, err = GetUserFromDB(id)
	if err != nil {
		return nil, err
	}

	// 缓存用户信息
	err = SetUserCache(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func UpdateUserInfo(id int, newName string, newEmail string) error {
	user := models.User{}
	if err := models.Db.First(&user, id).Error; err != nil {
		return fmt.Errorf("user not found: %w", err)
	}

	user.Name = newName
	user.Email = newEmail
	if err := models.Db.Save(&user).Error; err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	cacheKey := fmt.Sprintf("user:%d", id)
	if err := models.Rdb.Del(models.Ctx, cacheKey).Err(); err != nil {
		return fmt.Errorf("failed to delete cache: %w", err)
	}
	
	return nil
}

func GetUserFromDB(id int) (*models.User, error) {
	var user models.User
	if err := models.Db.First(&user, id).Error; err != nil {
		return nil, err
	}
	fmt.Println("从mysql中获取")
	return &user, nil
}

func GetUserFromCache(id int) (*models.User, error) {
	val, err := models.Rdb.Get(models.Ctx, fmt.Sprintf("user:%d", id)).Result()
	if err == redis.Nil {
		return nil, nil
	} else if err != nil {
		return nil, err
	}

	var user models.User
	err = json.Unmarshal([]byte(val), &user)
	if err != nil {
		return nil, err
	}
	fmt.Println("从redis中获取")
	return &user, nil
}

func SetUserCache(user *models.User) error {
	data, err := json.Marshal(user)
	if err != nil {
		return err
	}

	return models.Rdb.Set(models.Ctx, fmt.Sprintf("user:%d", user.ID), data, 100*time.Second).Err()
}
