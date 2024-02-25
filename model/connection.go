package model

import (
	db_query "github.com/aadgraha/school_management/model/sqlc"
)

type Connect struct {
	Query *db_query.Queries
}
