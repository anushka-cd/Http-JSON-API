//19BIT0279(Anushka R)
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Upload is a struct that represents a user in our application

type Upload struct {
	ID       string `json:"ID"`
	Caption  string `json:"caption"`
	ImageURL string `json:"imageurl"`
}

//User is a struct that represents a user in our application

type User struct { //exported
	ID       string `json:"ID"`
	Name     string `json:"name"`
	Password string `json:"password"`
}

//Post is a struct that represents a single post
type Post struct {
	Author       User     `json:"author"`
	UploadedPost []Upload `json:"uploadedpost"`
}

//Create a slice
var posts []Post = []Post{}
var users []User = []User{}

func main() {
	/*ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, _ = mongo.Connect(ctx, options.Client().ApplyURI("mongo://localhost:27017"))*/
	router := mux.NewRouter()
	//router.HandleFunc("/test", test)
	//router.HandleFunc("/add/{item}", addItem).Methods("POST") //route with dynamic variable
	router.HandleFunc("/users", CreateUser).Methods("POST")
	router.HandleFunc("/users/{id}", getUser).Methods("GET")
	router.HandleFunc("/posts", postItem).Methods("POST")
	router.HandleFunc("/posts", getAllPosts).Methods("GET")
	router.HandleFunc("/posts/users/{id}", getPost).Methods("GET")
	router.HandleFunc("/posts/users/{id}", updatePost).Methods("PUT")
	router.HandleFunc("/posts/users/{id}", patchPost).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":4000", router))

}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsons")

	//routeVariable := mux.Vars(r)["item"] //Vars returns the route variables for the current request, if any.
	//get Item value from JSON body
	var newUser User
	json.NewDecoder(r.Body).Decode(&newUser)

	users = append(users, newUser)
	json.NewEncoder(w).Encode(users) //return back the slice to caller
}
func getUser(w http.ResponseWriter, r *http.Request) {
	//get the id of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) //convert string to  int,returns error and integer
	if err != nil {
		//there is an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer")) //string to slice of bytes
		return
	}

	//error checking
	if id >= len(users) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	user := users[id]
	w.Header().Set("Content-Type", "application/jsons")
	json.NewEncoder(w).Encode(user)
}

func getPost(w http.ResponseWriter, r *http.Request) {
	//get the id of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam) //convert string to  int,returns error and integer
	if err != nil {
		//there is an error
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer")) //string to slice of bytes
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}

	post := posts[id]
	w.Header().Set("Content-Type", "application/jsons")
	json.NewEncoder(w).Encode(post)
}
func getAllPosts(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsons")
	json.NewEncoder(w).Encode(posts)
}

func postItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/jsons")

	//routeVariable := mux.Vars(r)["item"] //Vars returns the route variables for the current request, if any.
	//get Item value from JSON body

	var newPost Post
	json.NewDecoder(r.Body).Decode(&newPost)

	posts = append(posts, newPost)
	json.NewEncoder(w).Encode(posts) //return back the slice to caller
}
func updatePost(w http.ResponseWriter, r *http.Request) {

	//get the id of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}
	//get the value from the json body in the request
	var updatePost Post
	w.Header().Set("Content-Type", "application/jsons")
	json.NewDecoder(r.Body).Decode(&updatePost)
	posts[id] = updatePost
	json.NewEncoder(w).Encode(updatePost)

}
func patchPost(w http.ResponseWriter, r *http.Request) {
	//get the id of the post from the route parameter
	var idParam string = mux.Vars(r)["id"]
	id, err := strconv.Atoi(idParam)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte("ID could not be converted to integer"))
		return
	}

	//error checking
	if id >= len(posts) {
		w.WriteHeader(404)
		w.Write([]byte("No post found with specified id"))
		return
	}
	//get the current value
	post := &posts[id]
	json.NewDecoder(r.Body).Decode(post)

	w.Header().Set("Content-Type", "application/jsons")
	json.NewEncoder(w).Encode(post)
}
