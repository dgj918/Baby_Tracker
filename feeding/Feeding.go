package feeding

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Feeding struct {
	FeedingID int `json:"FeedingID"`
	BabyID int `json:"BabyID"`
	Start string `json:"Start"`
	End string `json:"End"`
	OuncesConsumed int `json:"OuncesConsumed"`
	FormulaID int `json:"FormulaID"`
}

type FeedingDisplay struct {
	FeedingID int `json:"FeedingID"`
	BabyID int `json:"BabyID"`
	Start string `json:"Start"`
	End string `json:"End"`
	OuncesConsumed int `json:"OuncesConsumed"`
	FormulaID int `json:"FormulaID"`
	Brand string `json:"Brand"`
}

type FeedingSumByDay struct {
	Day string `json:"Day"`
	MlPerDay int `json:"MlPerDay"`
	BabyID int `json:"BabyID"`
}


func GetAllFeeding(c *gin.Context) {
    data, err := GetAllFeedingData()
    if err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		c.JSON(http.StatusOK, gin.H{
			"data": data,
		})
	}
}

func GetFeeding(c *gin.Context) {
  feedingID := c.Param("feedingID")
  data, err := GetFeedingData(feedingID)
  fmt.Println("Here is the return data: ", data, err)
  if err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	c.JSON(http.StatusOK, gin.H{
		"data": data,
	  })
  }
}

func GetFeedingSumByDay(c *gin.Context) {
	dayLimit := c.Param("dayLimit")
	data, err := GetFeedingSumByDayData(dayLimit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H {
			"data": data,
		})
	}
}

func PostFeeding(c *gin.Context) {
    var feedingData Feeding
    if err := c.BindJSON(&feedingData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateFeeding(database.DB, feedingData)
		if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": data,
			})
		}	
	}
}

func PutFeeding(c *gin.Context) {
  var feedingData Feeding
  if err := c.BindJSON(&feedingData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", feedingData)
	data, err := UpdateFeeding(feedingData)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": data,
		})
	}
  }
}

func DeleteFeeding(c *gin.Context) {
  feedingID := c.Param("feedingID")
  fmt.Println("Here is the change id: ", feedingID)
  result := RemoveFeeding(feedingID)
  if result > 0  {
    c.JSON(http.StatusOK, gin.H{
		"data": "Rows Deleted",
	})
  } else {
	c.JSON(http.StatusInternalServerError, gin.H{
		"error": "Delete Failed",
	  })
  }
}

func GetFeedingCurrentSubset(c *gin.Context) {
	num := c.Param("number")
	result, err := GetFeedingCurrentSubsetData(num)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
		  "error": err.Error(),
		})
	  } else {
		c.JSON(http.StatusOK, gin.H{
			"data": result,
		  })
	  }
  }

func GetAllFeedingData() ([]FeedingDisplay, error) {
    var feedings []FeedingDisplay
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"FeedingID\", \"BabyID\", \"Start\", \"End\", \"OuncesConsumed\", \"Feeding\".\"FormulaID\", \"Formula\".\"Brand\" FROM \"Baby_Tracker\".\"Feeding\" JOIN \"Baby_Tracker\".\"Formula\" ON \"Feeding\".\"FormulaID\" = \"Formula\".\"FormulaID\" ORDER BY \"Start\" desc LIMIT 100")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var feeding FeedingDisplay
        if err := rows.Scan(&feeding.FeedingID, &feeding.BabyID, &feeding.Start, &feeding.End, &feeding.OuncesConsumed, &feeding.FormulaID, &feeding.Brand); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        feedings = append(feedings, feeding)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(feedings)
    return feedings, nil
}

func GetFeedingCurrentSubsetData(num string) ([]FeedingDisplay, error) {
    var feedings []FeedingDisplay
	numReturn, err := strconv.Atoi(num)
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"FeedingID\", \"BabyID\", \"Start\", \"End\", \"OuncesConsumed\", \"Feeding\".\"FormulaID\", \"Formula\".\"Brand\" FROM \"Baby_Tracker\".\"Feeding\" JOIN \"Baby_Tracker\".\"Formula\" ON \"Feeding\".\"FormulaID\" = \"Formula\".\"FormulaID\" ORDER BY \"Start\" desc limit $1", numReturn)
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var feeding FeedingDisplay
        if err := rows.Scan(&feeding.FeedingID, &feeding.BabyID, &feeding.Start, &feeding.End, &feeding.OuncesConsumed, &feeding.FormulaID, &feeding.Brand); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        feedings = append(feedings, feeding)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(feedings)
    return feedings, nil
}

func GetFeedingSumByDayData(dayLimit string) ([]FeedingSumByDay, error){
	var feedings []FeedingSumByDay
	dayLimitReturn, err := strconv.Atoi(dayLimit)
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query(`SELECT date_trunc('day', "Start") "Day", SUM("OuncesConsumed") as "MlPerDay", "BabyID" FROM "Baby_Tracker"."Feeding" GROUP BY 1, 3 ORDER BY "Day" desc limit $1`, dayLimitReturn)
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var feeding FeedingSumByDay
        if err := rows.Scan(&feeding.Day, &feeding.MlPerDay, &feeding.BabyID); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        feedings = append(feedings, feeding)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(feedings)
    return feedings, nil
}

func GetFeedingData(feedingIDstr string) (Feeding, error) {
  feedingID, err := strconv.Atoi(feedingIDstr)
  var feeding Feeding
  if err != nil {
    return feeding, err
  }
  DBErr := database.DB.QueryRow("SELECT \"FeedingID\", \"BabyID\", \"Start\", \"End\", \"OuncesConsumed\", \"FormulaID\" FROM \"Baby_Tracker\".\"Feeding\" WHERE \"FeedingID\" = $1", feedingID).Scan(&feeding.FeedingID, &feeding.BabyID, &feeding.Start, &feeding.End, &feeding.OuncesConsumed, &feeding.FormulaID)
  if DBErr != nil {
      return feeding, fmt.Errorf("Create Feeding: %v", DBErr.Error())
  }
  return feeding, nil
}


func CreateFeeding(db *sql.DB, feeding Feeding) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Feeding\" (\"BabyID\", \"Start\", \"End\", \"OuncesConsumed\", \"FormulaID\") VALUES ($1, $2, $3, $4, $5) RETURNING \"FeedingID\"", feeding.BabyID, feeding.Start, feeding.End, feeding.OuncesConsumed, feeding.FormulaID).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create Feeding: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateFeeding(feeding Feeding) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Feeding\" SET \"BabyID\" = $1, \"Start\" = $2, \"End\" = $3, \"OuncesConsumed\" = $4, \"FormulaID\" = $5 WHERE \"FeedingID\" = $6 RETURNING \"FeedingID\"", feeding.BabyID, feeding.Start, feeding.End, feeding.OuncesConsumed, feeding.FormulaID, feeding.FeedingID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Feeding: %v", err)
	}
	return lastInsertId, nil
}

func RemoveFeeding(feedingIDstr string) int64 {
	feedingID, err := strconv.Atoi(feedingIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Feeding\" WHERE \"FeedingID\" = $1", feedingID)
	fmt.Println(dbErr)	
	if dbErr != nil {
		fmt.Print("Err is not nil")
		return -1
	}

	count, delErr := res.RowsAffected()
	if delErr != nil {
		return -1
	}
	fmt.Print("Err is nil")
	return count
}