package main

import (
    "encoding/json"
    "net/http"
    "strconv"
    "sync"

    "github.com/gorilla/mux"
)

type User struct {
    ID    int    `json:"id"`
    Name  string `json:"name"`
    Email string `json:"email"`
}

var (
    users   = make(map[int]User)
    nextID  = 1
    userMux sync.Mutex
)

func createUser(w http.ResponseWriter, r *http.Request) {
    var u User
    json.NewDecoder(r.Body).Decode(&u)
    userMux.Lock()
    u.ID = nextID
    nextID++
    users[u.ID] = u
    userMux.Unlock()
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(u)
}

func getAllUsers(w http.ResponseWriter, r *http.Request) {
    userMux.Lock()
    defer userMux.Unlock()
    var all []User
    for _, u := range users {
        all = append(all, u)
    }
    json.NewEncoder(w).Encode(all)
}

func getUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    userMux.Lock()
    defer userMux.Unlock()
    u, ok := users[id]
    if !ok {
        w.WriteHeader(http.StatusNotFound)
        return
    }
    json.NewEncoder(w).Encode(u)
}

func updateUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    var u User
    json.NewDecoder(r.Body).Decode(&u)
    userMux.Lock()
    defer userMux.Unlock()
    u.ID = id
    users[id] = u
    json.NewEncoder(w).Encode(u)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
    id, _ := strconv.Atoi(mux.Vars(r)["id"])
    userMux.Lock()
    defer userMux.Unlock()
    delete(users, id)
    w.WriteHeader(http.StatusNoContent)
}

func deleteAllUsers(w http.ResponseWriter, r *http.Request) {
    userMux.Lock()
    defer userMux.Unlock()
    users = make(map[int]User)
    w.WriteHeader(http.StatusNoContent)
}

func main() {
    r := mux.NewRouter()
    r.HandleFunc("/users", createUser).Methods("POST")
    r.HandleFunc("/users", getAllUsers).Methods("GET")
    r.HandleFunc("/users/{id}", getUser).Methods("GET")
    r.HandleFunc("/users/{id}", updateUser).Methods("PUT")
    r.HandleFunc("/users/{id}", deleteUser).Methods("DELETE")
    r.HandleFunc("/users", deleteAllUsers).Methods("DELETE")
    http.ListenAndServe(":8080", r)
}