package repos

import (
	"context"

	"github.com/apcichewicz/gotodolist/models"
	"github.com/uptrace/bun"
)

type ToDoRepository struct {
	db *bun.DB
}

func newToDoRepo(db *bun.DB) ToDoRepository {
	return ToDoRepository{db: db}
}

func (tr ToDoRepository) GetAllToDos() ([]models.ToDo, error) {
	ctx := context.Background()
	var todos []models.ToDo
	err := tr.db.NewSelect().Model(&todos).Scan(ctx)
	return todos, err
}

func (tr ToDoRepository) GetToDoByID(id int) (models.ToDo, error) {
	ctx := context.Background()
	var todo models.ToDo
	err := tr.db.NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx)
	return todo, err
}

func (tr ToDoRepository) CreateToDo(todo models.ToDo) (models.ToDo, error) {
	ctx := context.Background()
	_, err := tr.db.NewInsert().Model(&todo).Exec(ctx)
	return todo, err
}

func (tr ToDoRepository) UpdateToDo(todo models.ToDo) (models.ToDo, error) {
	ctx := context.Background()
	_, err := tr.db.NewUpdate().Model(&todo).WherePK().Exec(ctx)
	return todo, err
}

func (tr ToDoRepository) DeleteToDo(id int) error {
	ctx := context.Background()
	_, err := tr.db.NewDelete().Model(&models.ToDo{}).Where("id = ?", id).Exec(ctx)
	return err
}
