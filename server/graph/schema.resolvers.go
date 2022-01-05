package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/siongui/go-graphql-server-client-example/server/graph/generated"
	"github.com/siongui/go-graphql-server-client-example/server/graph/model"
)

func (r *mutationResolver) UpdateTodoStatus(ctx context.Context, todoID []string, status model.TodoStatus) (string, error) {
	fmt.Println(todoID)
	fmt.Println(status)

	return "OK", nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
