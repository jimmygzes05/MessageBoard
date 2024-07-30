package mydb

import (
	"database/sql"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	// "gorm.io/driver/mysql"
	// "gorm.io/gorm"
)

func TestGetAllComment(t *testing.T) {
	var db *sql.DB
	var err error
	var mock sqlmock.Sqlmock

	db, mock, err = sqlmock.New() // mock sql.DB
	if err != nil {
		t.Errorf("Failed to open mock sql db, got error: %v", err)
	}
	defer db.Close()

	t.Run("Get_All_Comment_Success", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectExec(
			regexp.QuoteMeta(
				"SELECT * FROM `comment`",
			),
		)
	})

}
