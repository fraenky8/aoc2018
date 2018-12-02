package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	b, err := ioutil.ReadFile("../input.txt")
	if err != nil {
		log.Fatal(err)
	}

	bs := bytes.Split(b, []byte("\n"))

	_, _, common, err := common(bs)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("common letters: %v\n", common)
}

func common(bs [][]byte) (id1 []byte, id2 []byte, common string, err error) {

	found := false

	for i := 0; i < len(bs); i++ {
		for j := i + 1; j < len(bs)-1; j++ {
			id1 = bs[i]
			id2 = bs[j]
			differences := 0
			for k := range id1 {
				if id1[k] != id2[k] {
					differences++
				}
			}
			if differences == 1 {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if !found {
		return nil, nil, "", fmt.Errorf("no correct boxes found")
	}

	var pos int
	for i := range id1 {
		if id1[i] != id2[i] {
			pos = i
			break
		}
	}

	common = string(id1[0:pos]) + string(id1[pos+1:])

	return id1, id2, common, nil
}
