package formula


import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type Formula struct {
	FormulaID int `json:"FormulaID"`
	Brand string `json:"Brand"`
}

func GetAllFormula(c *gin.Context) {
    data, err := GetAllFormulaData()
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

func GetFormula(c *gin.Context) {
  formulaID := c.Param("formulaID")
  data, err := GetFormulaData(formulaID)
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

func PostFormula(c *gin.Context) {
    var formulaData Formula
    if err := c.BindJSON(&formulaData); err != nil {
      c.JSON(http.StatusInternalServerError, gin.H{
        "error": err.Error(),
      })
    } else {
		data, err := CreateWipe(database.DB, formulaData)
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

func PutFormula(c *gin.Context) {
  var formulaData Formula
  if err := c.BindJSON(&formulaData); err != nil {
    c.JSON(http.StatusInternalServerError, gin.H{
      "error": err.Error(),
    })
  } else {
	fmt.Println("Here is the params: ", formulaData)
	data, err := UpdateWipe(formulaData)
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

func DeleteFormula(c *gin.Context) {
  formulaID := c.Param("formulaID")
  fmt.Println("Here is the change id: ", formulaID)
  result := RemoveWipe(formulaID)
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

func GetAllFormulaData() ([]Formula, error) {
    var formulas []Formula
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"FormulaID\", \"Brand\" FROM \"Baby_Tracker\".\"Formula\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var formula Formula
        if err := rows.Scan(&formula.FormulaID, &formula.Brand); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        formulas = append(formulas, formula)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
	fmt.Println(formulas)
    return formulas, nil
}

func GetFormulaData(changeIDstr string) (Formula, error) {
  formulaID, err := strconv.Atoi(changeIDstr)
  var formula Formula
  if err != nil {
    return formula, err
  }
  DBErr := database.DB.QueryRow("SELECT \"FormulaID\", \"Brand\" FROM \"Baby_Tracker\".\"Formula\" WHERE \"FormulaID\" = $1", formulaID).Scan(&formula.FormulaID, &formula.Brand)
  if DBErr != nil {
      return formula, fmt.Errorf("Create Formula: %v", DBErr.Error())
  }
  return formula, nil
}


func CreateWipe(db *sql.DB, formula Formula) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"Formula\" (\"Brand\") VALUES ($1) RETURNING \"FormulaID\"", formula.Brand).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create Formula: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateWipe(formula Formula) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"Formula\" SET \"Brand\" = $1 WHERE \"FormulaID\" = $2 RETURNING \"FormulaID\"", formula.Brand, formula.FormulaID).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update Formula: %v", err)
	}
	return lastInsertId, nil
}

func RemoveWipe(changeIDstr string) int64 {
	formulaID, err := strconv.Atoi(changeIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"Formula\" WHERE \"FormulaID\" = $1", formulaID)
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