package weight

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Weight struct {
	WeightID int `json:"WeightID"`
	BabyID int `json:"BabyID"`
	Weight float64 `json:"Weight"`
	Date string `json:"Date"`
	Age float64 `json:"Age"`
}

func GetWeight(c *gin.Context) {
	babyID := c.Param("babyID")
	data, err := GetWeightData(babyID)
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

func PostWeight(c *gin.Context) {
	var weight Weight
    if err := c.BindJSON(&weight); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := PostWeightData(database.DB, weight)
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

func PutWeight(c *gin.Context) {
	var weight Weight
	if err := c.BindJSON(&weight); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		data, err := PutWeightData(database.DB, weight)
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

func DeleteWeight(c *gin.Context) {
	weightID := c.Param("weightID")
	data := DeleteWeightData(weightID)
	if data > 0 {
	  c.JSON(http.StatusOK, gin.H{
		"data": data,
	  })
	} else {
	  c.JSON(http.StatusOK, gin.H{
		"error": "Failed to delete Weight",
	  })
	}
}

func GetWeightData(babyIDstr string) ([]Weight, error) {
	babyID, err := strconv.Atoi(babyIDstr)
	var weight []Weight
	if err != nil {
		return weight, err
	}
	rows, DBErr := database.DB.Query("SELECT \"WeightID\", \"BabyID\", \"Date\", \"Weight\", \"Age\" FROM \"Baby_Tracker\".\"Weight\" WHERE \"BabyID\" = $1", babyID)
	if DBErr != nil {
		return weight, fmt.Errorf("Error getting weight data: ", DBErr.Error())
	}
	defer rows.Close()
	for rows.Next() {
		var w Weight
		if err := rows.Scan(&w.WeightID, &w.BabyID, &w.Date, &w.Weight, &w.Age); err != nil {
			if err == sql.ErrNoRows {
			
			}
		}
		weight = append(weight, w)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("Error getting weight data: %v", err.Error())
	}
	return weight, err
}

func PostWeightData(db *sql.DB ,weight Weight) (int, error) {
	lastInsertId := 0
	err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Weight\" (\"BabyID\", \"Date\", \"Weight\") VALUES ($1, $2, $3) RETURNING \"WeightID\"", weight.BabyID, weight.Date, weight.Weight).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("Error inserting weight data: ", err.Error())
	}
	return lastInsertId, nil
}

func PutWeightData(db *sql.DB, weight Weight) (int, error) {
	lastInsertId := 0
	err := db.QueryRow("UPDATE \"Baby_Tracker\".\"Weight\" SET \"BabyID\" = $1, \"Date\" = $2, \"Weight\" = $3 WHERE \"WeightID\" = $4 RETURNING \"WeightID\"", weight.BabyID, weight.Date, weight.Weight).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("Error inserting weight data: ", err.Error())
	}
	return lastInsertId, nil
}

func DeleteWeightData(weightIDstr string) int64 {
	weightID, err := strconv.Atoi(weightIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Weight\" WHERE \"WeightID\" = $1", weightID)
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
