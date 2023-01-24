package models

type DataTable struct {
	Table     map[string]string
	Equations []string
	Columns   []string
	Rows      []string
}

func NewDataTable() *DataTable {
	table := make(map[string]string, 0)
	equations := []string{}
	columns := []string{}
	rows := []string{}
	return &DataTable{
		Table:     table,
		Equations: equations,
		Columns:   columns,
		Rows:      rows,
	}
}
