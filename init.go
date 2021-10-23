/**
 * -------------------------------------------------------------------------------
 * @author Renu P
 *  Initialisation of all the process
 * List_Article,Create_Article,Approve_Decline_Article
 * -------------------------------------------------------------------------------
 */
 package Init
 import(
	"GoDB"
	"time"
	U "Utility"
 )
/*------------------------------------
 TO LIST ARTICLES
--------------------------------------*/
func List_Article(processCron *sync.WaitGroup) {
	pr := U.PR
	db, err := GoDB.Getdbc()
	defer db.Close()
	if err != nil {
		processCron.Done()
		return
	}
	
	//Get articles data from database
	var articleData map[int]map[string]string
	cond := "status = 1"
	articleData, err = GoDB.FetchAll(db, "articles", "", "", cond, "", "") // can add limit if the article table is bulk in size
	if err != nil {
		processCron.Done()
		return
	}
	//List Approved articles
	if len(articleData) > 0 {
		pr("================ ARTICLES =============")
		for id, value := range articleData {
			pr(id+": "+ value["tittle"])
			pr("------------------------------")
			pr(value[description])
			pr("------------------------------")
			pr(" Written By "+ value["name"]+" on "+value[created])
		}
	}
	
	//For garbage collection
	time.Sleep(10 * time.Millisecond)
	processCron.Done()
}
/*------------------------------------
TO CREATE NEW ARTICLE (Signin/Signup required)
--------------------------------------*/
func Create_Article(processCron *sync.WaitGroup) {
	startTime := time.Now()
	db, err := GoDB.Getdbc()
	defer db.Close()
	if err != nil {
		processCron.Done()
		return
	}
	//user signup/signin
	var uid int
	var login bool
	login,uid,err = U.UserLogin(db)
	if !login{
		fmt.Println(" User login failed ...! If you want to Signup? Y/N :"
		var choice string
		c, _ := fmt.Scanln(&choice)
		switch c {
			case 'Y'||'y':
				uid = U.UserSignup(db)//user Signup
			case 'N'||'n':
				processCron.Done()//exit
				return
			default:
				fmt.Println("Invalid Choice...!!")
				processCron.Done()
				return
		}
	}
	add_new_article(db,uid)
	
	//For garbage collection
	time.Sleep(10 * time.Millisecond)
	processCron.Done()
}
/**
 *  add_new_article is used to insert new article data into database tables
 *  @param dbconnection object,user id
 *  @return none
 */
func add_new_article(db *sql.DB, uid int){
	pr := U.PR
	created := time.Now()
	
	var tittle,description,created,status string
	pr(" Enter Tittle of the Article  : ")
	t, _ := fmt.Scanln(&tittle)
	pr(" Enter the content : ")
	desc,_ := fmt.Scanln(&description)
	
	username, _ := GoDB.FetchField(db, "users", "name", "uid = "+uid)
	
	Fields := map[string]string{}
	Fields["uid"] = uid
	Fields["name"] = username
	Fields["tittle"] = t
	Fields["description"] = desc
	Fields["created"] = created
	Fields["status"] = 0 // Admin/autherised user need to approve
	GoDB.GoInsert(db, "article", Fields)
	

}
/*------------------------------------
TO APPROVE NEWLY CREATED ARTICLE (Signin & authentication are required)
--------------------------------------*/
func Approve_Decline_Article(processCron *sync.WaitGroup) {
	pr := U.PR
	startTime := time.Now()
	db, err := GoDB.Getdbc()
	defer db.Close()
	if err != nil {
		processCron.Done()
		return
	}
	pr(" This functionality required authentication...")
	login,uid,err = U.UserLogin(db)
	
	// Check autherisation
	if !login && uid==0 && U.AuthenticatedUser(db,uid) {
		pr("You are not Autherised to access this functionality")
		processCron.Done()
		return
	}
	pr("============= ARTICLES WAITING FOR APPROVAL =============")
	var articleData map[int]map[string]string
	cond := "status = 0"
	articleData, err = GoDB.FetchAll(db, "articles", "", "", cond, "", "")
	if err != nil {
		processCron.Done()
		return
	}
	//List all unapproved articles
	if len(articleData) > 0 {
		fmt.printf(" Article Id\t Tittle\t Auther\t Created\n")
		pr("--------------------------------------------------)
		for id, value := range articleData {
			fmt.printf("%d\t%s\t%s\t%v\n",value["art_id",value["tittle"],value["name"],value["created"])
		}
		pr(" Please enter article number from above list to approve : ")
		var choice int
		article_id, _ := fmt.Scanln(&choice)
		err = U.approve_article(db,article_id)
		if err != nil {
			processCron.Done()
			return
		}
	}
	
	
	//For garbage collection
	time.Sleep(10 * time.Millisecond)
	processCron.Done()
}
/**
 *  approve_article is used to update the status of the article form declined to approve
 *  @param dbconnection object,article id
 *  @return error
 */
func approve_article(db *sql.DB, id int) err error{

	article_name, err := GoDB.FetchField(db, "article", "tittle", "art_id = "+id)
	if article_name !=""{
		upFiedls := map[string]string{
			"status": "1",
		}
		_, err = GoDB.GoUpdate(db, "article", upFiedls, "art_id = "+id, "")
		if err==nil{
			fmt.Println(" Article - "+ article_name +" Approved!")
			return nil
		}
	}
	return err
}