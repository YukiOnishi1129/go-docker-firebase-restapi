package repository

import (
	"cloud.google.com/go/firestore"
	"context"
	"github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/todo"
	"github.com/harakeishi/gats"
)

type todoRepository struct {
	client *firestore.Client
}

func NewTodoRepository(client *firestore.Client) todo.TodoRepository {
	return &todoRepository{client}
}

func (r *todoRepository) Save(ctx context.Context, todo *todo.Todo) error {
	_, err := r.client.Collection("todos").Doc(todo.ID()).Set(ctx, map[string]interface{}{
		"title":       todo.Title(),
		"description": todo.Description(),
	})
	if err != nil {
		return err
	}
	return nil
}

func (r *todoRepository) FindByID(ctx context.Context, id string) (*todo.Todo, error) {
	doc, err := r.client.Collection("todos").Doc(id).Get(ctx)
	if err != nil {
		return nil, err
	}

	docTitle, err := doc.DataAt("title")
	if err != nil {
		return nil, err
	}
	title, _ := gats.ToString(docTitle)

	docDescription, err := doc.DataAt("description")
	if err != nil {
		return nil, err
	}
	description, _ := gats.ToString(docDescription)

	td, err := todo.Reconstruct(
		doc.Ref.ID,
		title,
		description,
	)
	if err != nil {
		return nil, err
	}
	return td, nil
}

func (r *todoRepository) FindAll(ctx context.Context) ([]*todo.Todo, error) {
	doc, err := r.client.Collection("todos").Documents(ctx).GetAll()
	if err != nil {
		return nil, err
	}

	var todos []*todo.Todo
	for _, d := range doc {

		docTitle, err := d.DataAt("title")
		if err != nil {
			return nil, err
		}
		title, _ := gats.ToString(docTitle)

		docDescription, err := d.DataAt("description")
		if err != nil {
			return nil, err
		}
		description, _ := gats.ToString(docDescription)

		td, err := todo.Reconstruct(
			d.Ref.ID,
			title,
			description,
		)
		if err != nil {
			return nil, err
		}

		todos = append(todos, td)
	}
	return todos, nil
}
