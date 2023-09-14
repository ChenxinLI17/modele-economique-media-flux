package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fmt"
)

func Contains[T comparable](s []T, e T) bool {
	for _, v := range s {
		if v == e {
			return true
		}
	}
	return false
}

func RemoveVideo(list []Video, index int) []Video {
	// Check l'indice
	if index < 0 || index >= len(list) {
		return list
	}

	// Retourne une slice en excluant l'indice
	return append(list[:index], list[index+1:]...)
}

func FormatFloat(f float32) string {
	return fmt.Sprintf("%.2f", f)
}

func GenerateUniqueID() string {
	// Générer 16 octets aléatoires
	bytes := make([]byte, 16)
	_, err := rand.Read(bytes)
	if err != nil {
		panic(err)
	}

	// Retourner les octets en tant qu'une chaîne hexadécimale
	return hex.EncodeToString(bytes)
}

func CompareFormats(prefs []Format, subFormats []Format) bool{
	for _,format:=range(subFormats){
		if Contains(prefs,format){
			return true
		}
	}
	return false
}