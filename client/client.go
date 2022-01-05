package main

import (
	"context"
	"fmt"

	"github.com/machinebox/graphql"
)

// TodoStatusClearedInput
type TodoStatusClearedInput struct {
	TodoID []*string
}

// UpdateTodoStatusCleared updates status of transactions to cleared
func UpdateTodoStatusCleared(ctx context.Context, input TodoStatusClearedInput) (string, error) {
	client := graphql.NewClient("http://localhost:8080/query")

	request := graphql.NewRequest(`
	mutation UpdateTodo ($id: [String!]!) {
		updateTodoStatus(todoId: $id, status: cleared)
	}
	`)
	request.Var("id", input.TodoID)
	request.Header.Set("user-agent", "my user agent")

	var responseData map[string]interface{}
	if err := client.Run(ctx, request, &responseData); err != nil {
		return "", err
	}
	return responseData["updateTodoStatus"].(string), nil
}

func main() {
	ctx := context.Background()
	var t1 = "todos_01F2QQ3XZDZ0ND8EDZKP15N6ST"
	var t2 = "todos_01F2QQ40N9GPC3CA55KVTDV17W"
	var input = TodoStatusClearedInput{TodoID: []*string{&t1, &t2}}

	s, err := UpdateTodoStatusCleared(ctx, input)
	fmt.Println(s)
	fmt.Println(err)
}
