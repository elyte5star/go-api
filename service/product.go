package service

// import (
// 	"database/sql"
// 	"github.com/api/util"
// 	"github.com/gofiber/fiber/v2"
// )

// func GetAllProducts(c *fiber.Ctx) error{
// 	// query product table in the database
// 	rows, err := database.DB.Query("SELECT name, description, category, amount FROM products order by name")
// 	if err != nil {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"error":   err,
// 		})
		
// 	}
// 	defer rows.Close()
// 	result := model.Products{}
// 	for rows.Next() {
// 		product := model.Product{}
// 		err := rows.Scan(&product.Name, &product.Description, &product.Category, &product.Amount)
// 		// Exit if we get an error
// 		if err != nil {
// 			return c.Status(500).JSON(&fiber.Map{
// 				"success": false,
// 				"error":   err,
// 			})
			
// 		}
// 		// Append Product to Products
// 		result.Products = append(result.Products, product)
// 	}
// 	// Return Products in JSON format
// 	if err := c.JSON(&fiber.Map{
// 		"success": true,
// 		"product": result,
// 		"message": "All product returned successfully",
// 	}); err != nil {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"message": err,
// 		})
		
// 	}
// }

// // GetSingleProduct from db
// func GetSingleProduct(c *fiber.Ctx) error{
// 	id := c.Params("pid")
// 	product := model.Product{}
// 	// query product database
// 	row, err := database.DB.Query("SELECT * FROM products WHERE id = $1", id)
// 	if err != nil {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"message": err,
// 		})
		
// 	}
// 	defer row.Close()
// 	// iterate through the values of the row
// 	for row.Next() {
// 		switch err := row.Scan(&id, &product.Amount, &product.Name, &product.Description, &product.Category); err {
// 		case sql.ErrNoRows:
// 			util.Logger().Info("No rows were returned!")
// 			c.Status(500).JSON(&fiber.Map{
// 				"success": false,
// 				"message": err,
// 			})
// 		case nil:
// 			util.Logger().Info(product.Name, product.Description, product.Category, product.Amount)
// 		default:
// 			//   panic(err)
// 			c.Status(500).JSON(&fiber.Map{
// 				"success": false,
// 				"message": err,
// 			})
// 		}
// 	}
// }

// // DeleteProduct from db
// func DeleteProduct(c *fiber.Ctx) error{
// 	id := c.Params("pid")
// 	// query product table in database
// 	res, err := database.DB.Query("DELETE FROM products WHERE id = $1", id)
// 	if err != nil {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"error":   err,
// 		})
		
// 	}
// 	// Print result
// 	util.Logger().Info(res)
// 	// return product in JSON format
// 	if err := c.JSON(&fiber.Map{
// 		"success": true,
// 		"message": "product deleted successfully",
// 	}); err != nil {
// 		return c.Status(500).JSON(&fiber.Map{
// 			"success": false,
// 			"error":   err,
// 		})
	
// 	}
// }
