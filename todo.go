package main

type Todo struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Status string `json:"status"`
}

var todos = []Todo{
	{ID: "1", Title: "gym", Desc: "goto gym", Status: "Pending"},
	{ID: "2", Title: "bar", Desc: "goto bar", Status: "Pending"},
}
