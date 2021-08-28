package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"github.com/duyledat197/netfix-backend/internal/delivery/grpc"
	httpS "github.com/duyledat197/netfix-backend/internal/delivery/http"
	"github.com/duyledat197/netfix-backend/internal/domain"
	"github.com/duyledat197/netfix-backend/internal/repository"
	"github.com/duyledat197/netfix-backend/pb"
	"github.com/duyledat197/netfix-backend/utils/storage"
	"github.com/duyledat197/netfix-backend/utils/token"

	mongoC "github.com/duyledat197/netfix-backend/internal/repository/mongo"
	"github.com/urfave/cli"
	"go.mongodb.org/mongo-driver/mongo"
	mgoOption "go.mongodb.org/mongo-driver/mongo/options"
)

var (
	//MongoCollectionUser collection of user
	MongoCollectionUser     = "users"
	MongoCollectionPost     = "posts"
	MongoCollectionCategory = "categories"
)

type srv struct {
	cfg *Config

	mgoClient *mongo.Client
	mgoDB     *mongo.Database

	bucket *storage.Bucket

	authenticator token.Authenticator

	// load store
	userRepository     repository.UserRepository
	postRepository     repository.PostRepository
	categoryRepository repository.CategoryRepository

	// load domain
	userDomain     domain.UserDomain
	authDomain     domain.AuthDomain
	postDomain     domain.PostDomain
	categoryDomain domain.CategoryDomain

	// load delivery
	authDelivery     pb.AuthServiceServer
	postDelivery     pb.PostServiceServer
	categoryDelivery pb.CategoryServiceServer

	// define for logger
}

func (s *srv) connectMongo() error {
	mgoClientOptions := mgoOption.Client().ApplyURI(s.cfg.Mongo.getConnectString())

	// Connect to MongoDB
	var err error
	mgoClient, err := mongo.Connect(context.TODO(), mgoClientOptions)
	if err != nil {
		return err
	}
	s.mgoDB = mgoClient.Database(s.cfg.Mongo.Database)
	log.Println("connect mongodb success")
	return nil
}
func (s *srv) loadRepository() error {
	s.userRepository = mongoC.NewUserRepository(s.mgoDB.Collection(MongoCollectionUser))
	s.postRepository = mongoC.NewPostRepository(s.mgoDB.Collection(MongoCollectionPost))
	s.categoryRepository = mongoC.NewCategoryRepository(s.mgoDB.Collection(MongoCollectionCategory))
	return nil
}

func (s *srv) loadDomain() error {
	s.userDomain = domain.NewUserDomain(s.userRepository)
	s.authDomain = domain.NewAuthDomain(s.userRepository, s.authenticator)
	s.postDomain = domain.NewPostDomain(s.postRepository)
	s.categoryDomain = domain.NewCategoryDomain(s.categoryRepository)
	return nil
}

func (s *srv) loadBucket() error {
	var err error
	s.bucket, err = storage.NewBucket("netflixgiare")
	return err
}

func (s *srv) loadHTTPServer() error {
	ctx := context.Background()
	s.authDelivery = grpc.NewAuthDelivery(s.authDomain)
	s.postDelivery = grpc.NewPostDelivery(s.postDomain)
	s.categoryDelivery = grpc.NewCategoryDelivery(s.categoryDomain)
	mux, err := httpS.NewHTTPHandler(ctx, s.authenticator, s.authDelivery, s.postDelivery, s.categoryDelivery, s.bucket)
	if err != nil {
		return err
	}
	log.Println("start http success in port", s.cfg.HTTP.Port)
	return http.ListenAndServe(":"+s.cfg.HTTP.Port, mux)
}

func (s *srv) loadLogger() error {
	return nil
}

func (s *srv) loadConfig(ctx *cli.Context) error {
	s.cfg = &Config{
		HTTP: ConnAddress{
			Host: ctx.GlobalString(HTTPHostFlag.GetName()),
			Port: ctx.GlobalString(HTTPPortFlag.GetName()),
		},
		Mongo: Mongo{
			Host:     ctx.GlobalString(MongoHostFlag.GetName()),
			Port:     ctx.GlobalString(MongoPortFlag.GetName()),
			Database: ctx.GlobalString(MongoDatabaseFlag.GetName()),
		},
		TokenKey:   ctx.GlobalString(TokenKeyFlag.GetName()),
		CookieName: ctx.GlobalString(CookieNameFlag.GetName()),
	}
	return nil
}

func (s *srv) loadAuthenticator() error {
	// jwt.SetInstance(s.cfg.TokenKey, s.cfg.TokenExpiration)
	var err error
	s.authenticator, err = token.NewPasetoAuthenticator(s.cfg.TokenKey, 15*24*time.Hour)
	return err
}
