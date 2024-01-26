package stock

import "sectask/libs/table"

func (s *StockConfig) GetStockByName(name string) (StockDBResponse, error) {
	var result StockDBResponse
	err := s.Db.Table(table.TableStocks).Select("*").Where("name = ? and deleted_at IS NULL", name).Find(&result).Error
	return result, err
}

func (s *StockConfig) GetStockByMultipleName(names []string) ([]StockDBResponse, error) {
	var result []StockDBResponse
	err := s.Db.Table(table.TableStocks).Select("*").Where("name IN ? and deleted_at IS NULL", names).Find(&result).Error
	return result, err
}
