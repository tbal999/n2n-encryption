package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
)

type Storage struct {
	str   string
	by    []byte
	key   []int
	intby []byte
	code  string
}

type Key struct {
	key []int
}

func randomNumber(min, max int) int {
	z := rand.Intn(max)
	if z < min {
		z = min
	}
	return z
}

func check(n *[]int, y int) {
	nest := *n
	insertion := randomNumber(1, y)
	for i := range nest {
		if nest[i] == insertion {
			return
		}
	}
	nest = append(nest, insertion)
	*n = nest
}

func (s *Storage) Encode(k *Key) {
	storage := *s
	key := k
	storage.by = []byte(storage.str)
	counter := 0
	for counter != 1 {
		if len(key.key) == len(storage.by) {
			counter = 1
		}
		check(&key.key, len(storage.by)+1)
	}
	for index := range storage.by {
		for index2 := range key.key {
			storage.intby = append(storage.intby, byte(int(storage.by[index])+key.key[index2]))
		}
	}
	storage.code = string(storage.intby[:])
	fmt.Println(storage.code)
	fmt.Println("")
	fmt.Println(key.key)
	*s = storage
}

func (s *Storage) Decode(x Key) {
	storage := *s
	decoded := []byte{}
	counter := 0
	for i := range storage.intby {
		counter = counter + 1
		if counter == len(x.key) {
			decoded = append(decoded, byte(int(storage.intby[i])-x.key[counter-1]))
			counter = 0
		}
	}
	answer := string(decoded[:])
	fmt.Println(answer)
}

func saveGame(w Key, o Storage) {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in a name for the savefile (this will be saved in same folder as executable):")
	Scanner.Scan()
	savefile := Scanner.Text()
	convertmap := &w
	convertobject := &o
	output, err := json.Marshal(convertmap)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile("key"+savefile+".json", output, 0755)
	output2, err := json.Marshal(convertobject)
	if err != nil {
		fmt.Println(err)
		return
	}
	_ = ioutil.WriteFile("string"+savefile+".json", output2, 0755)
	fmt.Println("Saved " + savefile + "!")
}

func loadGame(w *Key, o *Storage) {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in name of savefile you wish to load (has to be in same folder as executable):")
	Scanner.Scan()
	savefile := Scanner.Text()
	worldmap := *w
	jsonFile, _ := ioutil.ReadFile("key" + savefile + ".json")
	_ = json.Unmarshal([]byte(jsonFile), &worldmap)
	*w = worldmap
	objectstorage := *o
	jsonFile2, _ := ioutil.ReadFile("string" + savefile + ".json")
	_ = json.Unmarshal([]byte(jsonFile2), &objectstorage)
	*o = objectstorage
	fmt.Println("Loaded " + savefile + "!")
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	storage := Storage{}
	key := Key{}
	gameover := 0
	for gameover != 1 {
		fmt.Println("1 to 1 encryption")
		fmt.Println("Type in command: s to save, l to load, t to type in string for encode, e to encode, d to decode and q to quit.")
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "s":
			saveGame(key, storage)
		case "l":
			loadGame(&key, &storage)
		case "t":
			Scanner.Scan()
			storage.str = Scanner.Text()
			storage.Encode(&key)
		case "e":
			storage.Encode(&key)
		case "d":
			storage.Decode(key)
		case "q":
			gameover = 1
		}
	}
}
