package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	api "github.com/udodinho/golangProjects/wallet-engine/internals/adapters/api/wallet"
	repository "github.com/udodinho/golangProjects/wallet-engine/internals/adapters/repository/mongodb"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/helpers"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/services"
	"github.com/udodinho/golangProjects/wallet-engine/internals/core/shared"
	"github.com/udodinho/golangProjects/wallet-engine/internals/ports"
	"log"
	"time"
)

func main() {
	helpers.InitializeLogDir()

	//mongoURL := "mongodb://localhost:27017"
	serviceAddress, servicePort, dbType, mongodbPort, _, mongodbDbHost, dbName, _ := helpers.LoadConfig()
	mongoURL := fmt.Sprintf("%s://%s:%s", dbType, mongodbDbHost, mongodbPort)
	dbRepository := ConnectToMongo(mongoURL, dbName)
	service := services.New(dbRepository)
	handler := api.NewHttpHandler(service)
	router := gin.Default()
	//router.Use(helpers.LogRequest)

	apirouter := router.Group("/api/v1")
	apirouter.POST("/createWallet", handler.CreateWallet())
	apirouter.POST("/creditWallet/:reference", handler.CreditWallet())
	apirouter.POST("/debitWallet/:reference", handler.DebitWallet())
	apirouter.PUT("/activateDeactivateWallet/:reference", handler.ActivateDeactivate())

	router.NoRoute(func(ctx *gin.Context) {
		ctx.JSON(404,
			helpers.PrintErrorMessage("404", shared.NoResourceFound))
	})

	fmt.Println("service running on " + serviceAddress + ":" + servicePort)
	helpers.LogEvent("Info", fmt.Sprintf("started wallet-engine application on "+serviceAddress+":"+servicePort+" in "+time.Since(time.Now()).String()))
	_ = router.Run(":" + servicePort)

}

func ConnectToMongo(mongoURL string, dbName string) ports.WalletRepository {
	repo, err := repository.NewMongoRepository(mongoURL, dbName, 2000)
	if err != nil {
		_ = helpers.PrintErrorMessage("500", err.Error())
		log.Fatal(err)
	}
	return services.New(repo)
}
