package main

import (
	"context"
	"log"
)

func main() {
	ctx := context.TODO()

	md := &DB{}
	if err := md.Connect(ctx); err != nil {
		log.Fatal(err)
	}
	defer md.Disconnect(ctx)

	if err := md.Ping(ctx); err != nil {
		log.Fatal(err)
	}

}
