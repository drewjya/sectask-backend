package orm

/**
*	Custom ORM package
* Based on gorm ORM.
* Make sure the every database actions within the transaction.
* Do begin, rollback and commit the db transaction automatically.
* Return proper error
*
 */

import (
	"fmt"
	"math"
	"reflect"

	// "sectask/entity"
	domain "sectask/domain/orm"

	"github.com/jinzhu/gorm"
)

type (
	PosgreSQL struct {
		db *gorm.DB
	}
)

func New(db *gorm.DB) *PosgreSQL {
	return &PosgreSQL{db}
}

var total_rec int

// Create
// Helper function to insert gorm model to database by using 'WithinTransaction'
func (psql *PosgreSQL) Create(v interface{}) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		// check new object
		if !psql.db.NewRecord(v) {
			return err
		}

		if err = tx.Create(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// CreateWithTx
// Helper function to insert gorm model to database
func (psql *PosgreSQL) CreateWithTx(tx *gorm.DB, v interface{}) (err error) {
	// check new object
	if !psql.db.NewRecord(v) {
		return err
	}

	if err = tx.Create(v).Error; err != nil {
		return err
	}

	return err
}

// Save
// Helper function to save gorm model to database by using 'WithinTransaction'
func (psql *PosgreSQL) Save(v interface{}) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		// check new object
		if psql.db.NewRecord(v) {
			return err
		}

		if err = tx.Debug().Save(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// SaveWithTx
// Helper function to save gorm model to database
func (psql *PosgreSQL) SaveWithTx(tx *gorm.DB, v interface{}) (err error) {
	// check new object
	if psql.db.NewRecord(v) {
		return err
	}

	if err = tx.Save(v).Error; err != nil {
		return err
	}

	return err
}

// Delete
// Helper function to save gorm model to database by using 'WithinTransaction'
func (psql *PosgreSQL) Delete(v interface{}) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		// check new object
		if err = tx.Delete(v).Error; err != nil {
			return err
		}

		return err
	})
}

// DeleteWithTx
// Helper function to save gorm model to database
func (psql *PosgreSQL) DeleteWithTx(tx *gorm.DB, v interface{}) (err error) {
	// check new object
	if err = tx.Delete(v).Error; err != nil {
		tx.Rollback() // rollback
		return err
	}

	return err
}

// FindOneByID
// Helper function to find a record by using 'WithinTransaction'
func (psql *PosgreSQL) FindOneByID(v interface{}, id uint64) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Last(v, id).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// FindAll
// Helper function to find records by using 'WithinTransaction'
func (psql *PosgreSQL) FindAll(v interface{}) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Find(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// FindOneByQuery
// Helper function to find a record by using 'WithinTransaction'
func (psql *PosgreSQL) FindOneByQuery(v interface{}, params map[string]interface{}, preload string) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		if preload != "" {
			err = tx.Preload(preload).Where(params).Last(v).Error
		} else {
			err = tx.Where(params).Last(v).Error
		}

		if err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// FindByQuery
// Helper function to find records by using 'WithinTransaction'
func (psql *PosgreSQL) FindByQuery(v interface{}, params map[string]interface{}, order string) (err error) {
	return psql.WithinTransaction(func(tx *gorm.DB) error {
		if err = tx.Where(params).Order(order).Find(v).Error; err != nil {
			tx.Rollback() // rollback
			return err
		}

		return err
	})
}

// FindAllWithPage
// Helper function to find all records in pagination by using 'WithinTransaction'
// v interface{} Gorm model struct
// page int Page number
// rp int Record per page to be showed
// filters int Gorm model struct for filters
func (psql *PosgreSQL) FindAllWithPage(v interface{}, order string, page int, limit int, filters interface{}) (resp *domain.PaginationResponse, err error) {
	if limit <= 0 {
		limit = 25
	}

	// tx := psql.db.Begin()
	tx := psql.db

	// loop through "filters"
	refOf := reflect.ValueOf(filters).Elem()
	typeOf := refOf.Type()
	for i := 0; i < refOf.NumField(); i++ {
		f := refOf.Field(i)

		// ignore if empty
		// just make sure ModelFilterable its all in string type
		if f.Interface() != "" {
			tx = tx.Where(fmt.Sprintf("%s = ?", typeOf.Field(i).Tag.Get("json")), f.Interface())
		}
	}

	// copy of tx
	ctx := tx

	// get total record include filters
	ctx.Find(v).Count(&total_rec)

	offset := (page * limit) - limit

	lastPage := int(math.Ceil(float64(total_rec) / float64(limit)))

	tx.Order(order).Limit(limit).Offset(offset).Find(v)

	resp = &domain.PaginationResponse{
		Total:       total_rec,
		PerPage:     limit,
		CurrentPage: page,
		LastPage:    lastPage,
		From:        offset + 1,
		To:          offset + limit,
		Data:        &v,
	}

	if err != nil {
		// tx.Rollback() // rollback db transaction
		return resp, err
	}

	// tx.Commit()

	return resp, err
}

// WithinTransaction
// accept DBFunc as parameter
// call DBFunc function within transaction begin, and commit and return error DBFunc
func (psql *PosgreSQL) WithinTransaction(fn domain.DBFunc) (err error) {
	tx := psql.db.Begin() // start db transaction
	defer tx.Commit()
	err = fn(tx)
	// close db transaction
	return err
}

// Pagination with raw query
func (psql *PosgreSQL) PaginateQuery(query, sort, search, newFilter, pagination string) string {
	return fmt.Sprintf(`
		SELECT *
		FROM (
			SELECT ROW_NUMBER () OVER (ORDER BY %[2]s) AS row_number, *
			FROM (%[1]s) rs
			WHERE (%[3]s) AND (%[4]s)
			ORDER BY %[2]s
		) applicants_result
		WHERE %[5]s
	`, query, sort, search, newFilter, pagination)
}

// Gorm: return GORM's postgres database connection instance
func (psql *PosgreSQL) DBManager() *gorm.DB {
	return psql.db
}
