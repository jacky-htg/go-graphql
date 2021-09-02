package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"first-go-grapql/graph/generated"
	"first-go-grapql/internal/model"
)

func (r *mutationResolver) MyServiceCreateTodo(ctx context.Context, input model.NewTodoInput) (*model.Todo, error) {
	return &model.Todo{ID: "1", Text: input.Text, User: &model.User{ID: input.UserID, Name: "Rijal"}, Done: false}, nil
}

func (r *mutationResolver) BargingCreateOrder(ctx context.Context, input model.NewOrderInput) (*model.Order, error) {
	return &model.Order{}, nil
}

func (r *queryResolver) MyServiceTodos(ctx context.Context) ([]*model.Todo, error) {
	return []*model.Todo{&model.Todo{ID: "1", Text: "Baca Buku", User: &model.User{ID: "1", Name: "Rijal"}, Done: true}}, nil
}

func (r *queryResolver) BargingOrders(ctx context.Context) ([]*model.Order, error) {
	return []*model.Order{&model.Order{ID: "1"}}, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
