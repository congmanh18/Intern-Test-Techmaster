package server

import (
	"03/handler"
	"03/migrate"
	"03/model"
	"03/provider"
	"03/repository"
	"03/utils/config"
	"03/utils/groq"
	"log"

	"github.com/kataras/iris/v12"
	"github.com/russross/blackfriday/v2"
)

func RunMigration(appProvider *provider.AppProvider, enableMigrate bool) {
	if enableMigrate {
		migrate.Migration(appProvider.Postgres.Executor)
	}
}

func LoadServiceConfig(confPath string) config.ServiceConfig {
	var serviceConf config.ServiceConfig
	config.MustLoadConfig(confPath, &serviceConf)
	return serviceConf
}
func Run(confPath string) {
	// Load configuration
	serviceConf := LoadServiceConfig(confPath)
	appProvider := provider.NewAppProvider(serviceConf)

	// Run migration if not already
	RunMigration(appProvider, serviceConf.EnableMigrations)

	dialogRepo := repository.NewDialogRepo(appProvider.Postgres)
	wordRepo := repository.NewWordRepo(appProvider.Postgres)
	wordDialogRepo := repository.NewWordDialogRepo(appProvider.Postgres)

	handler := handler.NewHandler(
		handler.HandlerInject{
			DialogRepo:     dialogRepo,
			WordRepo:       wordRepo,
			WordDialogRepo: wordDialogRepo,
		},
	)

	app := iris.New()
	// Serve static files
	app.HandleDir("/static", "./static")

	// Trang chính
	app.Get("/", func(ctx iris.Context) {
		ctx.View("index.html")
	})

	// API xử lý prompt
	app.Post("/ask", func(ctx iris.Context) {
		var req model.Request
		if err := ctx.ReadJSON(&req); err != nil {
			ctx.StatusCode(iris.StatusBadRequest)
			ctx.JSON(iris.Map{"error": "Lỗi khi đọc dữ liệu"})
			return
		}

		response, err := groq.CallGroqAPI(serviceConf.GroqUrl, serviceConf.GroqAPIKey, req.Prompt, req.Model)
		if err != nil {
			ctx.StatusCode(iris.StatusInternalServerError)
			ctx.JSON(iris.Map{"error": err.Error()})
			return
		}

		// Chuyển Markdown thành HTML
		htmlResponse := string(blackfriday.Run([]byte(response)))

		ctx.JSON(iris.Map{"response": htmlResponse})
	})

	// Định nghĩa API endpoint
	app.Post("/auto-process", handler.AutoProcessHandler)

	app.RegisterView(iris.HTML("./templates", ".html"))
	log.Println("🚀 Server đang chạy tại http://localhost:8080")
	app.Listen(":8080")
}
