package main

import (
	database "backend/model/generated_model"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

func run() error {
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, "postgres://postgres:postgres@127.0.0.1:5432/postgres?sslmode=disable")
	if err != nil {
		return err
	}
	defer conn.Close(ctx)

	queries := database.New(conn)

	// list all users
	users, err := queries.ListUsers(ctx)
	if err != nil {
		return err
	}
	log.Println(users)

	// create a user
	insertedUser, err := queries.CreateUser(ctx, database.CreateUserParams{
		Email:    "example@email.com",
		FullName: "John Doe",
	})
	if err != nil {
		return err
	}
	log.Println(insertedUser)

	// create a not
	insertedNote, err := queries.CreateNote(ctx, database.CreateNoteParams{
		UserID:  insertedUser.ID,
		Title:   "Hello, world!",
		Content: "This is a test note.",
	})
	if err != nil {
		return err
	}
	log.Println(insertedNote)

	// delete a note
	err = queries.DeleteNote(ctx, insertedNote.ID)
	if err != nil {
		return err
	}

	// delete a user
	err = queries.DeleteUser(ctx, insertedUser.ID)
	if err != nil {
		return err
	}

	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
