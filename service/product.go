package service

import (
	"fmt"
	"strings"

	"github.com/api/repository/request"
	"github.com/api/repository/response"
	"github.com/api/service/dbutils/schema"
	"github.com/api/util"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

// CreateProduct func creates a new product.
// @Description Create a new product.
// @Summary Create a new product
// @Tags Product
// @Accept json
// @Produce json
// @Param create_product body request.CreateProductRequest true "Create product"
// @Success 200 {object} response.RequestResponse
// @Security BearerAuth
// @Router /api/products/create [post]
func (cfg *AppConfig) CreateProduct(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	// Get claims from JWT.
	data := cfg.JwtCredentials(c)
	isAdmin := data["isAdmin"].(bool)
	if !isAdmin {
		newErr.Message = "Admin rights needed"
		newErr.Code = fiber.StatusForbidden
		cfg.Logger.Warn(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}

	createProduct := new(request.CreateProductRequest)
	// Check, if received JSON data is valid.
	if err := c.BodyParser(createProduct); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate createProduct fields.
	if err := cfg.Validate.Struct(createProduct); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}

	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}

	loggedInUserid := data["userid"].(string)
	audit := &schema.AuditEntity{CreatedAt: util.TimeNow(), LastModifiedAt: util.NullTime(), LastModifiedBy: "none", CreatedBy: loggedInUserid}
	newProduct := &schema.Product{Pid: util.Ident(), Name: createProduct.Name,
		Description: createProduct.Description, Category: createProduct.Category,
		Price: createProduct.Price, StockQuantity: createProduct.StockQuantity,
		Image: createProduct.Image, Details: createProduct.Details,
		ProductDiscount: 0.0, AuditInfo: *audit,
	}
	// Validate product fields.
	if err := cfg.Validate.Struct(newProduct); err != nil {
		// Return, if some fields are not valid.
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	fmt.Printf("%#v\n", newProduct)
	if err := db.CreateProduct(newProduct); err != nil {
		newErr.Message = err.Error()
		if strings.Contains(err.Error(), "Error 1062") {
			newErr.Message = "Duplicate key: product already exist"
		}
		cfg.Logger.Error(newErr.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = newProduct.Pid
	return c.Status(fiber.StatusOK).JSON(response)
}

// CreateReview func creates a new product review.
// @Description Create a new product review.
// @Summary Create a new product review
// @Tags Product
// @Accept json
// @Produce json
// @Param product_review body request.CreateProductReviewRequest true "Create a product review"
// @Success 200 {object} response.RequestResponse
// @Router /api/products/create/review [post]
func (cfg *AppConfig) CreateReview(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	createReview := new(request.CreateProductReviewRequest)
	// Check, if received JSON data is valid.
	if err := c.BodyParser(createReview); err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid JSON body"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)

	}
	// Validate createReview fields.
	if err := cfg.Validate.Struct(createReview); err != nil {
		// Return, if some fields are not valid.
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = fmt.Sprintf("Invalid Field(s) :%v", util.ValidatorErrors(err))
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	foundProduct, err := db.GetProductById(createReview.Pid)
	if err != nil {
		newErr.Message = "Product with the given ID is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	newReview := schema.Review{
		Rid:          util.Ident(),
		Pid:          foundProduct.Pid,
		CreatedAt:    util.TimeNow(),
		Rating:       createReview.Rating,
		ReviewerName: createReview.ReviewerName,
		Comment:      createReview.Comment,
		Email:        createReview.Email,
	}
	// Validate product review fields.
	if err := cfg.Validate.Struct(newReview); err != nil {
		// Return, if some fields are not valid.
		cfg.Logger.Error(util.ValidatorErrors(err))
		return c.Status(newErr.Code).JSON(newErr)
	}
	if err := db.CreateProductReview(&newReview); err != nil {
		newErr.Message = err.Error()
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	response := response.NewResponse(c)
	response.Result = newReview.Rid
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetProducts method for getting all existing products.
// @Description Get all existing products.
// @Summary Get all existing products
// @Tags Product
// @Accept json
// @Produce json
// @Failure 500 {object} response.ErrorResponse
// @Success 200 {array} response.RequestResponse
// @Router /api/products [get]
func (cfg *AppConfig) GetAllProducts(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	// Create database connection.
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	products, err := db.GetProducts()
	if err != nil {
		newErr.Message = "Products not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	result := response.GetProductsResponse{}
	for _, product := range products {
		result.Products = append(result.Products, response.GetProductResponse{
			Pid:             product.Pid,
			Name:            product.Name,
			Description:     product.Description,
			Category:        product.Category,
			Price:           product.Price,
			StockQuantity:   product.StockQuantity,
			Image:           product.Image,
			Details:         product.Details,
			ProductDiscount: product.ProductDiscount,
		})

	}
	response := response.NewResponse(c)
	response.Result = result
	return c.Status(fiber.StatusOK).JSON(response)
}

// GetSingleProduct from db
// @Description Get Product by given ID.
// @Summary Get product by given pid
// @Tags Product
// @Accept json
// @Produce json
// @Param pid path string true "pid"
// @Success 200 {object} response.RequestResponse
// @Failure 400 {object} response.ErrorResponse
// @Failure 500 {object} response.ErrorResponse
// @Router /api/products/{pid} [get]
func (cfg *AppConfig) GetSingleProduct(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	pid, err := uuid.Parse(c.Params("pid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid pid"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	product, err := db.GetProductById(pid)
	if err != nil {
		newErr.Message = "Product with pid is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	result := &response.GetProductResponse{
		Pid:             product.Pid,
		Name:            product.Name,
		Description:     product.Description,
		Category:        product.Category,
		Price:           product.Price,
		StockQuantity:   product.StockQuantity,
		Image:           product.Image,
		Details:         product.Details,
		ProductDiscount: product.ProductDiscount,
	}
	response := response.NewResponse(c)
	response.Result = result
	return c.Status(fiber.StatusOK).JSON(response)
}

// DeleteProduct from db
// @Description Delete product by a given pid.
// @Summary Delete Product by given pid
// @Tags Product
// @Accept json
// @Produce json
// @Param pid path string true "pid"
// @Success 200 {object} response.RequestResponse
// @Security BearerAuth
// @Router /api/products/{pid} [delete]
func (cfg *AppConfig) DeleteProduct(c *fiber.Ctx) error {
	newErr := response.NewErrorResponse()
	//Get claims from JWT.
	data := cfg.JwtCredentials(c)
	isAdmin := data["isAdmin"].(bool)
	if !isAdmin {
		newErr.Message = "Admin rights needed"
		newErr.Code = fiber.StatusForbidden
		cfg.Logger.Warn(newErr.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	pid, err := uuid.Parse(c.Params("pid"))
	if err != nil {
		newErr.Code = fiber.ErrBadRequest.Code
		newErr.Message = "Invalid pid"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	db, err := DbWithQueries(cfg)
	if err != nil {
		newErr.Message = "Couldnt connect to DB!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusInternalServerError).JSON(newErr)
	}
	foundProduct, err := db.GetProductById(pid)
	if err != nil {
		newErr.Message = "Product with the given ID is not found!"
		cfg.Logger.Error(err.Error())
		return c.Status(fiber.StatusNotFound).JSON(newErr)
	}
	// Delete product by given pid.
	if err := db.DeleteProduct(foundProduct.Pid); err != nil {
		newErr.Message = "Couldnt delete the product!"
		cfg.Logger.Error(err.Error())
		return c.Status(newErr.Code).JSON(newErr)
	}
	response := response.NewResponse(c)
	return c.Status(response.Code).JSON(response)

}
