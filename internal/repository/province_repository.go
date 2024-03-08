package repository

import (
	"context"
	"database/sql"
	"errors"
	"github.com/afdikomayte/go-api-regions-of-indonesia/entity"
)

type ProvinceRepository interface {
	GetAll(ctx context.Context, tx *sql.Tx) []entity.Province
	GetById(ctx context.Context, id int) (entity.Province, error)
}

type ProvinceRepositoryImpl struct {
}

func (repository *ProvinceRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []entity.Province {

	SQL := "select id, name from provinces"
	rows, err := tx.QueryContext(ctx, SQL)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	var provinces []entity.Province
	for rows.Next() {
		province := entity.Province{}
		err := rows.Scan(&province.Id, &province.Name)
		if err != nil {
			panic(err)
		}

		provinces = append(provinces, province)

	}

	return provinces

}

func (repository *ProvinceRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, provinceId int) (entity.Province, error) {
	SQL := "select id, name from province where id=?"
	rows, err := tx.QueryContext(ctx, SQL, provinceId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	province := entity.Province{}
	if rows.Next() {
		err := rows.Scan(&province.Id, &province.Name)
		if err != nil {
			panic(err)
		}
		return province, nil
	} else {
		return province, errors.New("Province is not found")
	}

}
