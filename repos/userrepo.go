package repos

import (
	"context"
	"fmt"

	"github.com/apcichewicz/gotodolist/models"
	"github.com/uptrace/bun"
)

type UserRepository struct {
	db *bun.DB
}

func newUserRepo(db *bun.DB) UserRepository {
	return UserRepository{db: db}
}

func (ur UserRepository) GetUserByID(id int) (models.User, error) {
	ctx := context.Background()
	var usr models.User
	query := fmt.Sprintf("id = %d", id)
	err := ur.db.NewSelect().Model(&usr).Where(query).Scan(ctx, &usr)
	return usr, err
}

func (ur UserRepository) GetUserByEmail(email string) (models.User, error) {
	ctx := context.Background()
	var usr models.User
	query := fmt.Sprintf("email = %s", email)
	err := ur.db.NewSelect().Model(&usr).Where(query).Scan(ctx, &usr)
	return usr, err
}

func (ur UserRepository) CreateUser(user models.User) (models.User, error) {
	ctx := context.Background()
	_, err := ur.db.NewInsert().Model(&user).Exec(ctx)
	return user, err
}

func (ur UserRepository) UpdateUser(user models.User) (models.User, error) {
	ctx := context.Background()
	_, err := ur.db.NewUpdate().Model(&user).WherePK().Exec(ctx)
	return user, err
}

func (ur UserRepository) DeleteUser(id int) error {
	ctx := context.Background()
	_, err := ur.db.NewDelete().Model(&models.User{}).Where("id = ?", id).Exec(ctx)
	return err
}
