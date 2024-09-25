package mysql_heap

import (
	"container/heap"
	"database/sql"
	"fmt"
)

type RowDataIterator interface {
	// Next seeks the next row data, it used when compared rows.
	Next() (map[string]*ColumnData, error)
	// Close release the resource.
	Close()
}

type TableSource struct {
	OriginSchema string
	OriginTable  string
}

type TableShardSource struct {
	Table  string
	DBConn *sql.Rows
}

type MySQLSource struct {
	sourceTables map[string][]TableShardSource
}

type Table struct {
	Name         string
	OrderKeyCols []Column
}

func (s *MySQLSource) GetRowsIterator(table Table) (RowDataIterator, error) {
	source, ok := s.sourceTables[table.Name]
	if !ok {
		return nil, fmt.Errorf("table %s not found", table.Name)
	}
	sourceRows := make(map[int]*sql.Rows)
	for i, ms := range source {
		sourceRows[i] = ms.DBConn
	}

	sourceRowDatas := &RowDatas{
		Rows:         make([]RowData, 0, len(sourceRows)),
		OrderKeyCols: table.OrderKeyCols,
	}
	heap.Init(sourceRowDatas)
	for src, sRows := range sourceRows {
		rowData, err := getRowData(sRows)
		if err != nil {
			return nil, err
		}

		if rowData != nil {
			heap.Push(sourceRowDatas, RowData{
				Data:   rowData,
				Source: src,
			})
		} else {
			if sRows.Err() != nil {
				return nil, sRows.Err()
			}
		}
	}
	return &MultiSourceRowsIterator{
		sourceRows:     sourceRows,
		sourceRowDatas: sourceRowDatas,
	}, nil
}

type MultiSourceRowsIterator struct {
	sourceRows     map[int]*sql.Rows
	sourceRowDatas *RowDatas
}

func (ms *MultiSourceRowsIterator) Next() (map[string]*ColumnData, error) {
	if len(ms.sourceRowDatas.Rows) == 0 {
		return map[string]*ColumnData{}, nil
	}
	rowData := heap.Pop(ms.sourceRowDatas).(RowData)
	newRowData, err := getRowData(ms.sourceRows[rowData.Source])
	if err != nil {
		return nil, err
	}
	if newRowData != nil {
		heap.Push(ms.sourceRowDatas, RowData{
			Data:   newRowData,
			Source: rowData.Source,
		})
	} else {
		if ms.sourceRows[rowData.Source].Err() != nil {
			return nil, ms.sourceRows[rowData.Source].Err()
		}
	}
	return rowData.Data, nil
}

func (ms *MultiSourceRowsIterator) Close() {
	for _, s := range ms.sourceRows {
		s.Close()
	}
	return
}

func getRowData(rows *sql.Rows) (rowData map[string]*ColumnData, err error) {
	for rows.Next() {
		rowData, err = ScanRow(rows)
		return
	}
	return
}

func ScanRow(rows *sql.Rows) (map[string]*ColumnData, error) {
	cols, err := rows.Columns()
	if err != nil {
		return nil, err
	}

	colVals := make([][]byte, len(cols))
	colValsI := make([]interface{}, len(colVals))
	for i := range colValsI {
		colValsI[i] = &colVals[i]
	}

	err = rows.Scan(colValsI...)
	if err != nil {
		return nil, err
	}

	result := make(map[string]*ColumnData)
	for i := range colVals {
		data := &ColumnData{
			Data:   colVals[i],
			IsNull: colVals[i] == nil,
		}
		result[cols[i]] = data
	}

	return result, nil
}
