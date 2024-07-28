package repository

import "github.com/jmoiron/sqlx"

type PaginationQueries struct {
	*sqlx.DB
}

type Pagination struct {
	Next          int
	Previous      int
	RecordPerPage int
	CurrentPage   int
	TotalPage     int
}

func (p *Pagination) Pageable(table string, page int) {}
