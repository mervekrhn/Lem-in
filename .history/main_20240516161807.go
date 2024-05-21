package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type Node struct { //dügümler için struct yapısı
	ID int
	X  int
	Y  int
}

type Edge struct { //kenarlar için struct yapısı
	Start int
	End   int
}

type Graph struct { //düğümler ve knarlar için oluşturulmuş bir graf struct yapısı
	Nodes       []Node
	Edges       []Edge
	StartNodeID int         // Başlangıç düğümünün ID'si
	EndNodeID   int         // Bitiş düğümünün ID'si
	AntCounts   map[int]int // dugum ıd sini ve karınca sayıısıını eslestirdigimiz harita
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

func main() {

	file, err := os.Open("text.txt") //dosya açma işlemi
	if err != nil {
		fmt.Println("Dosya açma hatasi:", err)
		return
	}
	defer file.Close()

	graph := Graph{
		AntCounts: make(map[int]int),
	}

	nodeID := 1
	antCount := graph.CheckAntCount(nodeID)
	fmt.Println("Node", nodeID, "deki karınca sayısı:", antCount)
	startID, endID := -1, -1 //ilk önce baslangic ve bitis satırlarına bi id atıyoruz

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if antCount == 0 { // Karınca sayısı daha önce belirlenmediyse
			antCount, err = strconv.Atoi(line) // Dosyadan karınca sayısını oku
			if err != nil {
				return
			}
			fmt.Println("number_of_ants :", antCount)
		}
		if strings.HasPrefix(line, "##start") { //içinde start var mı diye kontrol ediyor
			scanner.Scan()
			fields := strings.Fields(scanner.Text())
			startID, _ = strconv.Atoi(fields[0]) //ilk önce atadığımız id değiştiği için onun atamasını yapıyoruz
			startx, _ := strconv.Atoi(fields[1]) // başlangıç düğümünün x koordinatını al
			starty, _ := strconv.Atoi(fields[2]) // başlangıç düğümünün y koordinatını al
			graph.StartNodeID = startID

			graph.Nodes = append(graph.Nodes, Node{ID: startID, X: startx, Y: starty}) // başlangıç düğümünü düğümlere ekle
			fmt.Println("start_room:", graph.StartNodeID)

		} else if strings.HasPrefix(line, "##end") { //satır bitis satırı mı diye kontrol ediyor
			scanner.Scan()
			fields := strings.Fields(scanner.Text())
			endID, _ = strconv.Atoi(fields[0]) //ilk önce atadığımız id değiştiği için onun atamasını yapıyoruz
			endx, _ := strconv.Atoi(fields[1]) // bitiş düğümünün x koordinatını al
			endy, _ := strconv.Atoi(fields[2])
			graph.EndNodeID = endID
			graph.Nodes = append(graph.Nodes, Node{ID: endID, X: endx, Y: endy}) // bitiş düğümünü düğümlere ekle
			fmt.Println("end_room:", graph.EndNodeID)

		} else {
			fields := strings.Fields(line)
			if len(fields) == 3 { //eger satır 3 argumandan oluşuyorsa o zaman  bunun bir dügüm koordinatı belirtme olduğunu anlayıp onaa göre atama yapıyoruz
				id, _ := strconv.Atoi(fields[0])
				x, _ := strconv.Atoi(fields[1])
				y, _ := strconv.Atoi(fields[2])
				graph.Nodes = append(graph.Nodes, Node{ID: id, X: x, Y: y}) //grafdaki dügüm struct yapısına ekleme yapıyoruz çünkü artık id,x,y yapılarımız belirlendi
			} else if len(fields) == 1 && strings.Contains(line, "-") == false {
				// Eğer satır sadece bir sayıdan oluşuyorsa ve "-" karakteri içermiyorsa, bu karınca sayısıdır.
				antCount, _ = strconv.Atoi(fields[0])
			} else if len(fields) == 1 && strings.Contains(line, "-") {
				edgeParts := strings.Split(fields[0], "-") //- ile ayrılmış olanları alıyoruz
				start, _ := strconv.Atoi(edgeParts[0])     //bunlardan ilki bizim için başlangıç ikincisi bitiş olaark alınıyor çünkü aradaki uzunluğu bulmak için kullanıcaz
				end, _ := strconv.Atoi(edgeParts[1])
				graph.Edges = append(graph.Edges, Edge{Start: start, End: end}) //kenarlar struct  yapısına ekleme yapıyoruz çünkü artık başlangıç ve bitiş noktaları belirlendi
			}
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Okuma hatasi:", err)
		return
	}

	// Başlangıç ve bitiş düğümlerini kontrol et
	if startID == -1 || endID == -1 {
		fmt.Println("Hata: Başlangiç veya bitiş düğümü belirtilmemiş.")
		return
	}
	fmt.Println("son durum:")
	// Düğümleri ve kenarları yazdır
	printNodes(graph.Nodes)
	printEdges(graph.Edges)

	newAntCounts := make(map[int]int)
	for _, node := range graph.Nodes {
		// Her düğümdeki karınca sayısını hesapla ve yeni karınca sayılarına ekle
		yeniAntSayisi := node.ID
		newAntCounts[node.ID] = yeniAntSayisi
	}
	graph.UpdateAntCounts(newAntCounts)

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
func suurballe(graph Graph) []Edge {
	// Grafı temsil eden matris
	dist := make([][]int, len(graph.Nodes))
	for i := range dist {
		dist[i] = make([]int, len(graph.Nodes))
	}

	// Grafın kenarlarını ağırlıklarla doldur
	for _, edge := range graph.Edges {
		dist[edge.Start][edge.End] = edge.Weight
		dist[edge.End][edge.Start] = edge.Weight
	}

	// Dijkstra algoritması kullanarak en kısa yolu bul
	shortestPath := dijkstra(graph, dist)

	return shortestPath
}

func dijkstra(graph Graph, dist [][]int) []Edge {
	startNode := graph.StartNodeID
	endNode := graph.EndNodeID
	n := len(graph.Nodes)

	// Tüm düğümlerin en kısa yol uzaklığını tutacak bir dizi oluştur
	shortest := make([]int, n)
	for i := range shortest {
		shortest[i] = math.MaxInt32
	}
	shortest[startNode] = 0

	// Ziyaret edilmemiş düğümleri tutacak bir dizi oluştur
	unvisited := make(map[int]bool)
	for _, node := range graph.Nodes {
		unvisited[node.ID] = true
	}

	var path []Edge

	for len(unvisited) > 0 {
		// En kısa yolun uzunluğunu bul
		minDist := math.MaxInt32
		minNode := -1
		for id := range unvisited {
			if shortest[id] < minDist {
				minDist = shortest[id]
				minNode = id
			}
		}

		// Ziyaret edilen düğümü kaldır
		delete(unvisited, minNode)

		// En kısa yolun düğümlerini güncelle
		for i := 0; i < n; i++ {
			if dist[minNode][i] > 0 && shortest[minNode]+dist[minNode][i] < shortest[i] {
				shortest[i] = shortest[minNode] + dist[minNode][i]
				path = append(path, Edge{Start: minNode, End: i, Weight: dist[minNode][i]})
			}
		}
	}

	// Hedef düğüme giden en kısa yolu döndür
	var shortestPath []Edge
	currentNode := endNode
	for currentNode != startNode {
		for _, edge := range path {
			if edge.End == currentNode {
				shortestPath = append([]Edge{edge}, shortestPath...)
				currentNode = edge.Start
				break
			}
		}
	}
	return shortestPath
}
