package baby

import (
	"database/sql"
	"log"
	"fmt"
  "github.com/gin-gonic/gin"
  "net/http"
  "gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
  "strconv"
)

type Baby struct {
	BabyID int `json:"BabyID"`
	FirstName string `json:"FirstName"`
	LastName string `json:"LastName"`
	Birthday string `json:"Birthday"`
}

func GetBabies(c *gin.Context) {
    data, err := GetAllBabies()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err,
      })
    }
    c.JSON(http.StatusOK, gin.H{
      "data": data,
    })
}

func GetBaby(c *gin.Context) {
  babyID := c.Param("babyID")
  fmt.Println(babyID)
  data, err := GetBabyData(babyID)
  fmt.Println(data, err)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err,
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "data": data,
  })
}

func PostBaby(c *gin.Context) {
    var babyData Baby
    if err := c.BindJSON(&babyData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err,
      })
    }
    data, err := CreateBaby(database.DB, babyData)
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err,
      })
    }
    c.JSON(http.StatusOK, gin.H{
      "message": data,
    })
}

func PutBaby(c *gin.Context) {
  var babyData Baby
  if err := c.BindJSON(&babyData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err,
    })
  }
  data, err := UpdateBaby(babyData)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err,
    })
  }
  c.JSON(http.StatusOK, gin.H{
    "message": data,
  })
}

func DeleteBaby(c *gin.Context) {
  babyID := c.Param("babyID")
  fmt.Println("Here is the baby id: ", babyID)
  sucess := RemoveBaby(babyID)
  if sucess {
    c.JSON(http.StatusOK, gin.H{})
  }
  c.JSON(http.StatusInternalServerError, gin.H{
    "error": "Delete Failed",
  })
}

func GetAllBabies() ([]Baby, error) {
    var babys []Baby
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"BabyID\", \"FirstName\", \"LastName\", \"Birthday\" FROM \"Baby_Tracker\".\"Baby\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var baby Baby
        if err := rows.Scan(&baby.BabyID, &baby.FirstName, &baby.LastName, &baby.Birthday); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
		fmt.Println(baby)
        babys = append(babys, baby)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(babys)
    return babys, nil
}

func GetBabyData(babyIDstr string) (Baby, error) {
  babyID, err := strconv.Atoi(babyIDstr)
  var baby Baby
  fmt.Println(babyID, babyIDstr)
  if err != nil {
    return baby, err
  }
  DBErr := database.DB.QueryRow("SELECT \"BabyID\", \"FirstName\", \"LastName\", \"Birthday\" FROM \"Baby_Tracker\".\"Baby\" WHERE \"BabyID\" = $1", babyID).Scan(&baby.BabyID, &baby.FirstName, &baby.LastName, &baby.Birthday)
  fmt.Println(baby)
  if DBErr != nil {
      return baby, fmt.Errorf("Create Baby: %v", DBErr)
  }
  return baby, nil
}


func CreateBaby(db *sql.DB, baby Baby) (int, error) {
    fmt.Println("Baby Data: ", baby.FirstName, baby.LastName, baby.Birthday)
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Baby\" (\"FirstName\", \"LastName\", \"Birthday\") VALUES ($1, $2, $3) RETURNING \"BabyID\"", baby.FirstName, baby.LastName, baby.Birthday).Scan(&lastInsertId)
    if err != nil {
        return 0, fmt.Errorf("Create Baby: %v", err)
    }

    return lastInsertId, nil
}

func UpdateBaby(baby Baby) (int, error) {
  lastInsertId := 0
  err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Baby\" SET \"FirstName\" = $1, \"LastName\" = $2, \"Birthday\" = $3 WHERE \"BabyID\" = $4 RETURNING \"BabyID\"", baby.FirstName, baby.LastName, baby.Birthday, baby.BabyID).Scan(&lastInsertId)
  if err != nil {
      return 0, fmt.Errorf("Update Baby: %v", err)
  }
  return lastInsertId, nil
}

func RemoveBaby(babyIDstr string) bool {
  babyID, err := strconv.Atoi(babyIDstr)
  if err != nil {
    return false
  }
  dbErr := database.DB.QueryRow("DELETE FROM \"Baby_Tracker\".\"Baby\" WHERE \"BabyID\" = $1", babyID)
  if dbErr != nil {
    return false
  }
  return true
}