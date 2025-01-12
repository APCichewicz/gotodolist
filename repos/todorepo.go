package repos

import (
	"context"

	"github.com/apcichewicz/gotodolist/models"
	"github.com/uptrace/bun"
)

type TodoRepository struct {
	db *bun.DB
}

func newTodoRepo(db *bun.DB) TodoRepository {
	return TodoRepository{db: db}
}

func (tr TodoRepository) GetAllTodos() ([]models.ToDo, error) {
	ctx := context.Background()
	var todos []models.ToDo
	err := tr.db.NewSelect().Model(&todos).Scan(ctx)
	return todos, err
}

func (tr TodoRepository) GetTodoByID(id int) (models.ToDo, error) {
	ctx := context.Background()
	var todo models.ToDo
	err := tr.db.NewSelect().Model(&todo).Where("id = ?", id).Scan(ctx)
	return todo, err
}

func (tr TodoRepository) CreateTodo(todo models.ToDo) (models.ToDo, error) {
	ctx := context.Background()
	_, err := tr.db.NewInsert().Model(&todo).Exec(ctx)
	return todo, err
}

func (tr TodoRepository) UpdateTodo(todo models.ToDo) (models.ToDo, error) {
	ctx := context.Background()
	_, err := tr.db.NewUpdate().Model(&todo).WherePK().Exec(ctx)
	return todo, err
}

func (tr TodoRepository) DeleteTodo(id int) error {
	ctx := context.Background()
	_, err := tr.db.NewDelete().Model(&models.ToDo{}).Where("id = ?", id).Exec(ctx)
	return err
}
