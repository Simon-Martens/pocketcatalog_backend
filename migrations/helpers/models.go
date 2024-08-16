package helpers

import "encoding/xml"

type Library struct {
	Orte                       *Orte
	Akteure                    *Akteure
	Reihentitel                *Reihentitel
	Bände                      *Bände
	Inhalte                    *Inhalte
	Relationen_Bände_Akteure   *Relationen_Bände_Akteure
	Relationen_Inhalte_Akteure *Relationen_Inhalte_Akteure
	Relationen_Bände_Reihen    *Relationen_Bände_Reihen
}

type Akteure struct {
	XMLName xml.Name `xml:"dataroot"`
	Akteure []Akteur `xml:"Akteure"`
}

type Akteur struct {
	ID           string `xml:"ID"`
	Name         string `xml:"NAME"`
	Körperschaft bool   `xml:"ORGANISATION"`
	Beruf        string `xml:"BERUF"`
	Nachweis     string `xml:"NACHWEIS"`
	Pseudonyme   string `xml:"PSEUDONYM"`
	Lebensdaten  string `xml:"LEBENSDATEN"`
	Anmerkungen  string `xml:"ANMERKUNGEN"`
	GND          string `xml:"GND"`
}

type Reihentitel struct {
	XMLName xml.Name `xml:"dataroot"`
	Reihen  []Reihe  `xml:"Reihen"`
}

type Reihe struct {
	ID          string `xml:"ID"`
	Titel       string `xml:"NAME"`
	Sortiername string `xml:"SORTIERNAME"`
	Nachweis    string `xml:"NACHWEIS"`
	Anmerkungen string `xml:"Anmerkungen"`
}

type Bände struct {
	XMLName xml.Name `xml:"dataroot"`
	Bände   []Band   `xml:"Baende"`
}

type Orte struct {
	XMLName xml.Name `xml:"dataroot"`
	Orte    []Ort    `xml:"Orte"`
}

type Ort struct {
	ID          string `xml:"ID"`
	Name        string `xml:"NAME"`
	Fiktiv      bool   `xml:"FIKTIV"`
	Anmerkungen string `xml:"Anmerkungen"`
}

type Band struct {
	ID                        string       `xml:"ID"`
	BiblioID                  int          `xml:"BIBLIO-ID"`
	Titelangabe               string       `xml:"TITEL"`
	Ortsangabe                string       `xml:"ORT-ALT"`
	Orte                      []Ortverweis `xml:"ORTE"`
	Verantwortlichkeitsangabe string       `xml:"HERAUSGEBER"`
	Jahr                      int          `xml:"JAHR"`
	Gesichtet                 bool         `xml:"AUTOPSIE"`
	Erfasst                   bool         `xml:"ERFASST"`
	Nachweis                  string       `xml:"NACHWEIS"`
	Struktur                  string       `xml:"STRUKTUR"`
	Norm                      string       `xml:"NORM"`
	Status                    Status       `xml:"STATUS"`
	Anmerkungen               string       `xml:"ANMERKUNGEN"`
	ReihentitelALT            string       `xml:"REIHENTITEL-ALT"`
}

type Ortverweis struct {
	Value string `xml:"Value"`
}

type Status struct {
	Value []string `xml:"Value"`
}

type Inhalte struct {
	XMLName xml.Name `xml:"dataroot"`
	Inhalte []Inhalt `xml:"Inhalte"`
}

type Inhalt struct {
	ID            string `xml:"ID"`
	Titelangabe   string `xml:"TITEL"`
	Urheberangabe string `xml:"AUTOR"`
	Band          string `xml:"BAND"`
	Objektnummer  string `xml:"OBJEKTNUMMER"`
	Incipit       string `xml:"INCIPIT"`
	Paginierung   string `xml:"PAGINIERUNG"`
	Typ           Typ    `xml:"TYP"`
	Anmerkungen   string `xml:"ANMERKUNGEN"`
	Seite         string `xml:"SEITE"`
}

type Typ struct {
	Value []string `xml:"Value"`
}

type Relationen_Bände_Reihen struct {
	XMLName    xml.Name              `xml:"dataroot"`
	Relationen []Relation_Band_Reihe `xml:"_x002A_RELATION_BaendeReihen"`
}

type Relation_Band_Reihe struct {
	ID       string `xml:"ID"`
	Band     string `xml:"BAND"`
	Relation string `xml:"BEZIEHUNG"`
	Reihe    string `xml:"REIHE"`
}

type Relationen_Inhalte_Akteure struct {
	XMLName    xml.Name                 `xml:"dataroot"`
	Relationen []Relation_Inhalt_akteur `xml:"_x002A_RELATION_InhalteAkteure"`
}

type Relation_Inhalt_akteur struct {
	ID       string `xml:"ID"`
	Band     string `xml:"INHALT"`
	Relation string `xml:"BEZIEHUNG"`
	Akteur   string `xml:"AKTEUR"`
}

type Relationen_Bände_Akteure struct {
	XMLName    xml.Name               `xml:"dataroot"`
	Relationen []Relation_Band_Akteur `xml:"_x002A_RELATION_BaendeAkteure"`
}

type Relation_Band_Akteur struct {
	ID       string `xml:"ID"`
	Band     string `xml:"BAND"`
	Relation string `xml:"BEZIEHUNG"`
	Akteur   string `xml:"AKTEUR"`
}
