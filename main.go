package main

import (
    "encoding/json"
    "net/http"
    "sync"
)

type User struct {
    ID       string `json:"id"`
    Name     string `json:"name"`
    Island   string `json:"island"`
    Character string `json:"character"`
}

var users = make(map[string]User)
var mu sync.Mutex

func addUser(w http.ResponseWriter, r *http.Request) {
    var user User
    if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    mu.Lock()
    users[user.ID] = user
    mu.Unlock()

    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(user)
}
func getUser(w http.ResponseWriter, r *http.Request) {
    id := r.URL.Query().Get("id")
    mu.Lock()
    user, ok := users[id]
    mu.Unlock()

    if !ok {
        http.Error(w, "User not found", http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(user)
}
func listUsers(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()
    var userList []User
    for _, user := range users {
        userList = append(userList, user)
    }
    json.NewEncoder(w).Encode(userList)
}
func setupRoutes() {
    http.HandleFunc("/addUser", addUser)
    http.HandleFunc("/getUser", getUser)
    http.HandleFunc("/listUsers", listUsers)
}
