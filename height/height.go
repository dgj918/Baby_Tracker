package height

import (
	"database/sql"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Height struct {
	HeightID int `json:"HeightID"`
	BabyID int `json:"BabyID"`
	Height float64 `json:"Height"`
	Date string `json:"Date"`
}

func GetHeight(c *gin.Context) {
	babyID := c.Param("babyID")
	fmt.Printf("babyID: %s", babyID)
	data, err := GetHeightData(babyID)
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

func PostHeight(c *gin.Context) {
	var height Height
    if err := c.BindJSON(&height); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := PostHeightData(database.DB, height)
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

func PutHeight(c *gin.Context) {
	var height Height
	if err := c.BindJSON(&height); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		data, err := PutHeightData(database.DB, height)
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

func DeleteHeight(c *gin.Context) {
	heightID := c.Param("heightID")
	data := DeleteHeightData(heightID)
	if data > 0 {
	  c.JSON(http.StatusOK, gin.H{
		"data": data,
	  })
	} else {
	  c.JSON(http.StatusOK, gin.H{
		"error": "Failed to delete Height",
	  })
	}
}

func GetHeightData(babyIDstr string) (Height, error) {
	fmt.Printf("babyIDstr: %v", babyIDstr)
	babyID, err := strconv.Atoi(babyIDstr)
	fmt.Printf("babyID: %v", babyID)
	var height Height
	if err != nil {
		return height, err
	}
	DBErr := database.DB.QueryRow("SELECT \"HeightID\", \"BabyID\", \"Date\", \"Height\" FROM \"Height\" WHERE \"HeightID\" = $1", babyID).Scan(&height.HeightID, &height.BabyID, &height.Date, &height.Height)
	if DBErr != nil {
		return height, fmt.Errorf("Error getting height data: ", DBErr.Error())
	}
	return height, err
}

func PostHeightData(db *sql.DB ,height Height) (int, error) {
	lastInsertId := 0
	err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Height\" (\"BabyID\", \"Date\", \"Height\") VALUES ($1, $2, $3) RETURNING \"HeightID\"", height.BabyID, height.Date, height.Height).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("Error inserting height data: ", err.Error())
	}
	return lastInsertId, nil
}

func PutHeightData(db *sql.DB, height Height) (int, error) {
	lastInsertId := 0
	err := db.QueryRow("UPDATE \"Baby_Tracker\".\"Height\" SET \"BabyID\" = $1, \"Date\" = $2, \"Height\" = $3 WHERE \"HeightID\" = $4 RETURNING \"HeightID\"", height.BabyID, height.Date, height.Height).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("Error inserting height data: ", err.Error())
	}
	return lastInsertId, nil
}

func DeleteHeightData(heightIDstr string) int64 {
	heightID, err := strconv.Atoi(heightIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Height\" WHERE \"HeightID\" = $1", heightID)
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
