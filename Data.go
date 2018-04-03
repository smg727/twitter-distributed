package main

import (
	"fmt"

	"net/http"
)


var userdata = make(map[string]User)

type User struct {
	username string
	password string
	tweets []tweet
	follows map[string]bool
}

type tweet struct {
	text string
}

var debugon = true //if set to true debug outputs are printed

//function to print debug outputs if debugon=true
func debugPrint(text string){
	if(debugon){
		fmt.Println(text)
	}
}

func printer(){ //test function to call one function from another goclass. Remove at end
	debugPrint("Debug: function test")
}

//function to add user to data on registration
func addUser(usrname string, pwd string) int  {
	_, ok := userdata[usrname]
	if(ok){
		debugPrint("Debug: User already exists")
		return 0
	}
	usr := User{username:usrname,password:pwd}
	userdata[usrname] = usr
	debugPrint("Debug: User added")
	return 1
}

func getPassword(usrname string) (bool, string){ // returns users password

	user, ok := userdata[usrname]
	if(!ok){
		debugPrint("No such user")
		return false, "No such User"
	}
	return true, user.password

}

func deleteCookie(w http.ResponseWriter){
	cookie := http.Cookie{Name: "username", MaxAge: -1}
	http.SetCookie(w, &cookie)
	debugPrint("Debug:Cookie Deleted")
	return
}