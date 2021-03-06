package routes

import (
	handler_admin "ppob/admin/handler"
	"ppob/app/middlewares"
	"ppob/helper/valid"
	handler_products "ppob/products/handler"
	handler_transaction "ppob/transaction/handler"
	handler_users "ppob/users/handler"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type ControllerList struct {
	JWTMiddleware      middleware.JWTConfig
	UserHandler        handler_users.UsersHandler
	ProductsHandler    handler_products.ProductsHandler
	TransactionHandler handler_transaction.TransactionHandler
	AdminHandler       handler_admin.AdminHandler
}

// const server = "https://36e2-2001-448a-1102-1a0f-350a-677f-f95c-668a.ap.ngrok.io/"

func (cl *ControllerList) RouteRegister(e *echo.Echo) {

	// log
	middlewares.LogMiddleware(e)

	// access public
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders:     []string{"*"},
		AllowCredentials: true,
		MaxAge:           2592000,
	}))
	e.POST("/login", cl.UserHandler.Authorization)
	e.POST("/register", cl.UserHandler.Register)
	e.GET("/admin/users", cl.UserHandler.GetUsers)
	e.GET("/admin/users/:phone", cl.UserHandler.GetUserForAdmin)
	// validasi
	e.POST("/validation", cl.UserHandler.VerifUser)
	// access public product
	e.GET("/products", cl.ProductsHandler.GetAllProduct)
	e.GET("/products/category/:category_id", cl.ProductsHandler.GetProductByCategory)
	e.GET("/products/:id", cl.ProductsHandler.GetProduct)
	e.GET("/detail/:product_slug", cl.ProductsHandler.GetDetailsProduct)
	e.GET("/category", cl.ProductsHandler.GetCategories)
	e.GET("/category/:id", cl.ProductsHandler.GetCategoryByID)
	e.POST("/transaction/callback_invoice", cl.TransactionHandler.Callback_Invoice)

	// access customer
	authUser := e.Group("users")
	authUser.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("customer", cl.UserHandler))
	// make pin
	authUser.POST("/pin", cl.UserHandler.MakePin)
	authUser.GET("/session", cl.UserHandler.GetUserSession)
	authUser.POST("/profile", cl.UserHandler.UpdateProfile)
	// transaction
	authUser.POST("/checkout/:detail_slug", cl.TransactionHandler.Checkout)
	authUser.GET("/transaction/:id", cl.TransactionHandler.GetTransactionByPaymentSuccess)
	authUser.GET("/transaction/success/:payment_id", cl.TransactionHandler.GetTransactionByPaymentSuccess)
	authUser.GET("/history", cl.TransactionHandler.GetHistoryTransaction)
	authUser.GET("/favorite", cl.TransactionHandler.FavoriteUser)

	// manage product endpoint (admin)
	authAdmin := e.Group("admin")
	authAdmin.Use(middleware.JWTWithConfig(cl.JWTMiddleware), valid.RoleValidation("admin", cl.UserHandler))
	authAdmin.Use(middleware.CORS())

	authAdmin.POST("/products/:category_id", cl.ProductsHandler.InsertProduct)
	authAdmin.PUT("/products/edit/:id", cl.ProductsHandler.EditProduct)
	authAdmin.DELETE("/products/delete/:id", cl.ProductsHandler.DestroyProduct)
	// manage detail product (admin)

	authAdmin.POST("/detail/:product_slug", cl.ProductsHandler.InsertDetail)
	authAdmin.PUT("/detail/edit/:getID", cl.ProductsHandler.EditDetail)
	authAdmin.DELETE("/detail/delete/:getID", cl.ProductsHandler.DestroyDetail)
	// manage category (admin)

	authAdmin.POST("/category", cl.ProductsHandler.InsertCategory)
	authAdmin.PUT("/category/edit/:id", cl.ProductsHandler.EditCategory)
	authAdmin.DELETE("/category/delete/:id", cl.ProductsHandler.DestroyCategory)

	// transaction
	authAdmin.GET("/transaction", cl.AdminHandler.GetAllTransaction)
	authAdmin.GET("/countAllItems", cl.AdminHandler.CountAllItems)
	authAdmin.GET("/detail_transaction/:payment_id", cl.AdminHandler.DetailTransaction)
}
