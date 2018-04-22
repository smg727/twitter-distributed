package main

import (
	"fmt"
	"golang.org/x/net/context"
	pb "twitter-distributed/utils/ProtoDef"
	"net/http"
	"time"
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

//Function to print debug outputs if debugon=true
func debugPrint(text string){
	if(debugon){
		fmt.Println(text)
	}
}

func userExists(uname string) bool{
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := rpcCaller.UserExists(ctx, &pb.UserExistsRequest{Username:uname})
	if err==nil {
		return reply.Status
	}
	fmt.Println("Debug: userExists rpc returned false",err)
	return false
}

func addTweet(username string, tweettext string) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	_, err := rpcCaller.AddTweet(ctx, &pb.AddTweetRequest{Username:username, TweetText:tweettext})
	if(err!=nil){
		fmt.Println("Debug: tweet addition failed",err)
	}

}

func getMyTweets(username string) *pb.OwnTweetsReply  {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := rpcCaller.OwnTweets(ctx, &pb.OwnTweetsRequest{Username:username})
	if err!=nil {
		fmt.Println(err)
		return nil
	}
	return reply
}
//function to add user to data on registration
func addUser(usrname string, pwd string) int  {
	_, ok := userdata[usrname]
	if(ok){
		debugPrint("Debug: User already exists")
		return 0
	}
	usr := User{username:usrname,password:pwd}
	usr.follows = make(map[string]bool)
	userdata[usrname] = usr
	debugPrint("Debug: User added")
	return 1
}

//Delete a user account
func deleteUser(username string) int  {
	//TODO: for later stages, we might have to add Locks here
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := rpcCaller.DeleteUser(ctx, &pb.Credentials{Uname:username})
	if err == nil {
		fmt.Println("Delete User RPC successful",reply)
		return 0
	}else{
		fmt.Println("Delete User RPC failed",reply,err)
		return -1
	}
}

//Returns users password
func getPassword(usrname string) (bool, string){

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