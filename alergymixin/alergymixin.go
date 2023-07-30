package alergymixin

import (
	"database/sql"
	"fmt"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
)

// AlergyMixin is a struct that contains the BabyID, AlergyDate, and Mixin as a boolean
type AlergyMixin struct {
	BabyID int    `json:"BabyID"`
	AlergyDate   string `json:"AlergyDate"`
	AlergyMixInEaten  bool   `json:"AlergyMixInEaten"`
}

// GetMixin is a function that returns the AlergyMixin struct for the given BabyID
func GetMixin(c *gin.Context) {
	babyID := c.Param("babyID")
	data, err := GetMixinData(babyID)
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

// PostMixin is a function that inserts a new AlergyMixin struct into the database
func PostMixin(c *gin.Context) {
	var mixin AlergyMixin
	if err := c.BindJSON(&mixin); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
	} else {
		fmt.Printf("mixin Data from Call: %v\n", mixin)
		data, err := PostMixinData(database.DB, mixin)
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

// GetMixinData is a function that returns the AlergyMixin struct for the given BabyID
func GetMixinData(babyIDstr string) (AlergyMixin, error) {
	babyID, err := strconv.Atoi(babyIDstr)
	var mixin AlergyMixin
	if err != nil {
		return mixin, err
	}
	DBErr := database.DB.QueryRow("SELECT \"BabyID\", \"AlergyDate\", \"AlergyMixInEaten\" FROM \"Baby_Tracker\".\"AlergyMixin\" WHERE \"BabyID\" = $1", babyID).Scan(&mixin.BabyID, &mixin.AlergyDate, &mixin.AlergyMixInEaten)
	if DBErr != nil {
		return mixin, fmt.Errorf("Error getting alergy mixin data: ", DBErr.Error())
	}
	return mixin, err
}

// PostMixinData is a function that inserts a new AlergyMixin struct into the database
func PostMixinData(db *sql.DB , mixin AlergyMixin) (int, error) {
	lastInsertId := 0
	fmt.Printf("mixin: %v", mixin)
	err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"AlergyMixin\" (\"BabyID\", \"Date\", \"AlergyMixInEaten\") VALUES ($1, $2, $3) ON CONFLICT(\"BabyID\", \"Date\") DO UPDATE SET \"AlergyMixInEaten\" = $3 RETURNING \"BabyID\"", mixin.BabyID, mixin.AlergyDate, mixin.AlergyMixInEaten).Scan(&lastInsertId)
	if err != nil {
		return 0, fmt.Errorf("Error inserting alergy mixin data: ", err.Error())
	}
	return lastInsertId, nil
}

