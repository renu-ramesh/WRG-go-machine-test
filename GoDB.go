/**
 * -------------------------------------------------------------------------------
 * @author Renu P
 * Main package to handle the database related connections
 * -------------------------------------------------------------------------------
 */
package GoDB

import (
	"Settings"
	"database/sql"
	"encoding/hex"
	"errors"

	"fmt"

	_ "mysql"
)

//Database object creation
var (
	DBCon *sql.DB
)


/**
 *  Get Getdbc is used to connect the db
 *  @param none
 *  @return DBobject of type sql.DB
 */
func Getdbc() (*sql.DB, error) {
	DBCon, err := create_connection(Gethost(), Getport(), Getdb(), Getuser(), Getpass())
	DBCon.SetMaxOpenConns(5)
	DBCon.SetMaxIdleConns(3)
	DBCon.SetConnMaxLifetime(time.Second * 1)
	return DBCon, err
}
/**
 * private function create_connection is used to create the connection
 * @param Host string
 * @param Port string
 * @param dbname string
 * @param username string
 * @param password string
 * @return
 *        DBobject of type sql.DB
 *        error
 */
func create_connection(Host string, Port string, dbname string, username string, password string) (*sql.DB, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+Host+")/"+dbname)
	if err != nil {
		return db, err
	}

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		db.Close()
		return db, err
	}

	return db, nil
}
/**
 * Gethost is a function used to get the hostname
 * @param none
 *
 */
func Gethost() string {
	return Settings.DB.Host
}
/**
 * getHost is a function used to get the port
 * @param none
 *
 */
func Getport() string {
	return Settings.DB.Port
}
/**
 * Getpass is a function used to get the hostname
 * @param none
 *
 */
func Getpass() string {
	return Settings.DB.Password
}
/**
 * Getuser is a function used to get the username
 * @param none
 *
 */
func Getuser() string {
	return Settings.DB.Username
}
func Getdb() string {
	return Settings.DB.Database_name
}

func FetchAll(db *sql.DB, tb string, fields string, join string, cond string, offset string, limit string) (rData map[int]map[string]string, errR error) {

	sqlStmt := "SELECT "

	if len(fields) > 0 {
		sqlStmt += fields
	} else {
		sqlStmt += " * "
	}

	sqlStmt += " FROM " + tb + " "

	if len(join) > 0 {
		sqlStmt += join
	}

	if len(cond) > 0 {
		sqlStmt += " WHERE " + cond
	}

	if len(offset) <= 0 {
		offset = "0"
	}

	if len(limit) > 0 {
		sqlStmt += " limit " + offset + "," + limit
	}
	fmt.Println(sqlStmt)
	rows, err := db.Query(sqlStmt)
	
	if err != nil {
		return make(map[int]map[string]string, 0), err
	}

	cols, err := rows.Columns()

	if err != nil {
		return make(map[int]map[string]string, 0), err
	}

	// Result is your slice string.
	rawResult := make([][]byte, len(cols))

	result := make(map[int]map[string]string)
	
	dest := make([]interface{}, len(cols)) // A temporary interface{} slice
	for i := range rawResult {
		dest[i] = &rawResult[i] // Put pointers to each string in the interface slice
	}
	var j int
	for rows.Next() {
		err = rows.Scan(dest...)
		if err != nil {
			return make(map[int]map[string]string, 0), err
		}
		j++
		result[j] = make(map[string]string)
		for i, raw := range rawResult {

			if raw == nil {
				result[j][cols[i]] = ""
			} else {
				result[j][cols[i]] = string(raw)
			}
		}
	}

	return result, nil

}
/**
 * Function FetchField is used to fetch the field value
 * @param DB   *sql.DB
 * @param table string table name
 * @param field String field name
 * @param conditions []string condition string
 * @return field string
 */
func FetchField(db *sql.DB, table string, field string, conditions string) (string, error) {
	sqlStatement := "SELECT " + field + " FROM " + table + " WHERE " + conditions + ";"
	var val string
	// Replace 3 with an ID from your database or another random
	// value to test the no rows use case.
	// fmt.Println("FetchField", sqlStatement)

	fmt.Println(sqlStatement)
	row := db.QueryRow(sqlStatement)

	switch err := row.Scan(&val); err {
	case sql.ErrNoRows:
		return "", errors.New("error: Rows are empty")
	case nil:
		return val, nil
	default:
		return "", err
	}

}
func GoInsert(db *sql.DB, tb string, fields map[string]string) (status bool, errR error) {

	sqlStmt := "INSERT INTO " + tb + " ("
	valueStr := "("
	count := len(fields)
	var i int

	var valArr []interface{}
	for key, val := range fields {
		valArr = append(valArr, val)
		if i++; i < count {
			sqlStmt += key + ","
			valueStr += "?,"
		} else {
			sqlStmt += key
			valueStr += "?"
		}
	}

	sqlStmt += ") "
	valueStr += ")"

	sqlStmt += " VALUES " + valueStr

	stmt, err := db.Prepare(sqlStmt)

	if err != nil {
		return false, err
	}
	defer stmt.Close()
	res, err := stmt.Exec(valArr...)

	if res != nil {
		return true, nil
	} else {
		return false, err
	}

}

func GoRowCount(db *sql.DB, tb string, cond string) (count int, errR error) {
	sqlStmt := "SELECT COUNT(*) as count FROM  " + tb
	if len(cond) > 0 {
		sqlStmt += " where " + cond
	}
	fmt.Println(sqlStmt)
	rows, err := db.Query(sqlStmt)
	if err != nil {
		return 0, err
	}

	for rows.Next() {
		errR = rows.Scan(&count)
		if errR != nil {
			return 0, errR
		}
	}

	if err != nil {
		return 0, err
	}
	return count, nil
}

func GoUpdate(db *sql.DB, tb string, fields map[string]string, cond string, expr string) (status bool, errR error) {
	sqlStmt := "UPDATE " + tb + " "
	sqlStmt += " SET "
	count := len(fields)
	var i int
	var valArr []interface{}
	if count > 0 {
		for key, val := range fields {
			valArr = append(valArr, val)
			if i++; i < count {
				sqlStmt += key + " = ?, "
			} else {
				sqlStmt += key + " = ? "
			}
		}
	}

	if len(expr) > 0 {
		sqlStmt += expr
	}

	if len(cond) > 0 {
		sqlStmt += " WHERE " + cond
	}
	fmt.Println(sqlStmt)
	stmt, err := db.Prepare(sqlStmt)

	if err != nil {
		return false, err
	}
	defer stmt.Close()

	res, err := stmt.Exec(valArr...)

	if res != nil {
		return true, nil
	} else {
		return false, err
	}
}