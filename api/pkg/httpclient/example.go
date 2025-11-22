package httpclient

import (
	"context"
	"fmt"
)

func Example() {
	todo := New(Config{Host: "localhost", Port: "8080"})

	ctx := context.Background()

	id, err := todo.Create(ctx, "My first todo", "This is my first todo description")
	if err != nil {
		panic(err)
	}

	fmt.Println("Created todo with ID:", id)

	t, err := todo.Get(ctx, id.String())
	if err != nil {
		panic(err)
	}

	fmt.Printf("Fetched todo: %+v\n", t)

	err = todo.Delete(ctx, id.String())
	if err != nil {
		panic(err)
	}

	_, err = todo.Get(ctx, id.String())

	fmt.Println("Error fetching deleted todo:", err)
}
