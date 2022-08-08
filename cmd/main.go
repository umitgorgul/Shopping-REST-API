package main

import (
	"log"
	"net/http"

	"github/umitgorgul/Shopping-Cart-REST-API/internal/config"
	cart_controller "github/umitgorgul/Shopping-Cart-REST-API/internal/controllers/cart"
	category_controller "github/umitgorgul/Shopping-Cart-REST-API/internal/controllers/category"
	order_controller "github/umitgorgul/Shopping-Cart-REST-API/internal/controllers/order"
	product_controller "github/umitgorgul/Shopping-Cart-REST-API/internal/controllers/product"
	cart_repository "github/umitgorgul/Shopping-Cart-REST-API/internal/repositories/cart"
	category_repository "github/umitgorgul/Shopping-Cart-REST-API/internal/repositories/category"
	order_repository "github/umitgorgul/Shopping-Cart-REST-API/internal/repositories/order"
	product_repository "github/umitgorgul/Shopping-Cart-REST-API/internal/repositories/product"
	user_repository "github/umitgorgul/Shopping-Cart-REST-API/internal/repositories/user"
	database_handler "github/umitgorgul/Shopping-Cart-REST-API/pkg/database_handler"

	"github.com/gin-gonic/gin"
)

func main() {
	appConfig, err := config.GetAllConfigValues("C:/Users/PC/Desktop/Shopping-Cart-REST-API/config/location.qa.yaml")
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()
	// registerMiddlewares(r)

	db := database_handler.NewMySQLDB(appConfig.DatabaseURI)

	cartRepo := cart_repository.NewCartRepository(db)
	categoryRepo := category_repository.NewCategoryRepository(db)
	orderRepo := order_repository.NewOrderRepository(db)
	productRepo := product_repository.NewProductRepository(db)
	userRepo := user_repository.NewUserRepository(db)

	if cartRepo.Migration() != nil {
		log.Fatal("Cart Migration failed")
	}
	if categoryRepo.Migration() != nil {
		log.Fatal("Category Migration failed")
	}
	if orderRepo.Migration() != nil {
		log.Fatal("Order Migration failed")
	}
	if productRepo.Migration() != nil {
		log.Fatal("Product Migration failed")
	}
	if userRepo.Migration() != nil {
		log.Fatal("User Migration failed")
	}

	//userRepo.InsertSampleData()

	cartCtrl := cart_controller.NewCartController(cartRepo, productRepo)
	categoryCtrl := category_controller.NewCategoryController(categoryRepo)
	orderCtrl := order_controller.NewOrderController(orderRepo, productRepo)
	productCtrl := product_controller.NewProductController(productRepo, categoryRepo)
	//userCtrl := user_controller.NewUserController(userRepo)

	// srv := &http.Server{
	// 	Addr:    ":8080",
	// 	Handler: r,
	// }

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	//cart routes
	r.GET("/cart/get/:id", cartCtrl.GetCart)
	r.POST("/cart/create", cartCtrl.CreateCart)
	r.PUT("/cart/update/:id", cartCtrl.UpdateCart)
	r.DELETE("/cart/delete/:id", cartCtrl.DeleteCart)

	//category routes
	r.GET("/category/get/:id", categoryCtrl.GetCategory)
	r.GET("/category/get", categoryCtrl.GetAllCategories)
	r.POST("/category/create", categoryCtrl.CreateCategory)
	r.POST("/category/create/bulk", categoryCtrl.CreateCategories)
	r.PUT("/category/update/:id", categoryCtrl.UpdateCategory)
	r.DELETE("/category/delete/:id", categoryCtrl.DeleteCategory)

	//product routes
	r.GET("/product/get/:id", productCtrl.GetProduct)
	r.GET("/product/get", productCtrl.GetProducts)
	r.POST("/product/create", productCtrl.CreateProduct)
	r.PUT("/product/update/:id", productCtrl.UpdateProduct)
	r.DELETE("/product/delete/:id", productCtrl.DeleteProduct)

	//order routes
	r.GET("/order/get/:id", orderCtrl.GetOrder)
	r.GET("/order/get/customer/:customerID", orderCtrl.GetOrdersByCustomerID)
	r.POST("/order/create", orderCtrl.CreateOrder)
	r.PUT("/order/update/:id", orderCtrl.UpdateOrder)
	r.DELETE("/order/delete/:id", orderCtrl.DeleteOrder)

	//user routes
	log.Fatal(http.ListenAndServe(":8080", r))

	// go func() {
	// 	if err := srv.ListenAndServe(); err != nil {
	// 		log.Printf("listen: %s \n", err)
	// 	}
	// }()
}

// func registerMiddlewares(r *gin.Engine) {
// 	r.Use(gin.LoggerWithFormatter(func(param gin.LogFormatterParams) string {
// 		// your custom format
// 		return fmt.Sprintf("%s - [%s] \"%s %s %s %d %s \"%s\" %s\"\n",
// 			param.ClientIP,
// 			param.TimeStamp.Format(time.RFC1123),
// 			param.Method,
// 			param.Path,
// 			param.Request.Proto,
// 			param.StatusCode,
// 			param.Latency,
// 			param.Request.UserAgent(),
// 			param.ErrorMessage,
// 		)
// 	}))
// 	r.Use(gin.Recovery())
// }
