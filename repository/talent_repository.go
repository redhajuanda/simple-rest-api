package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Masterminds/squirrel"
)

// TalentRepository that related with talent repository
type TalentRepository interface {
	Create(ctx context.Context, talent entity.Talent) (int64, error)
}

// TalentRepositoryImpl implementation interface
type TalentRepositoryImpl struct {
	db *sql.Tx
}

// NewTalentRepositoryImpl constructor
func NewTalentRepositoryImpl(db *sql.Tx) TalentRepository {
	return TalentRepositoryImpl{
		db: db,
	}
}

// Create function for insert to talent table
func (t TalentRepositoryImpl) Create(ctx context.Context, talent entity.Talent) (int64, error) {
	query, args, err := squirrel.Insert("talent").Columns(
		"name",
	).Values(talent.Name).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := t.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	log.Printf("success insert talent : %s, with id %d", talent.Name, id)

	return id, nil
}

// FindTalent function to find talent by talent id
func (t TalentRepositoryImpl) FindTalent(ctx context.Context, ID int64) error {
	return nil
}