
	package main

	import(
		"log"
		"os"
		"github.com/gin-gonic/gin"
		"github.com/joho/godotenv"

		"github.com/fxivan/commitnamegen_ia/cmd/api/handlers/commit"
		"github.com/fxivan/commitnamegen_ia/internal/repositories/mongo"
		commitMongo "github.com/fxivan/commitnamegen_ia/internal/repositories/mongo/commit"
		commitService "github.com/fxivan/commitnamegen_ia/internal/services/commit"
	)

	func main() {
		err := godotenv.Load()
		if err != nil {
			log.Fatal("Error loading .env file")
		}

		ginEngine := gin.Default()

		client, err := mongo.ConnectClient(os.Getenv("MONGO_URI"))
		if err != nil {
			log.Fatal(err.Error())
		}

		commitRepo := commitMongo.Repository{
			Client: client,
		}

		commitSrv := commitService.Service{
			Repo: commitRepo,
		}

		commitHandler := commit.Handler{
			CommitService: commitSrv,
		}

		ginEngine.POST("/commit", commitHandler.CreateCommit)

		log.Fatal(ginEngine.Run(":8001"))
	}
