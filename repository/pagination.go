package repository

import (
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PaginationQueries struct {
	*sqlx.DB
}

type Pagination struct {
	Next             int  `json:"next"`
	Previous         int  `json:"previous"`
	RecordPerPage    int  `json:"recordPerPage"`
	CurrentPage      int  `json:"currentPage"`
	TotalPages       int  `json:"totalPages"`
	NumberOfElements int  `json:"numberOfElements"`
	Number           int  `json:"number"`
	TotalElements    int  `json:"totalElements"`
	Empty            bool `json:"empty"`
	Last             bool `json:"last"`
	First            bool `json:"first"`
}

func (q *PaginationQueries) Pageable(table string, page int) *Pagination {
	const LIMIT = 12
	var rowCount int
	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	_ = q.QueryRow(query).Scan(&rowCount)
	total := (rowCount / LIMIT)
	return &Pagination{}
}
