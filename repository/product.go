package repository

import (
	"github.com/api/service/dbutils/schema"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductQueries struct {
	*sqlx.DB
}

func (q *ProductQueries) GetProductById(pid uuid.UUID) (schema.Product, error) {
	product := schema.Product{}
	// Define query string.
	query := `SELECT * FROM products WHERE pid=?`
	// Send query to database.
	err := q.Get(&product, query, pid)
	if err != nil {
		// Return empty object and error.
		return product, err
	}
	return product, nil
}
func (q *ProductQueries) GetReviewsByPid(pid uuid.UUID) ([]schema.Review, error) {
	reviews := []schema.Review{}
	query := `SELECT * FROM reviews WHERE pid=?`
	err := q.Select(&reviews, query, pid)
	if err != nil {
		return reviews, err
	}
	return reviews, nil
}

func (q *ProductQueries) GetProducts() ([]schema.Product, error) {
	// Define products variable.
	products := []schema.Product{}
	// Define query string.
	query := `SELECT * FROM products`
	// Send query to database.
	err := q.Select(&products, query)
	if err != nil {
		// Return empty object and error.
		return products, err
	}
	return products, nil
}

func (q *ProductQueries) DeleteProduct(pid uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM users WHERE pid=?`
	// Send query to database.
	_, err := q.Exec(query, pid)
	if err != nil {
		// Return only error.
		return err
	}

	// This query returns nothing.
	return nil
}

func (q *ProductQueries) CreateProduct(product *schema.Product) error {
	// Define query string.
	query := `INSERT INTO products (pid,name,description,category,price,stockQuantity,image,details,productDiscount,auditInfo)
	 VALUES (:pid,:name,:description,:category,:price,:stockQuantity,:image,:details,:productDiscount,CONVERT(:auditInfo using utf8mb4))`
	// Send query to database.
	_, err := q.NamedExec(query, product)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}

func (q *ProductQueries) CreateProductReview(review *schema.Review) error {
	// Define query string.
	query := `INSERT INTO reviews (rid,pid,createdAt,rating,reviewerName,comment,email)
	 VALUES (:rid,:pid,:createdAt,:rating,:reviewerName,:comment,:email)`
	// Send query to database.
	_, err := q.NamedExec(query, review)
	if err != nil {
		// Return only error.
		return err
	}
	// This query returns nothing.
	return nil
}
