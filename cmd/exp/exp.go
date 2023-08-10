package main

import (
	"context"
	"fmt"
	"strings"
)

type ctxKey string

const (
	favoriteColorKey ctxKey = "favorite-color"
)

func main() {
	ctx := context.Background()

	ctx = context.WithValue(ctx, favoriteColorKey, "blue")

	value := ctx.Value(favoriteColorKey)
	strValue, ok := value.(string)
	if !ok {
		fmt.Println("failed to do type assertion")
		return
	}

	fmt.Println(strValue)
	fmt.Println(strings.HasPrefix(strValue, "b"))
}
