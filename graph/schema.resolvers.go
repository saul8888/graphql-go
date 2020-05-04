package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go-graphql/database"
	"go-graphql/graph/generated"
	"go-graphql/graph/model"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) InsertUser(ctx context.Context, input model.NewUser) (*model.User, error) {
	user := &model.User{
		ID:           primitive.NewObjectID(),
		CustomerName: input.CustomerName,
		Email:        input.Email,
		PhoneNumber:  input.PhoneNumber,
		Password:     input.Password,
		CreatedAt:    time.Now(),
		UpdateAt:     time.Now(),
	}
	//r.users = append(r.users, user)
	data.Insert(user)
	return user, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, id string, input model.UpdateUser) (*model.User, error) {
	customerID := id
	user := &model.UpdateUser{
		Email:       input.Email,
		PhoneNumber: input.PhoneNumber,
		Password:    input.Password,
		//UpdateAt:    time.Now(),
	}
	_, err := data.Update(customerID, user)
	if err != nil {
		log.Fatal(err)
	}
	//return c.JSON(http.StatusOK, updateCustomer)
	row, err := data.GetById(customerID)
	if err != nil {
		log.Fatal(err)
	}
	customer := &model.User{}
	for row.Next(context.TODO()) {
		row.Decode(customer)
	}
	return customer, nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, id string) (bool, error) {
	customerID := id
	delete, err := data.Delete(customerID)
	if err != nil {
		log.Fatal(err)
	}
	result := false
	if delete == 1 {
		result = true
	}
	return result, nil
}

func (r *queryResolver) GetUser(ctx context.Context, id string) (*model.User, error) {
	customerID := id
	row, err := data.GetById(customerID)
	if err != nil {
		log.Fatal(err)
	}
	customer := &model.User{}
	//var customer *model.User
	for row.Next(context.TODO()) {
		row.Decode(customer)
	}
	return customer, nil
}

func (r *queryResolver) GetTotalUser(ctx context.Context, input model.SeeData) (*model.Customers, error) {
	params := new(database.GetCustomersRequest)
	params.Limit = input.Limit
	params.Offset = input.Offset
	row, err := data.GetTotal(params)
	if err != nil {
		log.Fatal(err)
	}
	var users []*model.User
	for row.Next(context.TODO()) {
		customer := &model.User{}
		row.Decode(&customer)
		users = append(users, customer)
	}

	totalCustomers, err := data.GetCantTotal()
	if err != nil {
		log.Fatal(err)
	}
	customers := &model.Customers{}
	customers.Customers = users
	customers.Cant = totalCustomers
	return customers, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
var dbconection, _ = database.ConnectDB()
var data = database.NewDataBase(dbconection)
