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

	if err := md.AddMeigen(ctx, "author", "text"); err != nil {
		log.Fatal(err)
	}

	mfi, err := md.SearchMeigenFromID(ctx, 1)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mfi)

	mfa, err := md.SearchMeigenFromAuthor(ctx, "author")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mfa)

	mft, err := md.SearchMeigenFromText(ctx, "text")
	if err != nil {
		log.Fatal(err)
	}
	log.Println(mft)

	log.Println(md.CountMeigen(ctx))

}
