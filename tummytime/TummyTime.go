package tummytime

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type TummyTime struct {
	TummyTimeID int `json:"TummyTimeID"`
	BabyID int `json:"BabyID"`
	Start string `json:"Start"`
	End string `json:"End"`
	Location string `json:"Location"`
	Position string `json:"Position"`
	Temperament string `json:"Temperament"`
	Notes string `json:"Notes"`
}

func GetAllTummyTime(c *gin.Context) {
    data, err := GetAllTummyTimeData()
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

func GetTummyTime(c *gin.Context) {
  tummyTimeID := c.Param("tummyTimeID")
  data, err := GetSleepData(tummyTimeID)
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

func PostTummyTime(c *gin.Context) {
    var tummyTimeData TummyTime
    if err := c.BindJSON(&tummyTimeData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateTummyTime(database.DB, tummyTimeData)
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

func PutTummyTime(c *gin.Context) {
  var tummyTimeData TummyTime
  if err := c.BindJSON(&tummyTimeData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", tummyTimeData)
	data, err := UpdateTummyTime(tummyTimeData)
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

func DeleteTummyTime(c *gin.Context) {
  tummyTimeID := c.Param("tummytimeID")
  fmt.Println("Here is the change id: ", tummyTimeID)
  result := RemoveTummyTime(tummyTimeID)
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

func GetAllTummyTimeData() ([]TummyTime, error) {
    var tummyTimes []TummyTime
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"TummyTimeID\", \"BabyID\", \"Start\", \"End\", \"Location\", \"Position\", \"Temperament\", \"Notes\" FROM \"Baby_Tracker\".\"TummyTime\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var tummyTime TummyTime
        if err := rows.Scan(&tummyTime.TummyTimeID, &tummyTime.BabyID, &tummyTime.Start, &tummyTime.End, &tummyTime.Location, &tummyTime.Temperament, &tummyTime.Notes); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        tummyTimes = append(tummyTimes, tummyTime)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(tummyTimes)
    return tummyTimes, nil
}

func GetSleepData(tummyTimeIDstr string) (TummyTime, error) {
  tummyTimeID, err := strconv.Atoi(tummyTimeIDstr)
  var tummyTime TummyTime
  if err != nil {
    return tummyTime, err
  }
  DBErr := database.DB.QueryRow("SELECT \"TummyTimeID\", \"BabyID\", \"Start\", \"End\", \"Location\", \"Position\", \"Temperament\", \"Notes\" FROM \"Baby_Tracker\".\"TummyTime\" WHERE \"TummyTimeID\" = $1", tummyTimeID).Scan(&tummyTime.TummyTimeID, &tummyTime.BabyID, &tummyTime.Start, &tummyTime.End, &tummyTime.Location, &tummyTime.Position, &tummyTime.Temperament, &tummyTime.Notes)
  if DBErr != nil {
      return tummyTime, fmt.Errorf("Create TummyTime: %v", DBErr.Error())
  }
  return tummyTime, nil
}


func CreateTummyTime(db *sql.DB, tummyTime TummyTime) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"TummyTime\" (\"BabyID\", \"Start\", \"End\", \"Location\", \"Position\", \"Temperament\", \"Notes\") VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING \"TummyTimeID\"", tummyTime.BabyID, tummyTime.Start, tummyTime.End, tummyTime.Location, tummyTime.Position, tummyTime.Temperament, tummyTime.Notes).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create TummyTime: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateTummyTime(tummyTime TummyTime) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"TummyTime\" SET \"BabyID\" = $1, \"Start\" = $2, \"End\" = $3, \"Location\" = $4, \"Position\" = $5, \"Temperament\" = $6, \"Notes\" = $7 WHERE \"TummyTimeID\" = $8 RETURNING \"TummyTimeID\"", tummyTime.BabyID, tummyTime.Start, tummyTime.End, tummyTime.Location, tummyTime.Position, tummyTime.Temperament, tummyTime.Notes, tummyTime.TummyTimeID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update TummyTime: %v", err)
	}
	return lastInsertId, nil
}

func RemoveTummyTime(tummyTimeIDstr string) int64 {
	tummyTimeID, err := strconv.Atoi(tummyTimeIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"TummyTime\" WHERE \"TummyTimeID\" = $1", tummyTimeID)
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