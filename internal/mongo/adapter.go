package mongo

import (
	"context"
	"errors"
	"github.com/my-otp/otp-rest/internal/config"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// ctx returns a context.Context.
func ctx() context.Context {
	return context.Background()
}

type Service struct {
	Secret string `bson:"secret"`
}

type Adapter struct {
	db *mongo.Database

	services *mongo.Collection
}

func NewAdapter(config config.DatabaseConfig) *Adapter {
	adapter := &Adapter{}

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(config.URL).SetServerAPIOptions(serverAPI))
	if err != nil {
		panic(err)
	}
	adapter.db = client.Database(config.Database)

	adapter.services = adapter.db.Collection("services")
	return adapter
}

func (a *Adapter) LoadSecret(service string) (string, error) {
	filter := bson.M{"name": bson.M{"$eq": service}}
	v, err := a.loadService(filter)
	if err != nil {
		return "", err
	}

	return v.Secret, nil
}
func (a *Adapter) SaveSecret(service string, secret string) error {
	filter := bson.M{"name": bson.M{"$eq": service}}

	v, err := a.loadService(filter)
	if err != nil && !errors.Is(err, mongo.ErrNoDocuments) {
		return err
	}
	v.Secret = secret
	update := bson.M{"$set": v}

	return a.updateService(filter, update)
}

func (a *Adapter) loadService(filter any) (Service, error) {
	result := a.services.FindOne(ctx(), filter)

	if result.Err() != nil {
		return Service{}, result.Err()
	}

	v := Service{}
	err := result.Decode(&v)
	return v, err
}

func (a *Adapter) updateService(filter any, update any) error {
	_, err := a.services.UpdateOne(ctx(), filter, update)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			_, err = a.services.InsertOne(ctx(), update)
			return err
		}
		return err
	}
	return nil
}
