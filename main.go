package main

import (
  "log"
  "os"
  "fmt"
  "github.com/gin-gonic/gin"
  "github.com/gin-contrib/cors"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/baby"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/changing"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/wipes"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/formula"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/diapers"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/sleep"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/tummytime"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/feeding"
)

func respondWithError(c *gin.Context, code int, message interface{}) {
  c.AbortWithStatusJSON(code, gin.H{"error": message})
}

func TokenAuthMiddleware() gin.HandlerFunc {
  requiredToken := os.Getenv("API_TOKEN")

  // We want to make sure the token is set, bail if not
  if requiredToken == "" {
    log.Fatal("Please set API_TOKEN environment variable")
  }

  return func(c *gin.Context) {
    token := c.Request.FormValue("api_token")
    
    if (c.Request.URL.Path != "/healthcheck") {
      if token == "" {
        respondWithError(c, 401, "API token required")
        return
      }
  
      if token != requiredToken {
        respondWithError(c, 401, "Invalid API token")
        return
      }
    }
    c.Next()
  }
}

func HealthCheck() gin.HandlerFunc {
  return func(c *gin.Context) {
    fmt.Println(c.Request.URL.Path)
    if (c.Request.URL.Path == "/healthcheck") {
      c.JSON(200, "Health Check")
    }
  }
}
 
func main() {
  database.SetUpDatabaseConnection()
  router := gin.Default()

  router.Use(cors.Default())
  router.Use(HealthCheck())
  // router.Use(TokenAuthMiddleware())

  router.GET("/baby", baby.GetBabies)
  router.GET("/baby/:babyID", baby.GetBaby)
  router.POST("/baby", baby.PostBaby)
  router.PUT("/baby", baby.PutBaby)
  router.DELETE("/baby/:babyID", baby.DeleteBaby)

  router.GET("/changing", changing.GetChanges)
  router.GET("/changing/:changeID", changing.GetChange)
  router.POST("/changing", changing.PostChange)
  router.PUT("/changing", changing.PutChange)
  router.DELETE("/changing/:changeID", changing.DeleteChange)

  router.GET("/wipes", wipes.GetWipes)
  router.GET("/wipes/:wipeID", wipes.GetWipe)
  router.POST("/wipes", wipes.PostWipe)
  router.PUT("/wipes", wipes.PutWipe)
  router.DELETE("/wipes/:wipeID", wipes.DeleteWipe)

  router.GET("/formula", formula.GetAllFormula)
  router.GET("/formula/:formulaID", formula.GetFormula)
  router.POST("/formula", formula.PostFormula)
  router.PUT("/formula", formula.PutFormula)
  router.DELETE("/formula/:formulaID", formula.DeleteFormula)

  router.GET("/diapers", diapers.GetDiapers)
  router.GET("/diapers/:diaperID", diapers.GetDiaper)
  router.POST("/diapers", diapers.PostDiaper)
  router.PUT("/diapers", diapers.PutDiaper)
  router.DELETE("/diapers/:diaperID", diapers.DeleteDiaper)

  router.GET("/sleep", sleep.GetAllSleep)
  router.GET("/sleep/:sleepID", sleep.GetSleep)
  router.POST("/sleep", sleep.PostSleep)
  router.PUT("/sleep", sleep.PutSleep)
  router.DELETE("/sleep/:sleepID", sleep.DeleteSleep)

  router.GET("/tummytime", tummytime.GetAllTummyTime)
  router.GET("/tummytime/:tummytimeID", tummytime.GetTummyTime)
  router.POST("/tummytime", tummytime.PostTummyTime)
  router.PUT("/tummytime", tummytime.PutTummyTime)
  router.DELETE("/tummytime/:tummytimeID", tummytime.DeleteTummyTime)

  router.GET("/feeding", feeding.GetAllFeeding)
  router.GET("/feeding/:feedingID", feeding.GetFeeding)
  router.GET("/feeding/subset/:number", feeding.GetFeedingCurrentSubset)
  router.GET("/feeding/sum/byday/:dayLimit", feeding.GetFeedingSumByDay)
  router.POST("/feeding", feeding.PostFeeding)
  router.PUT("/feeding", feeding.PutFeeding)
  router.DELETE("/feeding/:feedingID", feeding.DeleteFeeding)

  router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}


