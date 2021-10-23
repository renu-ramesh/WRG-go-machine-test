/**
 * -------------------------------------------------------------------------------
 * @author Renu P
 * Import the common package for the application avoid the duplication of the package
 *  User Authentication process
 * -------------------------------------------------------------------------------
 */

package Utility
import (
 "runtime"
 "fmt"
 "ErrorHandler/EpsLogs"
 "database/sql"
 "crypto/md5"
 "encoding/hex"
)

var (
	PR = fmt.Println
	// PR = TestK
	PS = fmt.Sprintf
)

func GetMD5Hash(text string) string {
   hash := md5.Sum([]byte(text))
   return hex.EncodeToString(hash[:])
}


/**
 *  approve_article is used to update the status of the article form declined to approve
 *  @param dbconnection object,article id
 *  @return error
 */

func UserLogin(db *sql.DB)(status bool,uid int,err error){
	pr := U.PR
	var uname,password string
	pr(" Please Enter Your username : ")
	name, _ := fmt.Scanln(&uname)
	pr(" Please Enter Your Password : ")
	pass,_ := fmt.Scanln(&password)
	
	cond := "name = " + name + " AND password = "+ GetMD5Hash(pass) 
	uid, _ := GoDB.FetchField(db, "users", "id", cond)
	if len(uid)>0
		return true,uid,nil
	return false,0,nil
}
func UserSignup(db *sql.DB) (uid int){
	pr := U.PR
	uid = 0
	var uname,password,email string
	pr(" Please Enter Your username : ")
	name, _ := fmt.Scanln(&uname)
	pr(" Please Enter Your Password : ")
	pass,_ := fmt.Scanln(&password)
	pr(" Please Enter Your Email : ")
	Email,_ := fmt.Scanln(&email)
	
	Fields := map[string]string{}
	Fields["password"] = GetMD5Hash(pass)
	Fields["name"] = username
	Fields["email"] = Email
	Fields["is_admin"] = "0"
	GoDB.GoInsert(db, "users", Fields)
	cond := "name = " + name + " AND password = "+ pass 
	uid, _ = GoDB.FetchField(db, "users", "id", cond)
	
	return uid
}
func AuthenticatedUser(db *sql.DB,uid int)status bool{
	cond := "id = " + uid + " AND is_admin = 1"
	exist, _ := GoDB.GoRowCount(db, "users", cond)
	if exist>0
		return true
	else
		return false
}