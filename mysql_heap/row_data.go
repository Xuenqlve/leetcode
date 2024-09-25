package mysql_heap

import (
	"go.uber.org/zap"
	"log"
	"strconv"
)

type ColumnData struct {
	Data   []byte
	IsNull bool
}

type RowData struct {
	Data   map[string]*ColumnData
	Source int
}

type Column struct {
	Name string `json:"name"`
	Type byte   `json:"type"`
}

type RowDatas struct {
	Rows         []RowData
	OrderKeyCols []Column
}

func (r *RowDatas) Len() int {
	return len(r.Rows)
}

func (r *RowDatas) Swap(i, j int) {
	r.Rows[i], r.Rows[j] = r.Rows[j], r.Rows[i]
}

func (r *RowDatas) Less(i, j int) bool {
	for _, col := range r.OrderKeyCols {
		col1, ok := r.Rows[i].Data[col.Name]
		if !ok {
			log.Fatal("data don't have column", zap.String("column", col.Name), zap.Reflect("data", r.Rows[i].Data))
		}
		col2, ok := r.Rows[j].Data[col.Name]
		if !ok {
			log.Fatal("data don't have column", zap.String("column", col.Name), zap.Reflect("data", r.Rows[j].Data))
		}
		switch {
		case col1.IsNull && col2.IsNull:
			continue
		case col1.IsNull:
			return true
		case col2.IsNull:
			return false
		}

		strData1 := string(col1.Data)
		strData2 := string(col2.Data)

		if NeedQuotes(col.Type) {
			if strData1 == strData2 {
				continue
			}
			return strData1 < strData2
		}

		num1, err1 := strconv.ParseFloat(strData1, 64)
		if err1 != nil {
			log.Fatal("convert string to float failed", zap.String("column", col.Name), zap.String("data", strData1), zap.Error(err1))
		}
		num2, err2 := strconv.ParseFloat(strData2, 64)
		if err2 != nil {
			log.Fatal("convert string to float failed", zap.String("column", col.Name), zap.String("data", strData2), zap.Error(err2))
		}
		if num1 == num2 {
			continue
		}
		return num1 < num2
	}
	return false
}

func (r *RowDatas) Push(x any) {
	r.Rows = append(r.Rows, x.(RowData))
}

func (r *RowDatas) Pop() (x any) {
	if len(r.Rows) == 0 {
		return nil
	}
	r.Rows, x = r.Rows[:len(r.Rows)-1], r.Rows[len(r.Rows)-1]
	return
}
