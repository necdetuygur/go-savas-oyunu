package main

import (
	"fmt"
	"math/rand"
	"savas/oyun"
	"strconv"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	takim := 0
	dunyalilar := []oyun.Dunyali{}
	marslilar := []oyun.Marsli{}

	fmt.Print("\n\n\nTakım Sayısı Girin: ")
	fmt.Scan(&takim)

	for i := 0; i < takim; i++ {
		index := strconv.Itoa(i + 1)

		dunyali := new(oyun.Dunyali)
		dunyali.Ad = "DünyalıSavaşçı" + index
		dunyali.Birlik = "Kütahya"
		dunyali.Can = 100
		dunyali.Mermi = 90
		dunyalilar = append(dunyalilar, *dunyali)

		marsli := new(oyun.Marsli)
		marsli.Ad = "MarslıSavaşçı" + index
		marsli.UzayGemisi = "KTHY950"
		marsli.Can = 100
		marsli.Mermi = 90
		marslilar = append(marslilar, *marsli)
	}

	for {
		tumDunyalilarinCaniBitti := true
		tumMarslilarinCaniBitti := true

		for i := 0; i < takim; i++ {
			for _, dunyali := range dunyalilar {
				if dunyali.Can > 0 {
					tumDunyalilarinCaniBitti = false
				}
			}
		}

		if tumDunyalilarinCaniBitti {
			fmt.Println("Marslılar kazandı!")
			break
		}

		for i := 0; i < takim; i++ {
			for _, marsli := range marslilar {
				if marsli.Can > 0 {
					tumMarslilarinCaniBitti = false
				}
			}
		}

		if tumMarslilarinCaniBitti {
			fmt.Println("Dünyalılar kazandı!")
			break
		}

		// TODO: BUG: SAVAŞÇININ CANI BİTMİŞ OLSA BİLE, ATEŞ EDEBİLİYOR.
		saldiriRandom := rand.Intn(2)
		randomSavasci := rand.Intn(takim)
		if saldiriRandom == 0 {
			dunyalilar[randomSavasci].AtesEt()
			marslilar[randomSavasci].HasarAl()
		} else {
			marslilar[randomSavasci].AtesEt()
			dunyalilar[randomSavasci].HasarAl()
		}
	}

	fmt.Printf("%+v\n", dunyalilar)
	fmt.Printf("%+v\n", marslilar)
}
