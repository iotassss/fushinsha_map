package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
)

type PersonModel struct {
	UUID         string
	Emoji        string
	Sign         string
	RegisterUUID string
	SightedCount int
	SightingTime string
	X            float64 // 経度
	Y            float64 // 緯度
	Gender       string
	Clothing     string
	Accessories  string
	Vehicle      string
	Behavior     string
	Hairstyle    string
}

const (
	minLat = 35.2
	maxLat = 37.5
	minLon = 138.5
	maxLon = 141.0
)

func randInRange(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}

func randEmoji() string {
	code := rand.Intn(0x1F64A-0x1F600+1) + 0x1F600
	return string(rune(code))
}

func main() {
	n := flag.Int("n", 10, "number of persons")
	flag.Parse()
	rand.Seed(time.Now().UnixNano())

	fmt.Println("dummyPersons := []PersonModel{")
	for i := 0; i < *n; i++ {
		lon := randInRange(minLon, maxLon)
		lat := randInRange(minLat, maxLat)
		emoji := randEmoji()
		uuid := fmt.Sprintf("%08x-%04x-%04x-%04x-%012x", rand.Uint32(), rand.Uint32()&0xffff, rand.Uint32()&0xffff, rand.Uint32()&0xffff, rand.Uint64()&0xffffffffffff)
		fmt.Printf("\t{ UUID: %q, Emoji: %q, Sign: %q, RegisterUUID: %q, SightedCount: %d, SightingTime: %q, X: %.4f, Y: %.4f, Gender: %q, Clothing: %q, Accessories: %q, Vehicle: %q, Behavior: %q, Hairstyle: %q },\n",
			uuid,
			emoji,
			"Z",
			"a1a2b3c4-d5e6-7f89-0abc-def123456789",
			2,
			"13:30",
			lon,
			lat,
			"女性",
			"制服",
			"バッグ",
			"自動車",
			"暴力",
			"パーマ",
		)
	}
	fmt.Println("}")
}
