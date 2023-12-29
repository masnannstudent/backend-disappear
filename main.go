package main

import (
	"github.com/capstone-kelompok-7/backend-disappear/utils/database/seeder"
	"github.com/capstone-kelompok-7/backend-disappear/utils/email"
	"github.com/capstone-kelompok-7/backend-disappear/utils/sendnotif"
	"os"

	"github.com/capstone-kelompok-7/backend-disappear/config"
	hAddress "github.com/capstone-kelompok-7/backend-disappear/module/feature/address/handler"
	rAddress "github.com/capstone-kelompok-7/backend-disappear/module/feature/address/repository"
	sAddress "github.com/capstone-kelompok-7/backend-disappear/module/feature/address/service"
	hArticle "github.com/capstone-kelompok-7/backend-disappear/module/feature/article/handler"
	rArticle "github.com/capstone-kelompok-7/backend-disappear/module/feature/article/repository"
	sArticle "github.com/capstone-kelompok-7/backend-disappear/module/feature/article/service"
	hChatbot "github.com/capstone-kelompok-7/backend-disappear/module/feature/assistant/handler"
	rChatbot "github.com/capstone-kelompok-7/backend-disappear/module/feature/assistant/repository"
	sChatbot "github.com/capstone-kelompok-7/backend-disappear/module/feature/assistant/service"
	hAuth "github.com/capstone-kelompok-7/backend-disappear/module/feature/auth/handler"
	rAuth "github.com/capstone-kelompok-7/backend-disappear/module/feature/auth/repository"
	sAuth "github.com/capstone-kelompok-7/backend-disappear/module/feature/auth/service"
	hCarousel "github.com/capstone-kelompok-7/backend-disappear/module/feature/carousel/handler"
	rCarousel "github.com/capstone-kelompok-7/backend-disappear/module/feature/carousel/repository"
	sCarousel "github.com/capstone-kelompok-7/backend-disappear/module/feature/carousel/service"
	hCart "github.com/capstone-kelompok-7/backend-disappear/module/feature/cart/handler"
	rCart "github.com/capstone-kelompok-7/backend-disappear/module/feature/cart/repository"
	sCart "github.com/capstone-kelompok-7/backend-disappear/module/feature/cart/service"
	hCategory "github.com/capstone-kelompok-7/backend-disappear/module/feature/category/handler"
	rCategory "github.com/capstone-kelompok-7/backend-disappear/module/feature/category/repository"
	sCategory "github.com/capstone-kelompok-7/backend-disappear/module/feature/category/service"
	hChallenge "github.com/capstone-kelompok-7/backend-disappear/module/feature/challenge/handler"
	rChallenge "github.com/capstone-kelompok-7/backend-disappear/module/feature/challenge/repository"
	sChallenge "github.com/capstone-kelompok-7/backend-disappear/module/feature/challenge/service"
	hDashboard "github.com/capstone-kelompok-7/backend-disappear/module/feature/dashboard/handler"
	rDashboard "github.com/capstone-kelompok-7/backend-disappear/module/feature/dashboard/repository"
	sDashboard "github.com/capstone-kelompok-7/backend-disappear/module/feature/dashboard/service"
	hFcm "github.com/capstone-kelompok-7/backend-disappear/module/feature/fcm/handler"
	rFcm "github.com/capstone-kelompok-7/backend-disappear/module/feature/fcm/repository"
	sFcm "github.com/capstone-kelompok-7/backend-disappear/module/feature/fcm/service"
	hHome "github.com/capstone-kelompok-7/backend-disappear/module/feature/homepage/handler"
	rHome "github.com/capstone-kelompok-7/backend-disappear/module/feature/homepage/repository"
	sHome "github.com/capstone-kelompok-7/backend-disappear/module/feature/homepage/service"
	hOrder "github.com/capstone-kelompok-7/backend-disappear/module/feature/order/handler"
	rOrder "github.com/capstone-kelompok-7/backend-disappear/module/feature/order/repository"
	sOrder "github.com/capstone-kelompok-7/backend-disappear/module/feature/order/service"
	"github.com/capstone-kelompok-7/backend-disappear/module/feature/product/handler"
	"github.com/capstone-kelompok-7/backend-disappear/module/feature/product/repository"
	"github.com/capstone-kelompok-7/backend-disappear/module/feature/product/service"
	hReview "github.com/capstone-kelompok-7/backend-disappear/module/feature/review/handler"
	rReview "github.com/capstone-kelompok-7/backend-disappear/module/feature/review/repository"
	sReview "github.com/capstone-kelompok-7/backend-disappear/module/feature/review/service"
	hUser "github.com/capstone-kelompok-7/backend-disappear/module/feature/users/handler"
	rUser "github.com/capstone-kelompok-7/backend-disappear/module/feature/users/repository"
	sUser "github.com/capstone-kelompok-7/backend-disappear/module/feature/users/service"
	hVoucher "github.com/capstone-kelompok-7/backend-disappear/module/feature/voucher/handler"
	rVoucher "github.com/capstone-kelompok-7/backend-disappear/module/feature/voucher/repository"
	sVoucher "github.com/capstone-kelompok-7/backend-disappear/module/feature/voucher/service"
	"github.com/capstone-kelompok-7/backend-disappear/utils/caching/redis"
	"github.com/capstone-kelompok-7/backend-disappear/utils/payment"
	"github.com/sashabaranov/go-openai"

	"net/http"

	"github.com/capstone-kelompok-7/backend-disappear/module/middlewares"
	"github.com/capstone-kelompok-7/backend-disappear/routes"
	"github.com/capstone-kelompok-7/backend-disappear/utils"
	"github.com/capstone-kelompok-7/backend-disappear/utils/database"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	e := echo.New()
	var initConfig = config.InitConfig()

	//db := database.InitDatabase(*initConfig)
	//db := database.InitPGDatabase(*initConfig)
	db := database.InitPGSDatabase(*initConfig)
	rdb := redis.NewRedisClient(*initConfig)
	database.Migrate(db)
	err := seeder.DBSeed(db)
	if err != nil {
		return
	}
	jwtService := utils.NewJWT(initConfig.Secret)
	hash := utils.NewHash()
	generatorID := utils.NewGeneratorUUID(db)
	fcm := sendnotif.NewFcmService()
	coreApi := payment.InitSnapMidtrans(*initConfig)
	emailSender := email.NewEmailService()

	userRepo := rUser.NewUserRepository(db)
	userService := sUser.NewUserService(userRepo, hash)
	userHandler := hUser.NewUserHandler(userService)

	authRepo := rAuth.NewAuthRepository(db)
	authService := sAuth.NewAuthService(authRepo, jwtService, userService, hash, rdb, emailSender)
	authHandler := hAuth.NewAuthHandler(authService, userService)

	voucherRepo := rVoucher.NewVoucherRepository(db)
	voucherService := sVoucher.NewVoucherService(voucherRepo, userService)
	voucherHandler := hVoucher.NewVoucherHandler(voucherService)

	mgodb := database.InitMongoDB(*initConfig)
	var client = openai.NewClient(initConfig.OpenAiApiKey)
	chatbotRepo := rChatbot.NewAssistantRepository(mgodb, db)
	chatbotService := sChatbot.NewAssistantService(chatbotRepo, client, *initConfig)
	chatbotHandler := hChatbot.NewAssistantHandler(chatbotService)

	productRepo := repository.NewProductRepository(db)
	productService := service.NewProductService(productRepo, chatbotService)
	productHandler := handler.NewProductHandler(productService)

	categoryRepo := rCategory.NewCategoryRepository(db)
	categoryService := sCategory.NewCategoryService(categoryRepo)
	categoryHandler := hCategory.NewCategoryHandler(categoryService)

	articleRepo := rArticle.NewArticleRepository(db)
	articleService := sArticle.NewArticleService(articleRepo)
	articleHandler := hArticle.NewArticleHandler(articleService)

	challengeRepo := rChallenge.NewChallengeRepository(db)
	challengeService := sChallenge.NewChallengeService(challengeRepo, userService)
	challengeHandler := hChallenge.NewChallengeHandler(challengeService)

	carouselRepo := rCarousel.NewCarouselRepository(db)
	carouselService := sCarousel.NewCarouselService(carouselRepo)
	carouselHandler := hCarousel.NewCarouselHandler(carouselService)

	addressRepo := rAddress.NewAddressRepository(db)
	addressService := sAddress.NewAddressService(addressRepo)
	addressHandler := hAddress.NewAddressHandler(addressService)

	reviewRepo := rReview.NewReviewRepository(db)
	reviewService := sReview.NewReviewService(reviewRepo, productService)
	reviewHandler := hReview.NewReviewHandler(reviewService)

	fcmRepo := rFcm.NewFcmRepository(db, fcm)
	fcmService := sFcm.NewFcmService(fcmRepo)
	fcmHandler := hFcm.NewFcmHandler(fcmService)

	cartRepo := rCart.NewCartRepository(db)
	cartService := sCart.NewCartService(cartRepo, productService)
	cartHandler := hCart.NewCartHandler(cartService)

	orderRepo := rOrder.NewOrderRepository(db, coreApi)
	orderService := sOrder.NewOrderService(orderRepo, generatorID, productService,
		voucherService, addressService, userService, cartService, fcmService)
	orderHandler := hOrder.NewOrderHandler(orderService)

	dashboardRepo := rDashboard.NewDashboardRepository(db)
	dashboardService := sDashboard.NewDashboardService(dashboardRepo, rdb)
	dashboardHandler := hDashboard.NewDashboardHandler(dashboardService)

	homeRepo := rHome.NewHomepageRepository(db)
	homeService := sHome.NewHomepageService(homeRepo)
	homeHandler := hHome.NewHomepageHandler(homeService)

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))
	e.Use(middlewares.ConfigureLogging())

	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, Disappear!! 🦄✨🍩")
	})

	routes.RouteUser(e, userHandler, jwtService, userService)
	routes.RouteAuth(e, authHandler, jwtService, userService)
	routes.RouteVoucher(e, voucherHandler, jwtService, userService)
	routes.RouteProduct(e, productHandler, jwtService, userService)
	routes.RouteArticle(e, articleHandler, jwtService, userService)
	routes.RouteChallenge(e, challengeHandler, jwtService, userService)
	routes.RouteCategory(e, categoryHandler, jwtService, userService)
	routes.RouteCarousel(e, carouselHandler, jwtService, userService)
	routes.RouteAddress(e, addressHandler, jwtService, userService)
	routes.RouteReview(e, reviewHandler, jwtService, userService)
	routes.RouteCart(e, cartHandler, jwtService, userService)
	routes.RouteOrder(e, orderHandler, jwtService, userService)
	routes.RouteAssistant(e, chatbotHandler, jwtService, userService)
	routes.RouteDashboard(e, dashboardHandler, jwtService, userService)
	routes.RouteHomepage(e, homeHandler, jwtService, userService)
	routes.RouteFcm(e, fcmHandler, jwtService, userService)
	port := os.Getenv("PORT")
	if port == "" {
		port = "8000"
	}
	e.Logger.Fatalf(e.Start(":" + port).Error())
}
