package main

import (
	"fmt"
	"net/http"
)

func userHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	fmt.Println(w, "Success")
	userId := r.FormValue("user-id")
	name := r.FormValue("name")

	fmt.Println(userId, name)
}

func groupHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "err: %v", err)
		return
	}
	fmt.Println(w, "Success")
	groupId := r.FormValue("group-id")
	groupName := r.FormValue("group-name")

	fmt.Println(groupId, groupName)
}

func main() {
	http.HandleFunc("/user", userHandler)
	http.HandleFunc("/group", groupHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		fmt.Println("err: %v", err)
	}
}
