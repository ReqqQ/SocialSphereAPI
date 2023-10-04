package uicontrollers

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/go-redis/cache/v9"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"

	appuser "github.com/ReqqQ/SocialSphereAPI/src/app/users"
)

type ControllersInterface interface {
	InitGetRoutes(app *fiber.App)
}
type Controllers struct {
	appUser appuser.AppUser
}
type UsersList struct {
	Ids []string
}

type User struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	City string `json:"city"`
}

func InitDBData() {
	for {
		ctx := context.Background()

		client := redis.NewClient(
			&redis.Options{
				Addr:     "redis-11093.c304.europe-west1-2.gce.cloud.redislabs.com:11093",
				Password: "Y81lJ3sG7Cvv3YAV1k9FhS7qi3CrxYUy",
				DB:       0,
				PoolSize: 50,
			},
		)
		defer client.Close()

		pipe := client.Pipeline()
		usersIds := pipe.LRange(ctx, "users", int64(0), int64(-1))

		_, err := pipe.Exec(ctx)
		if err != nil {
			fmt.Println("Błąd podczas wykonania pipeline:", err)
		}

		rawUsers, err := usersIds.Result()
		if err != nil {
			fmt.Println("Błąd podczas odczytywania wyniku GET:", err)
		}
		usersList := UsersList{Ids: rawUsers}
		var usersPreparedList []*redis.StringCmd
		for _, userId := range usersList.Ids {
			usersPreparedList = append(usersPreparedList, pipe.Get(ctx, "user:"+userId))
		}
		pipe.Exec(ctx)
		cacheRedisClient := redis.NewClient(
			&redis.Options{
				Addr: "127.0.0.1:6379",
			},
		)
		cacheClient := cache.New(&cache.Options{Redis: cacheRedisClient})
		for _, rawData := range usersPreparedList {
			rawUser, _ := rawData.Result()

			var user User
			err = json.Unmarshal([]byte(rawUser), &user)
			if err != nil {
				fmt.Println("Błąd podczas deserializacji JSON:", err)
			}

			cacheClient.Set(
				&cache.Item{
					Ctx:   ctx,
					Key:   "user:" + user.Id,
					Value: user,
					TTL:   1800,
				},
			)
		}

		time.Sleep(30 * time.Second)
	}
}

func (co Controllers) InitGetRoutes(app *fiber.App) {
	v1 := app.Group(
		"/api/v1", func(c *fiber.Ctx) error {
			c.Set("Version", "v1")
			return c.Next()
		},
	)
	v1.Get(
		"/user/1", func(c *fiber.Ctx) error {
			co.appUser.GetUser(1)
			//redisClient := redis.NewClient(&redis.Options{
			//	Addr: "127.0.0.1:6379",
			//})
			//
			//cacheClient := cache.New(&cache.Options{Redis: redisClient})
			//
			//var userCache User
			//cacheClient.Get(context.Background(), "user:1", &userCache)

			return c.JSON("")
		},
	)
}
