package diapers


import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Diapers struct {
	DiaperID int `json:"DiaperID"`
	Brand string `json:"Brand"`
}

func GetDiapers(c *gin.Context) {
    data, err := GetFormulaData()
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

func GetDiaper(c *gin.Context) {
  formulaID := c.Param("formulaID")
  data, err := GetDiaperData(formulaID)
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

func PostDiaper(c *gin.Context) {
    var diaperData Diapers
    if err := c.BindJSON(&diaperData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateDiaper(database.DB, diaperData)
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

func PutDiaper(c *gin.Context) {
  var diaperData Diapers
  if err := c.BindJSON(&diaperData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", diaperData)
	data, err := UpdateDiaper(diaperData)
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

func DeleteDiaper(c *gin.Context) {
  formulaID := c.Param("formulaID")
  fmt.Println("Here is the change id: ", formulaID)
  result := RemoveDiaper(formulaID)
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

func GetFormulaData() ([]Diapers, error) {
    var formulas []Diapers
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"DiaperID\", \"Brand\" FROM \"Baby_Tracker\".\"Diapers\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var diaper Diapers
        if err := rows.Scan(&diaper.DiaperID, &diaper.Brand); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        formulas = append(formulas, diaper)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(formulas)
    return formulas, nil
}

func GetDiaperData(formulaIDstr string) (Diapers, error) {
  formulaID, err := strconv.Atoi(formulaIDstr)
  var diaper Diapers
  if err != nil {
    return diaper, err
  }
  DBErr := database.DB.QueryRow("SELECT \"DiaperID\", \"Brand\" FROM \"Baby_Tracker\".\"Diapers\" WHERE \"DiaperID\" = $1", formulaID).Scan(&diaper.DiaperID, &diaper.Brand)
  if DBErr != nil {
      return diaper, fmt.Errorf("Create Diapers: %v", DBErr.Error())
  }
  return diaper, nil
}


func CreateDiaper(db *sql.DB, diaper Diapers) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Diapers\" (\"Brand\") VALUES ($1) RETURNING \"DiaperID\"", diaper.Brand).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create Diapers: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateDiaper(diaper Diapers) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Diapers\" SET \"Brand\" = $1 WHERE \"DiaperID\" = $2 RETURNING \"DiaperID\"", diaper.Brand, diaper.DiaperID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Diapers: %v", err)
	}
	return lastInsertId, nil
}

func RemoveDiaper(formulaIDstr string) int64 {
	formulaID, err := strconv.Atoi(formulaIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Diapers\" WHERE \"DiaperID\" = $1", formulaID)
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