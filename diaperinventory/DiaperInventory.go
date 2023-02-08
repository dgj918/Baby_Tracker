package diaperinventory

import (
	"database/sql"
	"log"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"gjohnson/Baby_Tracker_Api/Baby_Tracker/database"
	"strconv"
)

type DiaperTransaction struct {
	TransactionID int `json:"TransactionID"`
	DiaperID int `json:"DiaperID"`
	CurrentQty int `json:"CurrentQty"`
	QtyChange int `json:"QtyChange"`
	DateTime string `json:"DateTime"`
}

func GetCurrentQty(c *gin.Context) {
    data, err := RetrieveCurrentQty()
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
  diaperTransactionID := c.Param("diaperTransactionID")
  data, err := GetDiaperData(diaperTransactionID)
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

func PostDiaperInventory(c *gin.Context) {
    var diaperData DiaperTransaction
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
  var diaperData DiaperTransaction
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

func DeleteDiaperInventory(c *gin.Context) {
  diaperTransactionID := c.Param("diaperTransactionID")
  fmt.Println("Here is the change id: ", diaperTransactionID)
  result := RemoveDiaperInventory(diaperTransactionID)
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

/*
### Database Queries and Business Logic
*/

func GetAllTransactions() ([]DiaperTransaction, error) {
    var diaperInventoryArr []DiaperTransaction
	// Define a prepared statement. You'd typically define the statement
    // elsewhere and save it for use in functions such as this one.
    rows, err := database.DB.Query("SELECT \"DateTime\", \"DiaperID\", \"CurrentQty\", \"QtyChange\" FROM \"Baby_Tracker\".\"DiaperTransaction\" ORDER BY \"DateTime\"")
    if err != nil {
		fmt.Println("Error with query")
        log.Fatal(err)
    }
	defer rows.Close()
    // Loop through rows, using Scan to assign column data to struct fields.
    for rows.Next() {
        var diaperInventory DiaperTransaction
        if err := rows.Scan(&diaperInventory.DiaperID, &diaperInventory.CurrentQty, &diaperInventory.QtyChange, &diaperInventory.DateTime); err != nil {
            if err == sql.ErrNoRows {
				// Handle the case of no rows returned.
			}
        }
        diaperInventoryArr = append(diaperInventoryArr, diaperInventory)
    }
    if err := rows.Err(); err != nil {
        return nil, fmt.Errorf("Error with Query: %v", err)
    }
    return diaperInventoryArr, nil
}

func RetrieveCurrentQty() (int, error) {
	var currentQty int
	DBErr := database.DB.QueryRow("SELECT \"CurrentQty\" FROM \"Baby_Tracker\".\"DiaperTransaction\" ORDER BY \"DateTime\" LIMIT 1").Scan(currentQty)
	if DBErr != nil {
		return currentQty, fmt.Errorf("Create DiaperTransaction: %v", DBErr.Error())
	}
	return currentQty, nil
}


func CreateDiaper(db *sql.DB, diaperTransaction DiaperTransaction) (int, error) {
    lastInsertId := 0
    err := db.QueryRow("INSERT INTO \"Baby_Tracker\".\"DiaperInventory\" (\"DiaperID\", \"CurrentQty\", \"QtyChange\", \"DateTime\") VALUES ($1) RETURNING \"DiaperID\"", diaperTransaction.DiaperID, diaperTransaction.CurrentQty, diaperTransaction.QtyChange, diaperTransaction.DateTime).Scan(&lastInsertId)
	if err != nil {
        return 0, fmt.Errorf("Create DiaperTransaction: %v", err.Error())
    }

    return lastInsertId, nil
}

func UpdateDiaperInventory(diaperTransaction DiaperTransaction) (int, error) {
	lastInsertId := 0
	err := database.DB.QueryRow("UPDATE \"Baby_Tracker\".\"DiaperInventory\" SET \"DiaperID\" = $1, \"CurrentQty\" = $2 \"QtyChange\" = $3, \"DateTime\" = $4 WHERE \"TransactionID\" = $5 RETURNING \"TransactionID\"", diaperTransaction.DiaperID, diaperTransaction.CurrentQty, diaperTransaction.QtyChange, diaperTransaction.DateTime).Scan(&lastInsertId)
	fmt.Println(lastInsertId, err)
	if err != nil {
		return 0, fmt.Errorf("Update DiaperTransaction: %v", err)
	}
	return lastInsertId, nil
}

func RemoveDiaperInventory(diaperTransactionIDstr string) int64 {
	diaperTransactionID, err := strconv.Atoi(diaperTransactionIDstr)
	if err != nil {
		return -1
	}
	res, dbErr := database.DB.Exec("DELETE FROM \"Baby_Tracker\".\"DiaperInventory\" WHERE \"TransactionID\" = $1", diaperTransactionID)
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