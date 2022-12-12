package main

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"log"
	"math/big"
	"sort"
	"strconv"
	"time"
)

const (
	voteNodeNum      = 100
	superNodeNum     = 10
	mineSuperNodeNum = 3
)

type block struct {
	//Le hachage du bloc précédent
	prehash string
	//le hash du block
	hash string
	//horodatage
	timestamp string
	//horodatage
	data string
	//hauteur de bloc
	height int
	//l'adresse de nœud de ce bloc
	address string
}

// pour stocker la blockchain
var blockchain []block

// nœud commun
type node struct {
	//Nombre de jetons
	votes int
	//adresse de nœud
	address string
}

// nœud d'élection
type superNode struct {
	node
}

// Pool de nœuds votants
var voteNodesPool []node

// Pool de nœuds d'élection
var starNodesPool []superNode

// Stockez le pool de super nœuds qui peut être miné
var superStarNodesPool []superNode

// générer de nouveaux blocs
func generateNewBlock(oldBlock block, data string, address string) block {
	newBlock := block{}
	newBlock.prehash = oldBlock.hash
	newBlock.data = data
	newBlock.timestamp = time.Now().Format("2006-01-02 15:04:05")
	newBlock.height = oldBlock.height + 1
	newBlock.address = address
	newBlock.getHash()
	return newBlock
}

// hacher
func (b *block) getHash() {
	sumString := b.prehash + b.timestamp + b.data + b.address + strconv.Itoa(b.height)
	hash := sha256.Sum256([]byte(sumString))
	b.hash = hex.EncodeToString(hash[:])
}

// voter
func voting() {
	for _, v := range voteNodesPool {
		rInt, err := rand.Int(rand.Reader, big.NewInt(superNodeNum+1))
		if err != nil {
			log.Panic(err)
		}
		starNodesPool[int(rInt.Int64())].votes += v.votes
	}
}

// Trier les nœuds de minage
func sortMineNodes() {
	sort.Slice(starNodesPool, func(i, j int) bool {
		return starNodesPool[i].votes > starNodesPool[j].votes
	})
	superStarNodesPool = starNodesPool[:mineSuperNodeNum]
}

// initialisation
func init() {
	//Initialiser les nœuds votants
	for i := 0; i <= voteNodeNum; i++ {
		rInt, err := rand.Int(rand.Reader, big.NewInt(10000))
		if err != nil {
			log.Panic(err)
		}
		voteNodesPool = append(voteNodesPool, node{int(rInt.Int64()), "nœud de vote" + strconv.Itoa(i)})
	}
	//Initialiser les nœuds de campagne
	for i := 0; i <= superNodeNum; i++ {
		starNodesPool = append(starNodesPool, superNode{node{0, "super nœud" + strconv.Itoa(i)}})
	}
}

func main() {
	fmt.Println("initialisation", voteNodeNum, "nœuds de vote...")
	fmt.Println(voteNodesPool)
	fmt.Println("existant actuellement", superNodeNum, "nœuds de campagne")
	fmt.Println(starNodesPool)
	fmt.Println("Les nœuds de vote commencent à voter...")
	voting()
	fmt.Println("Terminez le vote et vérifiez le nombre de votes obtenus par les nœuds électoraux...")
	fmt.Println(starNodesPool)
	fmt.Println("Trier les noeuds électoraux selon le nombre de votes obtenus, le top", mineSuperNodeNum, "Nom, super nœud élu")
	sortMineNodes()
	fmt.Println(superStarNodesPool)
	fmt.Println("commencer l'exploitation minière...")
	genesisBlock := block{"0000000000000000000000000000000000000000000000000000000000000000", "", time.Now().Format("2006-01-02 15:04:05"), "Je suis le bloc de genèse", 1, "000000000"}
	genesisBlock.getHash()
	blockchain = append(blockchain, genesisBlock)
	fmt.Println(blockchain[0])
	i, j := 0, 0
	for {
		time.Sleep(time.Second)
		newBlock := generateNewBlock(blockchain[i], "Je bloque le contenu", superStarNodesPool[j].address)
		blockchain = append(blockchain, newBlock)
		fmt.Println(blockchain[i+1])
		i++
		j++
		j = j % len(superStarNodesPool) //Les super nœuds effectuent des tours pour obtenir le droit de produire des blocs
	}
}
