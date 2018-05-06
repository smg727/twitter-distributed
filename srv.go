package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strings"
	"os"
	"time"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "twitter-distributed/utils/ProtoDef"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

var rpcCaller pb.GreeterClient

//Handler to deal with only / requests.
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Sayhello Handler")
	//Parse url parameters passed, then parse the response packet for the POST body (request body)
	r.ParseForm()
	// Print information on server side.
	fmt.Println(r.Form)
	fmt.Println("path", r.URL.Path)
	fmt.Println("scheme", r.URL.Scheme)
	fmt.Println(r.Form["url_long"])
	for k, v := range r.Form {
		fmt.Println("key:", k)
		fmt.Println("val:", strings.Join(v, ""))
	}
	// write data to response
	fmt.Fprintf(w, "Hello User!")
}

//Handler for login
func loginHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: login handler")
	fmt.Println("Debug: login method:", r.Method)

	if r.Method == "GET" {
		//User has directly come to login page
		t, _ := template.ParseFiles("login.html")
		t.Execute(w, nil)
	} else {
		//User has come to login via Post
		r.ParseForm()
		fmt.Println("username:", r.Form["username"])
		fmt.Println("password:", r.Form["password"])
		usr := r.Form["username"][0]
		pwd := r.Form["password"][0]
		//ok, actualPassword := getPassword(usr)
		//calling rpc to validate user
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		reply, err := rpcCaller.Login(ctx, &pb.Credentials{Uname: usr, Pwd: pwd})
		//User does not exist - send to registration page
		if err != nil {
			fmt.Println("Debug: Login rpc failed", err.Error())
			if err.Error() == "Wrong Password" {
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			} else {
				http.Redirect(w, r, "/registration", http.StatusSeeOther)
				return
			}

		} else if err == nil && reply.Status == true {
			debugPrint("debug: user successfully logged in")
			expiration := 3600
			cookie := http.Cookie{Name: "username", Value: usr, MaxAge: expiration}
			http.SetCookie(w, &cookie)
			http.Redirect(w, r, "/home", http.StatusSeeOther)
			return
		} else {
			log.Println("Major issue")
		}
	}
}

//Handler for registration
func registrationHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Registration handler")
	//Get request method, type of request
	fmt.Println("method:", r.Method)

	if r.Method == "GET" {
		t, _ := template.ParseFiles("Registration.html")
		t.Execute(w, nil)
		return
	} else {
		r.ParseForm()
		if debugon {
			fmt.Println("Debug: username in post: ", r.Form["username"][0])
			fmt.Println("Debug: password in post: ", r.Form["password_1"][0])
		}

		//Check for non-empty username and password values
		if len(r.Form["username"][0]) == 0 || len(r.Form["password_1"][0]) == 0 {
			if debugon {
				fmt.Println("Debug: Empty Username or Password value")
			}
			//TODO: Remove Alert?
			//fmt.Fprintln(w, "<script>alert(\"Please enter a valid Username and Password\")</script>")
			http.Redirect(w, r, "/registration", http.StatusSeeOther)
			return
		}

		//calling rpc to add user
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		reply, err := rpcCaller.Register(ctx, &pb.Credentials{Uname: r.Form["username"][0], Pwd: r.Form["password_1"][0], Broadcast: true})
		if err == nil {
			fmt.Println("User added using rpc", reply)
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		} else {
			fmt.Println("RPC to master failed", reply, err)
			http.Redirect(w, r, "/registration", http.StatusSeeOther)
			return
		}

		//Adding username and password to the Map

	}
}

//Home Page handler
func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: home handler")
	fmt.Println("method:", r.Method) //get request method
	cookie, ok := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Println(ok)

	//Cookie does not exist re-direct to login
	if ok != nil {
		fmt.Println("Debug: Cookie doesnt exist. re-direct to login")
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//Get User Data from Map
	username := cookie.Value
	isUserPresent := userExists(username)

	//If map returns false, the account has been deleted. Redirect to registration.
	if !isUserPresent {
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}

	if r.Method == "GET" {
		//Re-direct to homepage from login. Display Home
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, nil)
	} else {
		//Post: submission of new tweet. Save the tweet and then display Home.
		r.ParseForm()
		addTweet(username, r.Form["tweet"][0])
		t, _ := template.ParseFiles("home.html")
		t.Execute(w, nil)

	}

	//Display self tweets
	tweets := getMyTweets(username).TweetList
	if len(tweets) != 0 {
		fmt.Fprint(w, "<h>Here are your tweets, "+username+":<h><br />")
	} else {
		fmt.Fprint(w, "<h>What's on your mind? Make a tweet ! <h><br />")
	}
	for _, dispTweet := range tweets {
		fmt.Fprint(w, "<p>")
		fmt.Fprint(w, dispTweet.Text)
		fmt.Fprint(w, "</p>")
	}

	// Initiate RPC call to get all the friends tweets
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := rpcCaller.GetFriendsTweets(ctx, &pb.GetFriendsTweetsRequest{Username: username})
	if err == nil {
		fmt.Println("Successful GetFriendsTweets RPC")
		allFriendsTweets := reply.FriendsTweets
		fmt.Fprint(w, "<br /><br />")
		if len(allFriendsTweets) != 0 {
			fmt.Fprint(w, "<h>Here are your friends tweets:<h><br />")
		} else {
			fmt.Fprint(w, "<h>Go right ahead, Discover some users to follow<h><br />")
		}

		for _, eachUser := range allFriendsTweets {
			//Print Friend's name
			fmt.Fprint(w, "<br/>"+eachUser.Username.Username+":"+"<br/>")
			for _, eachTweet := range eachUser.Tweets {
				//Print Friend's all the tweets
				fmt.Fprint(w, "<p>")
				fmt.Fprint(w, eachTweet.Text)
				fmt.Fprint(w, "</p>")
			}
		}
		return
	} else {
		fmt.Println("GetFriendsTweets RPC failed", reply, err)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

}

//Logout handler
func logoutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: logout handler")
	//Print request method
	fmt.Println("method:", r.Method)
	deleteCookie(w)
	http.Redirect(w, r, "/login", http.StatusSeeOther)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: users handler")
	//Print request method
	fmt.Println("method:", r.Method)
	cookie, ok := r.Cookie("username")
	if ok != nil {
		//Cookie does not exist re-direct to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//Get User Data from Map
	username := cookie.Value
	isUserPresent := userExists(username)

	//If map returns false, the account has been deleted. Redirect to registration.
	if !isUserPresent {
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	}

	t, _ := template.ParseFiles("Home.html")
	t.Execute(w, nil)
	toFollow := r.URL.Query().Get("tofollow")
	fmt.Println("tofollow" + toFollow)

	if toFollow != "" {
		ctx, cancel := context.WithTimeout(context.Background(), time.Second)
		defer cancel()
		reply, err := rpcCaller.FollowUser(ctx, &pb.FollowUserRequest{SelfUsername: username, ToFollowUsername: toFollow, Broadcast:true})
		if err == nil {
			fmt.Println("User " + username + " successfully followed user " + toFollow)
		} else {
			fmt.Println("FollowUser RPC failed", reply, err)
			http.Redirect(w, r, "/users", http.StatusSeeOther)
			return
		}

	}

	fmt.Fprint(w, "<h>Follow some Users to see their Tweets<h><br/>")

	//RPC Call to get all the users to follow
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	reply, err := rpcCaller.UsersToFollow(ctx, &pb.UsersToFollowRequest{Username: username})
	if err == nil {
		fmt.Println("UsersToFollow RPC Sucessful", reply)
		allUsersToFollow := reply.UsersToFollowList
		for _, eachUser := range allUsersToFollow {
			//Adding all the users to follow on the website
			fmt.Fprintf(w, "%s <a href=users?tofollow=%s>Follow</a>", eachUser.Username, eachUser.Username)
			fmt.Fprint(w, "</br>")
		}
	} else {
		fmt.Println("UsersToFollow RPC failed", reply, err)
		http.Redirect(w, r, "/home", http.StatusSeeOther)
		return
	}

}

//Delete Account handler
func deleteHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Invoked Delete Handler")
	//get request method
	fmt.Println("Method:", r.Method)

	//Get cookie to identify the user
	cookie, ok := r.Cookie("username")
	if ok != nil {
		//Cookie does not exist, re-direct to login
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}
	username := cookie.Value

	//Remove user from the user Map
	deleteStatus := deleteUser(username)
	if deleteStatus == 0 {
		//Successfully deleted the User
		deleteCookie(w)
		http.Redirect(w, r, "/registration", http.StatusSeeOther)
		return
	} else {
		//Deleting user Failed
		http.Redirect(w, r, "/login", http.StatusSeeOther)
		return
	}

	//Delete cookie and redirect to register
	deleteCookie(w)
	http.Redirect(w, r, "/registration", http.StatusSeeOther)
}

func faviconHandler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {

	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	//c := pb.NewGreeterClient(conn)
	rpcCaller = pb.NewGreeterClient(conn)

	// Contact the server and print out its response. TO test if RPC is working
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := rpcCaller.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("Please bring up the Back-End Server first. Could not greet: %v", err)
	}
	log.Printf("RPC is working %s", r.Message)
	//end of test RPC

	//All handler functions
	http.HandleFunc("/", sayhelloName) //Keeping this for now to enable log analyzing in console. Lets change this later
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/registration", registrationHandler)
	http.HandleFunc("/home", homeHandler)
	http.HandleFunc("/logout", logoutHandler)
	http.HandleFunc("/users", usersHandler)
	http.HandleFunc("/deleteAccount", deleteHandler)
	http.HandleFunc("/favicon.ico", faviconHandler)

	//Our server listens on this port
	errls := http.ListenAndServe(":9090", nil)
	if errls != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
