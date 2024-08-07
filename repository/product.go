package repository

import (
	"fmt"

	"github.com/api/repository/response"
	"github.com/api/repository/schema"
	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
)

type ProductQueries struct {
	*sqlx.DB
}

func (q *ProductQueries) CreateProduct(product *schema.Product) error {
	// Define query string.
	query := `INSERT INTO products (pid,name,description,category,price,stockQuantity,image,details,productDiscount,auditInfo)
	 VALUES (:pid,:name,:description,:category,:price,:stockQuantity,:image,:details,:productDiscount,CONVERT(:auditInfo using utf8mb4))`
	// Send query to database.
	_, err := q.NamedExec(query, product)
	if err != nil {
		// Return only error.
		return fmt.Errorf("error, Creating a product:, %w", err)
	}
	// This query returns nothing.
	return nil
}
func (q *ProductQueries) CreateProducts(products []schema.Product) error {
	// Define query string.
	query := `INSERT INTO products (pid,name,description,category,price,stockQuantity,image,details,productDiscount,auditInfo)
	 VALUES (:pid,:name,:description,:category,:price,:stockQuantity,:image,:details,:productDiscount,CONVERT(:auditInfo using utf8mb4))`
	// Send query to database.
	_, err := q.NamedExec(query, products)
	if err != nil {
		// Return only error.
		return fmt.Errorf("error, Creating a products:, %w", err)
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
		return fmt.Errorf("error, Creating a product review:, %w", err)
	}
	// This query returns nothing.
	return nil
}

func (q *ProductQueries) GetProductById(pid uuid.UUID) (schema.Product, []response.GetProductReviewResponse, error) {
	query := `SELECT p.*,r.rid, r.createdAt,r.rating,r.reviewerName,r.comment,r.email
    FROM products AS p LEFT JOIN reviews AS r ON p.pid = r.pid WHERE p.pid = ?`
	product := schema.Product{}
	reviews := []response.GetProductReviewResponse{}
	rows, err := q.Queryx(query, pid)
	if err != nil {
		//Return empty object and error.
		return product, reviews, err
	}
out:
	for rows.Next() {
		review := response.GetProductReviewResponse{}
		if err := rows.Scan(&product.Pid, &product.Name, &product.Description, &product.Category,
			&product.Price,
			&product.StockQuantity, &product.Image, &product.Details, &product.ProductDiscount, &product.AuditInfo,
			&review.Rid, &review.CreatedAt, &review.Rating,
			&review.ReviewerName, &review.Comment, &review.Email); err != nil {
			return product, reviews, err
		}
		if review.Rid == nil {
			break out
		}
		reviews = append(reviews, review)
	}
	return product, reviews, nil
}

func (q *ProductQueries) GetProducts() ([]schema.Product, error) {
	// Define products variable.
	products := []schema.Product{}
	// // Define query string.
	query := `SELECT * FROM products LIMIT 12`
	// Send query to database.
	err := q.Select(&products, query)
	if err != nil {
		// Return empty object and error.
		return products, err
	}
	return products, nil
}

func (q *ProductQueries) GetReviewsByPid(pid uuid.UUID) ([]schema.Review, error) {
	reviews := []schema.Review{}
	query := `SELECT * FROM reviews WHERE pid=?`
	err := q.Select(&reviews, query, pid)
	if err != nil {
		return reviews, fmt.Errorf("error, Getting product reviews:, %w", err)
	}
	return reviews, nil
}

func (q *ProductQueries) DeleteProduct(pid uuid.UUID) error {
	// Define query string.
	query := `DELETE FROM users WHERE pid=?`
	// Send query to database.
	_, err := q.Exec(query, pid)
	if err != nil {
		// Return only error.
		return fmt.Errorf("error, Deleting the product:, %w", err)
	}

	// This query returns nothing.
	return nil
}
