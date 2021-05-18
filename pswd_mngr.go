package main
import (
	"golang.org/x/crypto/bcrypt"
	"bufio"
	"fmt"
	"os"
	"strings"	
)

const db = "pwdDB.db"

func main(){
	var args []string = os.Args

	if(args[1]=="put"){     //if CLA is "put" then we store the password in database
		save(args[2],args[3],args[4])
	}else if(args[1]=="get"){   //if CLA is "get" then we retrive the username and password from the database
		retrive(args[2],args[3])
	}else if(args[1]=="compare"){                  //if CLA is "compare" then we compare user input password 
		compare(args[2],args[3],args[4])            //and password stored in the database
	}else{
		fmt.Println("Platform is Not Applied ",args[2])
	}
}

func save(platform string, username string, password string){

	//converts string into byte array because bcrypt only accept byte array
	b:=[]byte(password)

	//encypting password into hash
	hash, err := bcrypt.GenerateFromPassword(b, bcrypt.MinCost) 
	if(err!=nil){
		fmt.Println(err)
		return
	}

	//stores platform and username in data variable
	data:=platform + "," + username + ","   

	//open the file in WRITE mode, if file is not present then it will create
	f,err := os.OpenFile(db,os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if(err!=nil){
		fmt.Println(err)
		return
	}

	//adding platform and username
	//l1 is size of the string passed
	l1,err := f.WriteString(data)
	if(err!=nil){
		fmt.Println(err)
		return
	}

	//adding password
	l2,err := f.WriteString(string(hash)) //string(hash) function save the hash int the string format
	if(err!=nil){
		fmt.Println(err)
		return
	}

	//adding new line
	l3,err := f.WriteString("\n")
	if(err!=nil){
		fmt.Println(err)
		return
	}

	//check the data was stored or not
	if(l1!=0 && l2!=0 && l3!=0){
		fmt.Print("Credentials Saved")
	}

	//close the file
	err =f.Close()
	if(err!=nil){
		fmt.Println(err)
		return
	}

}

func retrive(platform string, username string){

	//open file to retrive the data
	f,err := os.Open(db)
	if(err!=nil){
		fmt.Print(err)
		return
	}

	//sacan the line using bufio package which does buffered IO operations.
	input:=bufio.NewScanner(f)
	for input.Scan() {
		//split it using strings package’s Split method
		data := strings.Split(input.Text(), ",")
		if(data[0]==platform){
			if(data[1]==username){
				fmt.Println(data[1],": ", data[2])
				return
			}
		}
	}
	fmt.Print("Platform %s not known \n", platform)
}


func compare(platform string, username string, password string){

	//open the data base to compare the password 
	f,err := os.Open(db)
	if(err!=nil){
		fmt.Println(err)
		return
	}

	input:=bufio.NewScanner(f)
	for input.Scan() {
		data := strings.Split(input.Text(), ",")
		if(data[0]==platform){
			if(data[1]==username){

				//compare the password using CompareHashAndPassword method //RETURN NIL IF PASSWORD MATCHED
				err := bcrypt.CompareHashAndPassword([]byte(data[2]), []byte(password)); 
				if(err!=nil){
					fmt.Print(err,">>>Password Not Matched") ///print an error
					return
				}else{
					fmt.Print("Matched")
					return
				}
			}
		}
	}
	fmt.Println("Password Not Matched : ", username)
}


//Execution of program

//Before compilation make sure that you have installed bcrypt source file 

//Run the command : go build file_name.go

//Pass the argument :
//     filename operation platrform username password

//e.g. pswd_mngr compare facebook aniket43 hi@git