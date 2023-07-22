package sleep

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
	"gopkg.in/guregu/null.v4"
)

type Sleep struct {
	SleepID int `json:"SleepID"`
	BabyID int `json:"BabyID"`
	Start string `json:"Start"`
	End null.String `json:"End"`
	WakeUp null.String `json:"WakeUp"`
}

type SleepInProgess struct {
	BabyID int `json:"BabyID"`
	NapInProgress bool `json:"NapInProgress"`
}


func GetAllSleep(c *gin.Context) {
    data, err := GetAllSleepData("10000")
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

func GetSleep(c *gin.Context) {
  sleepID := c.Param("sleepID")
  data, err := GetSleepData(sleepID)
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

func GetInProgressSleep(c *gin.Context) {
	babyID := c.Param("babyID")
	data, err := GetInProgressSleepData(babyID)
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

func GetIsSleepInProgress(c *gin.Context) {
	babyID := c.Param("babyID")
	data, err := IsSleepInProgress(babyID)
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

func PostSleep(c *gin.Context) {
    var sleepData Sleep
    if err := c.BindJSON(&sleepData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateSleep(database.DB, sleepData)
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

func PutSleep(c *gin.Context) {
  var sleepData Sleep
  if err := c.BindJSON(&sleepData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", sleepData)
	data, err := UpdateSleep(sleepData)
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

func PutStartSleep(c *gin.Context) {
	babyID := c.Param("babyID")
	var sleepInProgress SleepInProgess
	sleepInProgress.BabyID, _ = strconv.Atoi(babyID)
	sleepInProgress.NapInProgress = true
	data, err := UpdateCurrentSleepStatus(sleepInProgress)
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

func PutStopSleep(c *gin.Context) {
	babyID := c.Param("babyID")
	var sleepInProgress SleepInProgess
	sleepInProgress.BabyID, _ = strconv.Atoi(babyID)
	sleepInProgress.NapInProgress = false
	data, err := UpdateCurrentSleepStatus(sleepInProgress)
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

func DeleteSleep(c *gin.Context) {
  sleepID := c.Param("sleepID")
  fmt.Println("Here is the change id: ", sleepID)
  result := RemoveSleep(sleepID)
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

func GetSleepCurrentSubset(c *gin.Context) {
	num := c.Param("num")
	result, err := GetAllSleepData(num)
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

/*
 * Database Functions
*/

func GetAllSleepData(num string) ([]Sleep, error) {
    var sleeps []Sleep
	numReturn, err := strconv.Atoi(num)
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"SleepID\", \"BabyID\", \"Start\", \"End\", \"WakeUp\" FROM \"Baby_Tracker\".\"Sleep\" ORDER BY \"Start\" desc limit $1", numReturn)
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var sleep Sleep
        if err := rows.Scan(&sleep.SleepID, &sleep.BabyID, &sleep.Start, &sleep.End, &sleep.WakeUp); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        sleeps = append(sleeps, sleep)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(sleeps)
    return sleeps, nil
}

func GetSleepData(sleepIDstr string) (Sleep, error) {
  sleepID, err := strconv.Atoi(sleepIDstr)
  var sleep Sleep
  if err != nil {
    return sleep, err
  }
  DBErr := database.DB.QueryRow("SELECT \"SleepID\", \"BabyID\", \"Start\", \"End\", \"WakeUp\" FROM \"Baby_Tracker\".\"Sleep\" WHERE \"SleepID\" = $1", sleepID).Scan(&sleep.SleepID, &sleep.BabyID, &sleep.Start, &sleep.End, &sleep.WakeUp)
  if DBErr != nil {
      return sleep, fmt.Errorf("Create Sleep: %v", DBErr.Error())
  }
  return sleep, nil
}

func IsSleepInProgress(babyIDstr string) (SleepInProgess, error) {
	babyID, err := strconv.Atoi(babyIDstr)
	var sleepInProgess SleepInProgess
	if err != nil {
		return sleepInProgess, err
	}	
	DBErr := database.DB.QueryRow("SELECT \"BabyID\", \"NapInProgress\" FROM \"Baby_Tracker\".\"SleepInProgress\" WHERE \"BabyID\" = $1", babyID).Scan(&sleepInProgess.BabyID, &sleepInProgess.NapInProgress)
	if DBErr != nil {
		return sleepInProgess, fmt.Errorf("Database Error: %v", DBErr.Error())
	}
	return sleepInProgess, nil
}

func GetInProgressSleepData(babyIDstr string) (Sleep, error) {
	babyID, err := strconv.Atoi(babyIDstr)
	var sleep Sleep
	if err != nil {
		return sleep, err
	}	
	DBErr := database.DB.QueryRow("SELECT \"SleepID\", \"BabyID\", \"Start\", \"End\", \"WakeUp\" FROM \"Baby_Tracker\".\"Sleep\" WHERE \"BabyID\" = $1 AND \"End\" IS NULL", babyID).Scan(&sleep.SleepID, &sleep.BabyID, &sleep.Start, &sleep.End, &sleep.WakeUp)
	if DBErr != nil {
		return sleep, fmt.Errorf("Database Error: %v", DBErr.Error())
	}
	return sleep, nil
}


func CreateSleep(db *sql.DB, sleep Sleep) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Sleep\" (\"BabyID\", \"Start\", \"End\", \"WakeUp\") VALUES ($1, $2, $3, $4) RETURNING \"SleepID\"", sleep.BabyID, sleep.Start, sleep.End, sleep.WakeUp).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create Sleep: %v", err.Error())
    }

    return lastInsertId, nil
}
 

func UpdateSleep(sleep Sleep) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Sleep\" SET \"BabyID\" = $1, \"Start\" = $2, \"End\" = $3, \"WakeUp\" = $4 WHERE \"SleepID\" = $5 RETURNING \"SleepID\"", sleep.BabyID, sleep.Start, sleep.End.String, sleep.WakeUp, sleep.SleepID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Sleep: %v", err)
	}
	return lastInsertId, nil
}

func UpdateCurrentSleepStatus(sleepStatus SleepInProgess) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"SleepInProgress\" SET \"NapInProgress\" = $1 WHERE \"BabyID\" = $2 RETURNING \"BabyID\"", sleepStatus.NapInProgress, sleepStatus.BabyID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Updated Sleep Status: %v", err)
	}
	return lastInsertId, nil
}

func RemoveSleep(sleepIDstr string) int64 {
	sleepID, err := strconv.Atoi(sleepIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Sleep\" WHERE \"SleepID\" = $1", sleepID)
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