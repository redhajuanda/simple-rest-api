package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/Alfabetss/simple-rest-api/entity"
	"github.com/Masterminds/squirrel"
)

// ExperienceRepository interface that related with experience repository
type ExperienceRepository interface {
	Create(ctx context.Context, exp entity.Experience) (int64, error)
	FindTalentExperiences(ctx context.Context, talentID int64) ([]entity.Experience, error)
	Delete(ctx context.Context, talentID int64) error
	Update(ctx context.Context, exp entity.Experience) error
	FindExperience(ctx context.Context, ID int64) (*entity.Experience, error)
}

// ExperienceRepositoryImpl implementation interface
type ExperienceRepositoryImpl struct {
	db *sql.Tx
}

// NewExperienceRepositoryImpl constructor
func NewExperienceRepositoryImpl(db *sql.Tx) ExperienceRepository {
	return ExperienceRepositoryImpl{
		db: db,
	}
}

// FindTalentExperiences function to find experience by talent id
func (e ExperienceRepositoryImpl) FindTalentExperiences(ctx context.Context, talentID int64) ([]entity.Experience, error) {
	var experiences []entity.Experience
	query, args, err := squirrel.Select("id", "company", "talent_id").
		From("experience").
		Where(squirrel.Eq{"talent_id": talentID}).ToSql()

	if err != nil {
		return experiences, err
	}

	rows, err := e.db.Query(query, args...)
	if err != nil {
		return experiences, err
	}

	for rows.Next() {
		var exp entity.Experience
		err := rows.Scan(
			&exp.ID,
			&exp.Company,
			&exp.TalentID,
		)
		if err != nil {
			return experiences, err
		}

		experiences = append(experiences, exp)
	}

	return experiences, nil
}

// Create function to insert to experience table
func (e ExperienceRepositoryImpl) Create(ctx context.Context, exp entity.Experience) (int64, error) {
	query, args, err := squirrel.Insert("experience").Columns("company", "talent_id").
		Values(exp.Company, exp.TalentID).ToSql()
	if err != nil {
		return 0, err
	}

	res, err := e.db.ExecContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}

	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	log.Printf("success insert experience for talent id : %d, company name : %s", id, exp.Company)
	return id, nil
}

// Delete function to delete experience row by talent id
func (e ExperienceRepositoryImpl) Delete(ctx context.Context, talentID int64) error {
	query, args, err := squirrel.Delete("experience").
		Where(squirrel.Eq{"talent_id": talentID}).ToSql()
	if err != nil {
		return err
	}

	_, err = e.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// Update function to update experience by id
func (e ExperienceRepositoryImpl) Update(ctx context.Context, exp entity.Experience) error {
	query, args, err := squirrel.Update("experience").
		Set("company", exp.Company).
		Where(squirrel.Eq{"id": exp.ID}).ToSql()
	if err != nil {
		return err
	}

	_, err = e.db.Exec(query, args...)
	if err != nil {
		return err
	}

	return nil
}

// FindExperience function to find experience by id
func (e ExperienceRepositoryImpl) FindExperience(ctx context.Context, ID int64) (*entity.Experience, error) {
	var experience entity.Experience
	query, args, err := squirrel.Select("id", "company", "talent_id").
		From("experience").
		Where(squirrel.Eq{"id": ID}).ToSql()
	if err != nil {
		return &experience, err
	}

	row := e.db.QueryRow(query, args...)
	err = row.Scan(
		&experience.ID,
		&experience.Company,
		&experience.TalentID,
	)
	if err != nil {
		return &experience, err
	}

	return &experience, nil
}
