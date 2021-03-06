package main

import (
	"fmt"
)
  var graph = make(map[string]map[string]bool)
func main() {	
	addEdge("apple","tree")
	fmt.Println(hasEdge("apple","tree"))
	fmt.Println(hasEdge("da vinci","tree"))
}

// graph by maps
func addEdge(start, end string){
        edges := graph[start]
	if edges == nil{
		edges = make(map[string]bool)
		graph[start] = edges
	}
	edges[end] = true
}
func hasEdge(start, end string)bool{
	return graph[start][end]
}

/* CASE-1 DATA NOT EXIT IT WILL RETURN false
   CASE-2 Data exit it always return true */
https://play.golang.org/p/lyZsmnpw1Tu




