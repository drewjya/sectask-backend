package orm

import (
	"github.com/jinzhu/gorm"
)

type Repository interface {
	Create(v interface{}) (err error)
	CreateWithTx(tx *gorm.DB, v interface{}) (err error)
	Save(v interface{}) (err error)
	SaveWithTx(tx *gorm.DB, v interface{}) (err error)
	Delete(v interface{}) (err error)
	DeleteWithTx(tx *gorm.DB, v interface{}) (err error)
	FindOneByID(v interface{}, id uint64) (err error)
	FindAll(v interface{}) (err error)
	FindOneByQuery(v interface{}, params map[string]interface{}, preload string) (err error)
	FindByQuery(v interface{}, params map[string]interface{}, order string) (err error)
	FindAllWithPage(v interface{}, order string, page int, limit int, filters interface{}) (resp *PaginationResponse, err error)
	WithinTransaction(fn DBFunc) (err error)
	PaginateQuery(query, sort, search, newFilter, pagination string) string
	DBManager() *gorm.DB
}
