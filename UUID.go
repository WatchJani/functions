package main

import (
	"crypto/rand"
	"fmt"
)

//generise random listu bajtova od 0 do 255
//UUID je dug 128 bita odnosno 16 bajta
//poslije generisanja niza od 16bajta, formatira se u string pomocu hexadecimalnog zapisa
//da bi se zapisala jedna cifra u hexadecimalno zapisu treba 4bita
//samim tim nas krajnji string ID je dug 32 hexadecimalna karaktera a za svaki treba 4bita da se zapise
//samim tim se duzina zapisa je teska 128 bita

func generateUUIDv4() (string, error) {
	uuid := make([]byte, 16)
	_, err := rand.Read(uuid)
	if err != nil {
		return "", err
	}

	// Postavljanje verzije i varijante za UUID verzije 4
	uuid[6] = (uuid[6] & 0x0F) | 0x40 // Verzija 4
	uuid[8] = (uuid[8] & 0x3F) | 0x80 // Varijanta DCE 1.1

	// Formatiranje UUID u string
	uuidStr := fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
	return uuidStr, nil
}
