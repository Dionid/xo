package models

// Code generated by xo. DO NOT EDIT.

import (
	"context"
)

// ForeignKey is a foreign key.
type ForeignKey struct {
	ForeignKeyName string `json:"foreign_key_name"` // foreign_key_name
	ColumnName     string `json:"column_name"`      // column_name
	RefTableName   string `json:"ref_table_name"`   // ref_table_name
	RefColumnName  string `json:"ref_column_name"`  // ref_column_name
	KeyID          int    `json:"key_id"`           // key_id
}

// PostgresTableForeignKeys runs a custom query, returning results as [ForeignKey].
func PostgresTableForeignKeys(ctx context.Context, db DB, schema, table string) ([]*ForeignKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`tc.constraint_name, ` + // ::varchar AS foreign_key_name
		`kcu.column_name, ` + // ::varchar AS column_name
		`ccu.table_name, ` + // ::varchar AS ref_table_name
		`ccu.column_name, ` + // ::varchar AS ref_column_name
		`0 ` + // ::integer AS key_id
		`FROM information_schema.table_constraints tc ` +
		`JOIN information_schema.key_column_usage AS kcu ON tc.constraint_name = kcu.constraint_name ` +
		`AND tc.table_schema = kcu.table_schema ` +
		`AND tc.table_name = kcu.table_name ` +
		`JOIN ( ` +
		`SELECT ` +
		`ROW_NUMBER() OVER ( ` +
		`PARTITION BY ` +
		`table_schema, ` +
		`table_name, ` +
		`constraint_name ` +
		`ORDER BY row_num ` +
		`) AS ordinal_position, ` +
		`table_schema, ` +
		`table_name, ` +
		`column_name, ` +
		`constraint_name ` +
		`FROM ( ` +
		`SELECT ` +
		`ROW_NUMBER() OVER (ORDER BY 1) AS row_num, ` +
		`table_schema, ` +
		`table_name, ` +
		`column_name, ` +
		`constraint_name ` +
		`FROM information_schema.constraint_column_usage ` +
		`) t ` +
		`) AS ccu ON ccu.constraint_name = tc.constraint_name ` +
		`AND ccu.table_schema = tc.table_schema ` +
		`AND ccu.ordinal_position = kcu.ordinal_position ` +
		`WHERE tc.constraint_type = 'FOREIGN KEY' ` +
		`AND tc.table_schema = $1 ` +
		`AND tc.table_name = $2`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ForeignKey
	for rows.Next() {
		var fk ForeignKey
		// scan
		if err := rows.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefTableName, &fk.RefColumnName, &fk.KeyID); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &fk)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// MysqlTableForeignKeys runs a custom query, returning results as [ForeignKey].
func MysqlTableForeignKeys(ctx context.Context, db DB, schema, table string) ([]*ForeignKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`constraint_name AS foreign_key_name, ` +
		`column_name AS column_name, ` +
		`referenced_table_name AS ref_table_name, ` +
		`referenced_column_name AS ref_column_name ` +
		`FROM information_schema.key_column_usage ` +
		`WHERE referenced_table_name IS NOT NULL ` +
		`AND table_schema = ? ` +
		`AND table_name = ?`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ForeignKey
	for rows.Next() {
		var fk ForeignKey
		// scan
		if err := rows.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefTableName, &fk.RefColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &fk)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// Sqlite3TableForeignKeys runs a custom query, returning results as [ForeignKey].
func Sqlite3TableForeignKeys(ctx context.Context, db DB, schema, table string) ([]*ForeignKey, error) {
	// query
	sqlstr := `/* ` + schema + ` */ ` +
		`SELECT ` +
		`id AS key_id, ` +
		`"table" AS ref_table_name, ` +
		`"from" AS column_name, ` +
		`"to" AS ref_column_name ` +
		`FROM pragma_foreign_key_list($1)`
	// run
	logf(sqlstr, table)
	rows, err := db.QueryContext(ctx, sqlstr, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ForeignKey
	for rows.Next() {
		var fk ForeignKey
		// scan
		if err := rows.Scan(&fk.KeyID, &fk.RefTableName, &fk.ColumnName, &fk.RefColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &fk)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// SqlserverTableForeignKeys runs a custom query, returning results as [ForeignKey].
func SqlserverTableForeignKeys(ctx context.Context, db DB, schema, table string) ([]*ForeignKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`fk.name AS foreign_key_name, ` +
		`col.name AS column_name, ` +
		`pk_tab.name AS ref_table_name, ` +
		`pk_col.name AS ref_column_name ` +
		`FROM sys.tables tab ` +
		`INNER JOIN sys.columns col ON col.object_id = tab.object_id ` +
		`LEFT OUTER JOIN sys.foreign_key_columns fk_cols ON fk_cols.parent_object_id = tab.object_id ` +
		`AND fk_cols.parent_column_id = col.column_id ` +
		`LEFT OUTER JOIN sys.foreign_keys fk ON fk.object_id = fk_cols.constraint_object_id ` +
		`LEFT OUTER JOIN sys.tables pk_tab ON pk_tab.object_id = fk_cols.referenced_object_id ` +
		`LEFT OUTER JOIN sys.columns pk_col ON pk_col.column_id = fk_cols.referenced_column_id ` +
		`AND pk_col.object_id = fk_cols.referenced_object_id ` +
		`WHERE schema_name(tab.schema_id) = @p1 ` +
		`AND tab.name = @p2 ` +
		`AND fk.object_id IS NOT NULL`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ForeignKey
	for rows.Next() {
		var fk ForeignKey
		// scan
		if err := rows.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefTableName, &fk.RefColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &fk)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}

// OracleTableForeignKeys runs a custom query, returning results as [ForeignKey].
func OracleTableForeignKeys(ctx context.Context, db DB, schema, table string) ([]*ForeignKey, error) {
	// query
	const sqlstr = `SELECT ` +
		`LOWER(a.constraint_name) AS foreign_key_name, ` +
		`LOWER(a.column_name) AS column_name, ` +
		`LOWER(c_pk.table_name) AS ref_table_name, ` +
		`LOWER(b.column_name) AS ref_column_name ` +
		`FROM user_cons_columns a ` +
		`JOIN user_constraints c ON a.owner = c.owner ` +
		`AND a.constraint_name = c.constraint_name ` +
		`JOIN user_constraints c_pk ON c.r_owner = c_pk.owner ` +
		`AND c.r_constraint_name = c_pk.constraint_name ` +
		`JOIN user_cons_columns b ON C_PK.owner = b.owner ` +
		`AND C_PK.CONSTRAINT_NAME = b.constraint_name AND b.POSITION = a.POSITION ` +
		`WHERE c.constraint_type = 'R' ` +
		`AND a.owner = UPPER(:1) ` +
		`AND a.table_name = UPPER(:2)`
	// run
	logf(sqlstr, schema, table)
	rows, err := db.QueryContext(ctx, sqlstr, schema, table)
	if err != nil {
		return nil, logerror(err)
	}
	defer rows.Close()
	// load results
	var res []*ForeignKey
	for rows.Next() {
		var fk ForeignKey
		// scan
		if err := rows.Scan(&fk.ForeignKeyName, &fk.ColumnName, &fk.RefTableName, &fk.RefColumnName); err != nil {
			return nil, logerror(err)
		}
		res = append(res, &fk)
	}
	if err := rows.Err(); err != nil {
		return nil, logerror(err)
	}
	return res, nil
}
