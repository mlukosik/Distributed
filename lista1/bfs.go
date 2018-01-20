package main

import (
	"bufio"
    "os"
	"fmt"
	"strconv"
	"regexp"
	"time"
)

type vertexStruct struct {
		vertexID int
		parentID int
		mainID int
}


type singleVertex struct {
		id int
		parent int
		children []int
}


func neigh(i int, ch chan vertexStruct, adj_list map[int][]vertexStruct, visited []int, parent []int, vertexArray []singleVertex) {

	for j := 0; j < len(adj_list[i]); j++ {
		visited[i] = 1

		if parent[adj_list[i][j].vertexID - 1] == 0 {
			adj_list[i][j].mainID = i
			adj_list[i][j].parentID = i
			vertexArray[i-1].children = append(vertexArray[i-1].children, adj_list[i][j].vertexID)
			go bfs(adj_list[i][j].vertexID, ch, adj_list[i][j], visited, i, adj_list, parent, vertexArray)
		}
	}
}


func bfs(i int, ch chan vertexStruct, adj vertexStruct, visited []int, j int, adj_list map[int][]vertexStruct, parent []int, vertexArray []singleVertex) {

	if visited[i] == 0 {
		parent[adj.vertexID-1] = j
		vertexArray[adj.vertexID-1].parent = j
		ch <- adj
		go neigh(i, ch, adj_list, visited, parent, vertexArray)
	}
}


func main() {

	r, _ := regexp.Compile("[0-9]+")	

	file, err := os.Open("test2")
	defer file.Close()
    if err != nil {
		fmt.Println("ERR: ", err)
    }

	reader := bufio.NewReader(file)
	var line string
	var vertex int

	line, err = reader.ReadString('\n')
	temp := r.FindAllString(line, -1)
	n, err := strconv.Atoi(temp[0])

	adj_list := make(map[int][]vertexStruct)
	visited := make([]int, n+1)
	parent := make([]int, n)
	vertexArray := make([]singleVertex, n)

	visited[1] = 0
	for i := 1; i <= n; i++ {
		visited[i-1] = 0 
		line, err = reader.ReadString('\n')
		temp = r.FindAllString(line, -1)
		for j := 1; j < len(temp); j++ {
			vertex, err = strconv.Atoi(temp[j])
			str := vertexStruct{vertexID: vertex, parentID: 0, mainID: i}
			adj_list[i] = append(adj_list[i], str)
			vertexArray[i-1] = singleVertex{id: i}
		}
	}
	vertexArray[0].parent = 1
	visited[1] = 1
	parent[0] = 1

	channel := make(chan vertexStruct)

	go neigh(1, channel, adj_list, visited, parent, vertexArray)

	for i := 0; i < n-1; i++ {
		time.Sleep(1* time.Millisecond)
		select {
			case msg := <-channel:
					parent[msg.vertexID-1] = msg.parentID
			default:
				fmt.Println("ok")
		}
	}

	for i := range vertexArray {
		fmt.Printf("Wierzchołek: %d Rodzic: %d Dzieci: %v\n", vertexArray[i].id, vertexArray[i].parent, vertexArray[i].children)
	}
}


