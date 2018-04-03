package main

import (
    "fmt"
    "html/template"
    "log"
    "net/http"
    "strings"
)

//handler to deal with only / requests. Default behaviour needs to be defined
func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Sayhello Handler")
    r.ParseForm() //Parse url parameters passed, then parse the response packet for the POST body (request body)
    // attention: If you do not call ParseForm method, the following data can not be obtained form
    fmt.Println(r.Form) // print information on server side.
    fmt.Println("path", r.URL.Path)
    fmt.Println("scheme", r.URL.Scheme)
    fmt.Println(r.Form["url_long"])
    for k, v := range r.Form {
        fmt.Println("key:", k)
        fmt.Println("val:", strings.Join(v, ""))
    }
    fmt.Fprintf(w, "Hello astaxie!") // write data to response
}

//Handler for login
func login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: login handler")
    fmt.Println("Debug: login method:", r.Method)
    if r.Method == "GET" { 											//user has directly come to login page
        t, _ := template.ParseFiles("login.php")
        t.Execute(w, nil)
    } else {  														//user has come to login via post
        r.ParseForm()
        fmt.Println("username:", r.Form["username"])
        fmt.Println("password:", r.Form["password"])
        usr := r.Form["username"][0]
        pwd := r.Form["password"][0]
        ok, actualpassword := getPassword(usr)

        if(!ok){													//user does not exist - send to registration page
        	http.Redirect(w,r,"/registration",http.StatusSeeOther)
        	return
		}

		if(pwd==actualpassword){									//login successful goto home
			expiration := 3600										//set cookie to validate other pages
			cookie := http.Cookie{Name: "username", Value: usr, MaxAge: expiration}
			http.SetCookie(w, &cookie)
			http.Redirect(w,r,"/home",http.StatusSeeOther)
			return
		}else{														//login unsuccessful go back to login page
			http.Redirect(w,r,"/login",http.StatusSeeOther)
			return
		}
    }
}

//handler for registration
func registration(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: Registration handler")
	fmt.Println("method:", r.Method) //get request method
	if r.Method == "GET" {
		t, _ := template.ParseFiles("Registration.php")
		t.Execute(w, nil)
		return
	} else {															//arrived by post
		r.ParseForm()
		if debugon{ fmt.Println("Debug: username in post: ",r.Form["username"][0]) }
		if debugon{ fmt.Println("Debug: password in post: ",r.Form["password_1"][0]) }
		result := addUser(r.Form["username"][0],r.Form["password_1"][0])
		if result==1 {													//succesfully added user go to login page
			http.Redirect(w,r,"/login",http.StatusSeeOther)
			return
		}else{															//user already exists send back to registration page
			if debugon { fmt.Println("Debug: User already exists") }
			http.Redirect(w,r,"/registration",http.StatusSeeOther)
			return
		}

	}
	}

//handlr for home
func home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: home handler")
	fmt.Println("method:", r.Method) //get request method
	cookie, ok := r.Cookie("username")
	fmt.Println(cookie)
	fmt.Println(ok)
	if(ok!=nil){						//cookie does not exist re-direct to login
		http.Redirect(w,r,"/login",http.StatusSeeOther)
		return
	}
	username := cookie.Value
	user := userdata[username]//username of logged in user, key to userdata hashmap
	if r.Method == "GET" {					//re-direct to homepage from login. Display Home
		t, _ := template.ParseFiles("home.php")
		t.Execute(w, nil)
	} else {								//Post: submission of new tweet. Save the tweet and then display Home.
		r.ParseForm()
		newtweet := tweet{text:r.Form["tweet"][0]}
		user.tweets = append(user.tweets,newtweet)
		userdata[username]=user
		t, _ := template.ParseFiles("home.php")
		t.Execute(w, nil)
		fmt.Println(user.tweets)

	}
	fmt.Fprint(w, "<h>Here are your tweets:<h><br />")
	for _, disptweet := range user.tweets{
		fmt.Fprint(w,disptweet.text)
		fmt.Fprint(w,"<br />")
	}
	//fmt.Fprint(w,user.tweets)


}

//logout handler
func logout(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: logout handler")
	fmt.Println("method:", r.Method) //get request method
	deleteCookie(w)
	http.Redirect(w,r,"/login",http.StatusSeeOther)
}

func users(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: users handler")
	fmt.Println("method:", r.Method) //get request method
	cookie, ok := r.Cookie("username")
	if(ok!=nil){						//cookie does not exist re-direct to login
		http.Redirect(w,r,"/login",http.StatusSeeOther)
		return
	}
	username := cookie.Value
	user := userdata[username]
	t, _ := template.ParseFiles("home.php")
	t.Execute(w, nil)
	fmt.Fprint(w, "<h>Folow some Users to see their tweets<h><br/>")
	for k,_ := range userdata{
		_, ok := user.follows[k]
		fmt.Println(k)
		if(ok==false && k!=user.username){
			//user is not already following the person. Checking if the user1 from userdata exists in current users follows
			//fmt.Fprint(w, k)
			fmt.Fprintf(w,"%s <a href=follow?tofollow=%s>Follow</a>",k,k)
			fmt.Fprint(w,"</br>")
		}

	}

}

func follow(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Debug: follow handler")
	fmt.Println("method:", r.Method) //get request method
	/*cookie, ok := r.Cookie("username")
	if(ok!=nil){						//cookie does not exist re-direct to login
		http.Redirect(w,r,"/login",http.StatusSeeOther)
		return
	}
	username := cookie.Value
	user := userdata[username]*/
	tofollow := r.URL.Query().Get("tofollow")
	fmt.Fprint(w,tofollow)
	//user.follows[tofollow]=true
	//http.Redirect(w,r,"/users",http.StatusSeeOther)

}



//entry point.
func main() {
    http.HandleFunc("/", sayhelloName) // setting router rule
    http.HandleFunc("/login", login)
	http.HandleFunc("/registration", registration)
	http.HandleFunc("/home", home)
	http.HandleFunc("/logout", logout)
	http.HandleFunc("/users", users)
	http.HandleFunc("/follow", follow)
    err := http.ListenAndServe(":9090", nil) // setting listening port
    if err != nil {
        log.Fatal("ListenAndServe: ", err)
    }
}