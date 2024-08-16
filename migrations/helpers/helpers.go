package helpers

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func NormalizeString(s *string) string {
	s1 := strings.ReplaceAll(*s, "'", "’")
	s2 := strings.TrimSpace(s1)
	return s2
}

func ReadXMLData(path string) (*Library, error) {
	var Akteure Akteure
	var Reihentitel Reihentitel
	var Bände Bände
	var Inhalte Inhalte
	var Orte Orte
	var Relationen_Bände_Akteure Relationen_Bände_Akteure
	var Relationen_Bände_Reihen Relationen_Bände_Reihen
	var Relationen_Inhalte_Akteure Relationen_Inhalte_Akteure

	if err := UnmarshalFile(path+"Akteure.xml", &Akteure); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"Orte.xml", &Orte); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"Reihen.xml", &Reihentitel); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"Baende.xml", &Bände); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"Inhalte.xml", &Inhalte); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"_RELATION_BaendeAkteure.xml", &Relationen_Bände_Akteure); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"_RELATION_BaendeReihen.xml", &Relationen_Bände_Reihen); err != nil {
		return nil, err
	}

	if err := UnmarshalFile(path+"_RELATION_InhalteAkteure.xml", &Relationen_Inhalte_Akteure); err != nil {
		return nil, err
	}

	lib := Library{
		Orte:                       &Orte,
		Akteure:                    &Akteure,
		Reihentitel:                &Reihentitel,
		Bände:                      &Bände,
		Inhalte:                    &Inhalte,
		Relationen_Bände_Akteure:   &Relationen_Bände_Akteure,
		Relationen_Bände_Reihen:    &Relationen_Bände_Reihen,
		Relationen_Inhalte_Akteure: &Relationen_Inhalte_Akteure,
	}

	return &lib, nil
}

func UnmarshalFile[T any](filename string, data *T) error {
	xmlFile, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return err
	}
	fmt.Println("Successfully opened " + filename)
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, data)

	return nil
}

func MakeMap[T any, U comparable](data []T, f func(T) U) map[U]T {
	m := make(map[U]T)
	for _, v := range data {
		m[f(v)] = v
	}
	return m
}
