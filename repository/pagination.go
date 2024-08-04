package repository

import (
	"database/sql"
	"fmt"

	"github.com/jmoiron/sqlx"
)

type PaginationQueries struct {
	*sqlx.DB
}

const limit = 12

type Page struct {
	Next             int  `json:"next,omitempty"`
	Previous         int  `json:"previous,omitempty"`
	CurrentPage      int  `json:"currentPage,omitempty"`
	TotalPages       int  `json:"totalPages,omitempty"`
	NumberOfElements int  `json:"numberOfElements,omitempty"`
	Number           int  `json:"number,omitempty"`
	TotalElements    int  `json:"totalElements"`
	Empty            bool `json:"empty"`
	Size             int  `json:"size"`
}

func (q *PaginationQueries) Pageable(table string, page int) (*Page, error) {
	var (
		pagination = Page{}
		rowCount   int
	)

	query := fmt.Sprintf("SELECT COUNT(*) FROM %s", table)
	err := q.QueryRow(query).Scan(&rowCount)
	if err == sql.ErrNoRows {
		pagination.Empty = true
		pagination.TotalElements = 0
		pagination.NumberOfElements = 0
		pagination.Size = limit
		return &pagination, nil
	} 
	pagination.Empty = false
	pagination.TotalElements = rowCount
	pagination.NumberOfElements = limit
	totalPages := (rowCount / limit)
	remainder := (rowCount % limit)
	if remainder == 0 {
		pagination.TotalPages = totalPages
	} else {
		pagination.TotalPages = totalPages + 1
	}
	pagination.CurrentPage = page
	pagination.Size = limit
	if page <= 0 {
		pagination.Next = page + 1
	} else if page < pagination.TotalPages {
		pagination.Previous = page - 1
		pagination.Next = page + 1
	} else if page == pagination.TotalPages {
		pagination.Previous = page - 1
		pagination.Next = 0

	}
	return &pagination, nil
}
