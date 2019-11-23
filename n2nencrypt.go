package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

//Storage - The object is accessed by index.
type Storage struct {
	Str  string `json:"storageString"`
	Code string `json:"storageCode"`
}

//Key -  The key is accessed by index.
type Key struct {
	KKey []int `json:"storageKey"`
}

func saveG(w Key, o Storage) {
	Scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Type in a name for the savefile (this will be saved in same folder as executable):")
	Scanner.Scan()
	savefile := Scanner.Text()
	convertkey := &w
	convertobject := &o
	output, err := json.Marshal(convertkey)
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

func loadG(w *Key, o *Storage) {
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
	key := *k
	By := []byte(storage.Str)
	Intby := []byte{}
	counter := 0
	for counter != 1 {
		if len(key.KKey) == len(By) {
			counter = 1
		}
		check(&key.KKey, len(By)+1)
	}
	for index := range By {
		for index2 := range key.KKey {
			Intby = append(Intby, byte(int(By[index])+key.KKey[index2]))
		}
	}
	storage.Code = string(Intby[:])
	storage.Str = "this has been encoded"
	fmt.Println(storage.Code)
	fmt.Println("")
	fmt.Println(key.KKey)
	*s = storage
	*k = key
}

func (s *Storage) Decode(x Key) {
	storage := *s
	By := []byte(storage.Code)
	decoded := []byte{}
	counter := 0
	for i := range By {
		counter = counter + 1
		if counter == len(x.KKey) {
			decoded = append(decoded, byte(int(By[i])-x.KKey[counter-1]))
			counter = 0
		}
	}
	answer := string(decoded[:])
	fmt.Println(answer)
}

func main() {
	Scanner := bufio.NewScanner(os.Stdin)
	storage := Storage{}
	key := Key{}
	gameover := 0
	fmt.Println("1 to 1 encryption")
	for gameover != 1 {
		fmt.Println("Type in command: s to save, l to load, t to type in string for encode, e to encode, d to decode and q to quit.")
		Scanner.Scan()
		result := Scanner.Text()
		switch result {
		case "s":
			saveG(key, storage)
		case "l":
			loadG(&key, &storage)
		case "t":
			fmt.Println("Type in your string:")
			Scanner.Scan()
			storage.Str = Scanner.Text()
			storage.Encode(&key)
		case "e":
			storage.Encode(&key)
		case "d":
			fmt.Println("Do you have key loaded already? type 'y' to decode with loaded key, 'n' to type in key manually.")
			Scanner.Scan()
			res := Scanner.Text()
			switch res {
			case "y":
				fmt.Println("Decoding:")
				fmt.Println("")
				storage.Decode(key)
				fmt.Println("")
				fmt.Println("Decoded!")
			case "n":
				key.KKey = nil
				fmt.Println("Enter key like this '1 2 3 4' etc:")
				Scanner.Scan()
				keyscan := Scanner.Text()
				stringarray := strings.Fields(keyscan)
				for i := range stringarray {
					no, err := strconv.Atoi(stringarray[i])
					if err != nil {
						fmt.Println("You typed in key incorrectly.")
					}
					key.KKey = append(key.KKey, no)
				}
				fmt.Println(key.KKey)
				fmt.Println("Decoding:")
				fmt.Println("")
				storage.Decode(key)
				fmt.Println("")
				fmt.Println("Decoded!")
			}

		case "q":
			gameover = 1
		}
	}
}
