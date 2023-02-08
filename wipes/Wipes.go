package wipes


import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Wipe struct {
	WipesID int `json:"WipesID"`
	Brand string `json:"Brand"`
}

func GetWipes(c *gin.Context) {
    data, err := GetWipesData()
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

func GetWipe(c *gin.Context) {
  wipeID := c.Param("wipeID")
  data, err := GetWipeData(wipeID)
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

func PostWipe(c *gin.Context) {
    var wipeData Wipe
    if err := c.BindJSON(&wipeData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateWipe(database.DB, wipeData)
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

func PutWipe(c *gin.Context) {
  var wipeData Wipe
  if err := c.BindJSON(&wipeData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", wipeData)
	data, err := UpdateWipe(wipeData)
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

func DeleteWipe(c *gin.Context) {
  wipeID := c.Param("wipeID")
  fmt.Println("Here is the change id: ", wipeID)
  result := RemoveWipe(wipeID)
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

func GetWipesData() ([]Wipe, error) {
    var wipes []Wipe
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"WipesID\", \"Brand\" FROM \"Baby_Tracker\".\"Wipes\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var wipe Wipe
        if err := rows.Scan(&wipe.WipesID, &wipe.Brand); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        wipes = append(wipes, wipe)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(wipes)
    return wipes, nil
}

func GetWipeData(changeIDstr string) (Wipe, error) {
  wipeID, err := strconv.Atoi(changeIDstr)
  var wipe Wipe
  if err != nil {
    return wipe, err
  }
  DBErr := database.DB.QueryRow("SELECT \"WipesID\", \"Brand\" FROM \"Baby_Tracker\".\"Wipes\" WHERE \"WipesID\" = $1", wipeID).Scan(&wipe.WipesID, &wipe.Brand)
  if DBErr != nil {
      return wipe, fmt.Errorf("Create Wipe: %v", DBErr.Error())
  }
  return wipe, nil
}


func CreateWipe(db *sql.DB, wipe Wipe) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Wipes\" (\"Brand\") VALUES ($1) RETURNING \"WipesID\"", wipe.Brand).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create Wipe: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateWipe(wipe Wipe) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Wipes\" SET \"Brand\" = $1 WHERE \"WipesID\" = $2 RETURNING \"WipesID\"", wipe.Brand, wipe.WipesID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Wipe: %v", err)
	}
	return lastInsertId, nil
}

func RemoveWipe(changeIDstr string) int64 {
	wipeID, err := strconv.Atoi(changeIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Wipes\" WHERE \"WipesID\" = $1", wipeID)
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