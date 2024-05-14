package customer

import (
	"context"
	"errors"

	"coffee-culture.uk/internal/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CustomerRepository interface {
	CreateCustomer(c context.Context, customer Customer) (Customer, error)
	// GetUserByID(c context.Context, id primitive.ObjectID) (*Customer, error)
	GetCustomerByEmail(c context.Context, email string) (*Customer, error)
	// GetCustomerByUsername(c context.Context, username string) (*Customer, error)
	EditCredentials(c context.Context, customer Customer) (*Customer, error)
	GetStampHistory(c context.Context, id string) ([]Stamp, error)
}

type repositoryImpl struct {
	collection *mongo.Collection
}

func NewCustomerRepository() CustomerRepository {
	collection := db.Client.Database(db.DATABASE_NAME).Collection(db.COLLECTION_CUSTOMER)
	return &repositoryImpl{collection: collection}
}

func (r *repositoryImpl) CreateCustomer(c context.Context, customer Customer) (Customer, error) {
	filter := bson.M{"email": customer.Email}

	var existingCustomer Customer
	err := r.collection.FindOne(c, filter).Decode(&existingCustomer)
	if err == nil {
		return Customer{}, errors.New("a user with this email already exists")
	} else if err != mongo.ErrNoDocuments {
		return Customer{}, err
	}

	result, err := r.collection.InsertOne(c, customer)
	if err != nil {
		return Customer{}, err
	}

	_, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return Customer{}, errors.New("failed to convert inserted ID to ObjectID")
	}

	return customer, nil
}

// EditCredentials implements CustomerRepository.
func (r *repositoryImpl) EditCredentials(c context.Context, customer Customer) (*Customer, error) {
	panic("unimplemented")
}

func (r *repositoryImpl) GetCustomerByEmail(c context.Context, email string) (*Customer, error) {
	filter := bson.M{"email": email}

	var customer Customer
	err := r.collection.FindOne(c, filter).Decode(&customer)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}

// GetStampHistory implements CustomerRepository.
func (r *repositoryImpl) GetStampHistory(c context.Context, id string) ([]Stamp, error) {
	panic("unimplemented")
}


