/**
 * -------------------------------------------------------------------------------
 * @author Renu P
 * Settings for the Database Connection
 * Users can change host,port,dbname,user and password
 * -------------------------------------------------------------------------------
 */
package Settings

// DB connection structure
// Store DB settings
type DBConnection struct {
	Host, Port, Database_name, Username, Password string
}
var (

	/**
	 * Database configuration comes here
	 * Host : Host for the database
	 * Port : TCP port
	 * Database_name : Database name
	 * Username : username
	 * Password : database password
	 */

	/*---------- Local Settings ------------*/
	DB = DBConnection{
		Host:          "localhost",
		Port:          "2222",
		Database_name: "WRG_GO_TEST",
		Username:      "root",
		Password:      "root",
	}

)
