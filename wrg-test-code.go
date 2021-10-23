/**
 * -------------------------------------------------------------------------------
 * @author Renu P
 * Main package to handle the User choices
 * This package act as the GUI for user selection
 * -------------------------------------------------------------------------------
 */
 package main
 
 import(
	 "fmt"
	 U "Utility"
	 "gocron"
	 "sync"
	 Init "/init"
 )
 /*------------------------------------
 Main function
 --------------------------------------*/
 func main() {
	pr := U.PR
	pr("Welcome...! Please Enter your choice ?")
	pr("1: List Articles")
	pr("2: Create Articles")
	pr("3: Approve/Decline Articles")
	pr("4: Exit")
	
	//go channel process
	go startCron()
	ch := make(chan string)
	<-ch
			
	 
 }
 /*------------------------------------
 concurrent channel process to start on every 2seconds
 --------------------------------------*/
 func startCron() {
	gocron.Every(2).Seconds().Do(addCronJobs)
	<-gocron.Start()
}
 /*------------------------------------
 Process begins according to the user selection
 --------------------------------------*/
func addCronJobs() {

	var processCron sync.WaitGroup

	var input int
	for ok := true; ok; ok = (input != 4) {
		n, err := fmt.Scanln(&input)
		if n < 1 || err != nil {
			panic(err)
			return
		}
		else{
			switch input {
				case 1:
					go Init.List_Article(&processCron) //List Approved Articles
				case 2:
					go Init.Create_Article(&processCron) //Create New Articles
				case 3:
					go Init.Approve_Decline_Article(&processCron) // Approve articles by autherised users only
				default:
					fmt.Println("Invalid Choice...!!")
			}

		}
	}
	processCron.Wait()

}














