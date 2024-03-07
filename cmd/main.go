package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/usawyer/urlShortener/internal/handlers"
	"github.com/usawyer/urlShortener/internal/routes"
	"github.com/usawyer/urlShortener/internal/service"

	"log"
)

func main() {
	srvc := service.New(5)
	handler := handlers.New(srvc)

	app := fiber.New()
	app.Use(logger.New())
	routes.InitRoutes(app, handler)
	log.Fatal(app.Listen(":8080"))
}

//import (
//	"context"
//	"fmt"
//	"github.com/go-gorm/caches/v4"
//	"github.com/redis/go-redis/v9"
//	"gorm.io/driver/postgres"
//	"os"
//	"time"
//
//	"gorm.io/gorm"
//)
//
//type UserRoleModel struct {
//	gorm.Model
//	Name string `gorm:"unique"`
//}
//
//type UserModel struct {
//	gorm.Model
//	Name   string
//	RoleId uint
//	Role   *UserRoleModel `gorm:"foreignKey:role_id;references:id"`
//}
//
//type redisCacher struct {
//	rdb *redis.Client
//}
//
//func (c *redisCacher) Get(ctx context.Context, key string, q *caches.Query[any]) (*caches.Query[any], error) {
//	res, err := c.rdb.Get(ctx, key).Result()
//	if err == redis.Nil {
//		return nil, nil
//	}
//
//	if err != nil {
//		return nil, err
//	}
//
//	if err := q.Unmarshal([]byte(res)); err != nil {
//		return nil, err
//	}
//
//	return q, nil
//}
//
//func (c *redisCacher) Store(ctx context.Context, key string, val *caches.Query[any]) error {
//	res, err := val.Marshal()
//	if err != nil {
//		return err
//	}
//
//	c.rdb.Set(ctx, key, res, 30*time.Second) // Set proper cache time
//	return nil
//}
//
//func (c *redisCacher) Invalidate(ctx context.Context) error {
//	var (
//		cursor uint64
//		keys   []string
//	)
//	for {
//		var (
//			k   []string
//			err error
//		)
//		k, cursor, err = c.rdb.Scan(ctx, cursor, fmt.Sprintf("%s*", caches.IdentifierPrefix), 0).Result()
//		if err != nil {
//			return err
//		}
//		keys = append(keys, k...)
//		if cursor == 0 {
//			break
//		}
//	}
//
//	if len(keys) > 0 {
//		if _, err := c.rdb.Del(ctx, keys...).Result(); err != nil {
//			return err
//		}
//	}
//	return nil
//}
//
//func getEnv(key, defaultValue string) string {
//	value := os.Getenv(key)
//	if value == "" {
//		return defaultValue
//	}
//	return value
//}
//
//func main() {
//	connectionParams := map[string]string{
//		"host":     getEnv("DB_HOST", "localhost"),
//		"user":     getEnv("POSTGRES_USER", "postgres"),
//		"password": getEnv("POSTGRES_PASSWORD", "postgres"),
//		"dbname":   getEnv("POSTGRES_DB", "team00"),
//		"port":     getEnv("DB_PORT", "5432"),
//		"sslmode":  "disable",
//		"TimeZone": "Asia/Novosibirsk",
//	}
//	var dsn string
//	for key, value := range connectionParams {
//		dsn += fmt.Sprintf("%s=%s ", key, value)
//	}
//
//	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
//
//	rdb := redis.NewClient(&redis.Options{
//		Addr:     "localhost:6379",
//		Password: "",
//		DB:       0,
//	})
//
//	cacher := &redisCacher{rdb: rdb}
//
//	cachesPlugin := &caches.Caches{Conf: &caches.Config{
//		Cacher: cacher,
//	}}
//
//	_ = db.Use(cachesPlugin)
//	//
//	//_ = db.AutoMigrate(&UserRoleModel{})
//	//_ = db.AutoMigrate(&UserModel{})
//	//
//	//db.Delete(&UserRoleModel{})
//	//db.Delete(&UserModel{})
//	//
//	//adminRole := &UserRoleModel{
//	//	Name: "Admin",
//	//}
//	//db.Save(adminRole)
//	//
//	//guestRole := &UserRoleModel{
//	//	Name: "Guest",
//	//}
//	//db.Save(guestRole)
//	//
//	//db.Save(&UserModel{
//	//	Name: "ktsivkov",
//	//	Role: adminRole,
//	//})
//	//
//	//db.Save(&UserModel{
//	//	Name: "anonymous",
//	//	Role: guestRole,
//	//})
//	//
//	//q1User := &UserModel{}
//	//db.WithContext(context.Background()).Find(q1User, "Name = ?", "ktsivkov")
//	//q2User := &UserModel{}
//	//db.WithContext(context.Background()).Find(q2User, "Name = ?", "ktsivkov")
//	//
//	//fmt.Println(fmt.Sprintf("%+v", q1User))
//	//fmt.Println(fmt.Sprintf("%+v", q2User))
//
//	fmt.Println("lol")
//
//}
