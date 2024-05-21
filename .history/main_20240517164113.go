package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	ID int
	X  int
	Y  int
}

type Edge struct {
	Start int
	End   int
	Cost  int // Kenar maliyeti
}

type Graph struct {
	Nodes       []Node
	Edges       []Edge
	StartNodeID int
	EndNodeID   int
	AntCounts   map[int]int
	AdjList     map[int][]Edge // Komşuluk listesi
}

func (g *Graph) Dijkstra(startNodeID int) map[int]int {
	dist := make(map[int]int)
	for _, node := range g.Nodes {
		dist[node.ID] = int(^uint(0) >> 1) // Max int
	}
	dist[startNodeID] = 0

	pq := &PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, &Item{nodeID: startNodeID, priority: 0})

	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		current := item.nodeID

		for _, edge := range g.AdjList[current] {
			next := edge.End
			newDist := dist[current] + edge.Cost
			if newDist < dist[next] {
				dist[next] = newDist
				heap.Push(pq, &Item{nodeID: next, priority: newDist})
			}
		}
	}
	return dist
}

// BFS Algoritması
func (g *Graph) BFS(startNodeID int, endNodeID int) [][]int {
	paths := [][]int{}
	queue := [][]int{{startNodeID}}
	visited := map[int]bool{}

	for len(queue) > 0 {
		path := queue[0]
		queue = queue[1:]
		node := path[len(path)-1]

		if node == endNodeID {
			paths = append(paths, path)
			continue
		}

		if visited[node] {
			continue
		}

		visited[node] = true

		for _, edge := range g.AdjList[node] {
			if !visited[edge.End] {
				newPath := make([]int, len(path))
				copy(newPath, path)
				newPath = append(newPath, edge.End)
				queue = append(queue, newPath)
			}
		}
	}
	return paths
}

// Düğümleri yazdırmak için fonksiyon
func printNodes(nodes []Node) {
	fmt.Println("\n the_rooms:")
	for _, node := range nodes {
		fmt.Printf("%d: (%d, %d)\n", node.ID, node.X, node.Y)
	}
}

// Kenarları yazdırmak için fonksiyon
func printEdges(edges []Edge) {
	fmt.Println("\nthe_links:")
	for _, edge := range edges {
		fmt.Printf("%d - %d\n", edge.Start, edge.End)
	}
}

// PriorityQueue item'ları
type Item struct {
	nodeID   int
	priority int
	index    int
}

// PriorityQueue yapısı
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	item.index = -1
	*pq = old[0 : n-1]
	return item
}


	
	// BFS Algoritması ile başlangıçtan bitişe yolları bul
	paths := graph.BFS(graph.StartNodeID, graph.EndNodeID)
	fmt.Println("BFS ile bulunan yollar:", paths)

	// Karınca hareketlerini yazdırmak için mantık
	if len(paths) > 0 {
		antPaths := make([][]int, antCount)
		for i := 0; i < antCount; i++ {
			antPaths[i] = append(antPaths[i], graph.StartNodeID)
		}

		for i := 1; i < len(paths[0]); i++ {
			for j := 0; j < antCount; j++ {
				if i < len(antPaths[j]) {
					antPaths[j][i] = paths[j%len(paths)][i]
				} else {
					antPaths[j] = append(antPaths[j], paths[j%len(paths)][i])
				}
			}
		}

		occupied := make(map[int]bool)
		for i := 0; i < len(paths[0]); i++ {
			for j := 0; j < antCount; j++ {
				if i < len(antPaths[j]) {
					if !occupied[antPaths[j][i]] {
						fmt.Printf("L%d-%d ", j+1, antPaths[j][i])
						occupied[antPaths[j][i]] = true
					} else {
						// Başka yol bulma mantığı eklenebilir
						fmt.Printf("L%d-%d ", j+1, antPaths[(j+1)%antCount][i])
						occupied[antPaths[(j+1)%antCount][i]] = true
					}
				}
			}
			fmt.Println()
			occupied = make(map[int]bool) // Reset the occupied map for the next step
		}
	} else {
		fmt.Println("Yol bulunamadı.")
	}
}

func (g *Graph) CheckAntCount(nodeID int) int {
	antCount, exist := g.AntCounts[nodeID]
	if !exist {
		return 0
	}
	return antCount
}

func (g *Graph) UpdateAntCounts(newAntCounts map[int]int) {
	for id, count := range newAntCounts {
		g.AntCounts[id] = count
	}
}
// Karıncaların hedef odalara ulaşmasını sağlar
func main() {
    // Önceki kod kısmı buraya gelecek

    // BFS Algoritması ile başlangıçtan bitişe yolları bul
    paths := graph.BFS(graph.StartNodeID, graph.EndNodeID)
    fmt.Println("BFS ile bulunan yollar:", paths)

    // Karıncaların hedefe ulaşana kadar ilerlemesi
    if len(paths) > 0 {
        antPaths := make([][]int, antCount)
        for i := 0; i < antCount; i++ {
            antPaths[i] = make([]int, len(paths[i])) // Her karınca için yol oluştur
            copy(antPaths[i], paths[i]) // BFS ile bulunan yolları kopyala
        }

        occupied := make(map[int]bool)
        for i := 0; i < len(paths[0]); i++ {
            for j := 0; j < antCount; j++ {
                if i < len(antPaths[j]) {
                    if !occupied[antPaths[j][i]] {
                        fmt.Printf("L%d-%d ", j+1, antPaths[j][i])
                        occupied[antPaths[j][i]] = true
                    } else {
                        // Başka yol bulma mantığı eklenebilir
                        fmt.Printf("L%d-%d ", j+1, antPaths[(j+1)%antCount][i])
                        occupied[antPaths[(j+1)%antCount][i]] = true
                    }
                }
            }
            fmt.Println()
            occupied = make(map[int]bool) // Reset the occupied map for the next step
        }
    } else {
        fmt.Println("Yol bulunamadı.")
    }
}
