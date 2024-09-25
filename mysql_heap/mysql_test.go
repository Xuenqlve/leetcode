package mysql_heap

import (
	"context"
	"database/sql"
	"database/sql/driver"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/bxcodec/faker/v4"
	"testing"
)

func MockSQLRow(ctx context.Context, column []string, mockData [][]driver.Value) (rows *sql.Rows, err error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		return
	}
	expectedRows := sqlmock.NewRows(column)
	for _, row := range mockData {
		expectedRows.AddRow(row...)
	}
	mock.ExpectQuery("select*").WillReturnRows(expectedRows)
	return db.QueryContext(ctx, "select")
}

func MakeMockData(sum, count int) (column []string, mockDatas [][][]driver.Value) {
	mockColumn := []string{"id", "name", "email", "create_time"}
	mockDatas = make([][][]driver.Value, count)
	for i := 0; i < count; i++ {
		mockDatas[i] = make([][]driver.Value, 0, sum)
	}
	mockData := make([][]driver.Value, 0, sum)
	for i := 1; i < sum; i++ {
		index := i % count
		mockDatas[index] = append(mockDatas[index], []driver.Value{
			i, faker.Name(), faker.Email(), faker.Timestamp(),
		})
		mockData = append(mockData, []driver.Value{
			i, faker.Name(), faker.Email(), faker.Timestamp(),
		})
	}
	return mockColumn, mockDatas
}

func TestMySQLHeap(t *testing.T) {
	column, mockdata := MakeMockData(50, 4)
	orderKeyCols := []Column{
		{
			Name: "id",
			Type: TypeInt24,
		},
	}

	s := MySQLSource{
		sourceTables: map[string][]TableShardSource{},
	}
	tableName := "test"
	s.sourceTables[tableName] = make([]TableShardSource, 0, len(mockdata))
	for index, data := range mockdata {
		sqlRows, err := MockSQLRow(context.Background(), column, data)
		if err != nil {
			t.Errorf("make mock sql row err:%v", err)
			return
		}
		s.sourceTables[tableName] = append(s.sourceTables[tableName], TableShardSource{
			Table:  fmt.Sprintf("%s_%d", tableName, index),
			DBConn: sqlRows,
		})
	}
	table := Table{
		Name:         tableName,
		OrderKeyCols: orderKeyCols,
	}

	rows, err := s.GetRowsIterator(table)
	if err != nil {
		t.Errorf("get rows iterator err:%v", err)
		return
	}
	index := 1
	t.Log("start for range ")
	for {
		data := map[string]*ColumnData{}
		if data, err = rows.Next(); err != nil {
			t.Errorf("rows next err:%v", err)
			return
		}

		if len(data) == 0 {
			break
		}

		for k, v := range data {
			if k == "id" {
				t.Logf("index:%d column:%v value:%v", index, k, string(v.Data))
			}
		}
		index++
	}
	t.Log("success")
}
