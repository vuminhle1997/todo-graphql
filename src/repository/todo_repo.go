package repository

import (
	"context"
	"fmt"
	"time"
	"todo-graphql/src/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	DATABASE        = "todo-graphql"
	TODO_COLLECTION = "todos"
)

type TodoRepository interface {
	Save(todo models.Todo) *models.Todo
	GetById(id string) *models.Todo
	GetToDos() []*models.Todo
	UpdateById(id string) *models.Todo
	DeleteById(id string) *models.Todo
}

type TodoInput struct {
	Title string
	Description string
	Done bool
}

func (db *Database) Save(todo models.Todo) *models.Todo {
	collection := db.Client.Database(DATABASE).Collection(TODO_COLLECTION)
	res, err := collection.InsertOne(context.TODO(), todo)

	if err != nil {
		fmt.Println("Could not insert document into Collection")
	}
	fmt.Println("Stored into DB ")

	filter := bson.M{
		"_id": res.InsertedID,
	}

	var doc *models.Todo
	er := collection.FindOne(context.TODO(), filter).Decode(&doc)

	if er != nil {
		fmt.Println("404 not found")
	}

	return doc
}

func (db *Database) GetById(id string) *models.Todo {
	collection := db.Client.Database(DATABASE).Collection(TODO_COLLECTION)

	var todo *models.Todo
	_id, _err := primitive.ObjectIDFromHex(id)
	if _err != nil {
		fmt.Println("not valid ObjectID", _err)
		return nil
	}

	filter := bson.M{
		"_id": _id,
	}

	err := collection.FindOne(context.TODO(), filter).Decode(&todo)
	if err != nil {
		fmt.Println("500 Decode Error", err)
		return nil
	}

	return todo
}

func (db *Database) GetToDos() []*models.Todo {
	collection := db.Client.Database(DATABASE).Collection(TODO_COLLECTION)

	cursor, err := collection.Find(context.TODO(), bson.D{})

	if err != nil {
		fmt.Println("500")
		return nil
	}
	defer cursor.Close(context.TODO())
	var todos []*models.Todo

	for cursor.Next(context.TODO()) {
		var todo *models.Todo

		err := cursor.Decode(&todo)
		if err != nil {
			fmt.Println("error")
			return nil
		}
		todos = append(todos, todo)
	}

	return todos
}


// TODO: Fix this
func (db *Database) UpdateById(id string, todoInput TodoInput) *models.Todo {
	collection := db.Client.Database(DATABASE).Collection(TODO_COLLECTION)

	_id, _err := primitive.ObjectIDFromHex(id)
	if _err != nil {
		fmt.Println("Not valid Object ID")
		return nil
	}
	filter := bson.M{
		"_id": _id,
	}
	update := bson.M{
		"title": todoInput.Title,
		"description": todoInput.Description,
		"done": todoInput.Done,
		"updatedAt": time.Now(),
	}
	fmt.Println(filter, update)
	var res *models.Todo
	err := collection.FindOneAndUpdate(context.TODO(), filter, update).Decode(&res)
	if err != nil {
		fmt.Println(500, "Error occured")
		return nil
	}
	return res
}