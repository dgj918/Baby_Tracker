package changing

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Change struct {
	ChangeID int `json:"ChangeID"`
	BabyID int `json:"BabyID"`
	WipesID int `json:"WipeID"`
	DateTime string `json:"DateTime"`
	WipesUsed int `json:"WipesUsed"`
	Notes string `json:"Notes"`
}

func GetChanges(c *gin.Context) {
    data, err := GetChangesData()
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

func GetChange(c *gin.Context) {
  changeID := c.Param("changeID")
  data, err := GetChangeData(changeID)
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

func PostChange(c *gin.Context) {
    var changeData Change
    if err := c.BindJSON(&changeData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateChange(database.DB, changeData)
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

func PutChange(c *gin.Context) {
  var changeData Change
  if err := c.BindJSON(&changeData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", changeData)
	data, err := UpdateChange(changeData)
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

func DeleteChange(c *gin.Context) {
  changeID := c.Param("changeID")
  fmt.Println("Here is the change id: ", changeID)
  result := RemoveChange(changeID)
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

func GetChangesData() ([]Change, error) {
    var changes []Change
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"ChangeID\", \"BabyID\", \"WipesID\", \"DateTime\", \"WipesUsed\", \"Notes\" FROM \"Baby_Tracker\".\"Changing\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var change Change
        if err := rows.Scan(&change.ChangeID, &change.BabyID, &change.WipesID, &change.DateTime, &change.WipesUsed, &change.Notes); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        changes = append(changes, change)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(changes)
    return changes, nil
}

func GetChangeData(changeIDstr string) (Change, error) {
  changeID, err := strconv.Atoi(changeIDstr)
  var change Change
  if err != nil {
    return change, err
  }
  DBErr := database.DB.QueryRow("SELECT \"ChangeID\", \"BabyID\", \"WipesID\", \"DateTime\", \"WipesUsed\", \"Notes\" FROM \"Baby_Tracker\".\"Changing\" WHERE \"ChangeID\" = $1", changeID).Scan(&change.ChangeID, &change.BabyID, &change.WipesID, &change.DateTime, &change.WipesUsed, &change.Notes)
  if DBErr != nil {
      return change, fmt.Errorf("Create Change: %v", DBErr)
  }
  return change, nil
}


func CreateChange(db *sql.DB, change Change) (int, error) {
	fmt.Println("Here are the create change: ", change)
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Changing\" (\"BabyID\", \"WipesID\", \"DateTime\", \"WipesUsed\", \"Notes\") VALUES ($1, $2, $3, $4, $5) RETURNING \"ChangeID\"", change.BabyID, change.WipesID, change.DateTime, change.WipesUsed, change.Notes).Scan(&lastInsertId)
    fmt.Println(err)
	if err != nil {
        return 0, fmt.Errorf("Create Change: %v", err)
    }

    return lastInsertId, nil
}

func UpdateChange(change Change) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Changing\" SET \"BabyID\" = $1, \"WipesID\" = $2, \"DateTime\" = $3, \"WipesUsed\" = $4, \"Notes\" = $5 WHERE \"ChangeID\" = $6 RETURNING \"ChangeID\"", change.BabyID, change.WipesID, change.DateTime, change.WipesUsed, change.Notes, change.ChangeID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Change: %v", err)
	}
	return lastInsertId, nil
}

func RemoveChange(changeIDstr string) int64 {
	changeID, err := strconv.Atoi(changeIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Changing\" WHERE \"ChangeID\" = $1", changeID)
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