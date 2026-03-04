package main

import "time"

type Ansprechpartner struct {
	Mandant     int       `db:"Mandant" json:"Mandant"`
	AdrNr       int       `db:"AdrNr" json:"AdrNr"`
	LaufNr      int       `db:"LaufNr" json:"LaufNr"`
	AdrNrVerk   int       `db:"AdrNr_Verk" json:"AdrNr_Verk"`
	DomizilVerk int       `db:"Domizil_Verk" json:"Domizil_Verk"`
	Anrede      string    `db:"Anrede" json:"Anrede"`
	Name        string    `db:"Name" json:"Name"`
	Vorname     string    `db:"Vorname" json:"Vorname"`
	Kuerzel     *string   `db:"Kuerzel" json:"Kuerzel"`
	Telefon     *string   `db:"Telefon" json:"Telefon"`
	Funktion    *string   `db:"Funktion" json:"Funktion"`
	Titel       *string   `db:"Titel" json:"Titel"`
	Info        *string   `db:"Info" json:"Info"`
	Briefanrede string    `db:"Briefanrede" json:"Briefanrede"`
	INS_DATE    time.Time `db:"INS_DATE" json:"INS_DATE"`
	UPD_DATE    time.Time `db:"UPD_DATE" json:"UPD_DATE"`
	INS_SABE    string    `db:"INS_SABE" json:"INS_SABE"`
	UPD_SABE    string    `db:"UPD_SABE" json:"UPD_SABE"`
	INS_SERVER  string    `db:"INS_SERVER" json:"INS_SERVER"`
}

type Domizil struct {
	Mandant        int       `db:"Mandant" json:"Mandant"`
	AdrNr          int       `db:"AdrNr" json:"AdrNr"`
	LaufNr         int       `db:"LaufNr" json:"LaufNr"`
	Sort           int       `db:"Sort" json:"Sort"`
	Art            int       `db:"Art" json:"Art"`
	ArtBezeichnung string    `db:"ArtBezeichnung" json:"ArtBezeichnung"`
	LandCode       string    `db:"LandCode" json:"LandCode"`
	Land           string    `db:"Land" json:"Land"`
	PLZ            string    `db:"PLZ" json:"PLZ"`
	Ort            string    `db:"Ort" json:"Ort"`
	Adresse        *string   `db:"Adresse" json:"Adresse"`
	TelFax         *string   `db:"TelFax" json:"TelFax"`
	Notizen        *string   `db:"Notizen" json:"Notizen"`
	Briefanrede    *string   `db:"Briefanrede" json:"Briefanrede"`
	Darstellung    string    `db:"Darstellung" json:"Darstellung"`
	INS_DATE       time.Time `db:"INS_DATE" json:"INS_DATE"`
	UPD_DATE       time.Time `db:"UPD_DATE" json:"UPD_DATE"`
	INS_SABE       string    `db:"INS_SABE" json:"INS_SABE"`
	UPD_SABE       string    `db:"UPD_SABE" json:"UPD_SABE"`
	INS_SERVER     string    `db:"INS_SERVER" json:"INS_SERVER"`
}

type Fall struct {
	Mandant                 int       `db:"Mandant" json:"Mandant"`
	FallNo                  int       `db:"FallNo" json:"FallNo"`
	FallNo_TopParent        int       `db:"FallNo_TopParent" json:"FallNo_TopParent"`
	FallNo_Parent           int       `db:"FallNo_Parent" json:"FallNo_Parent"`
	Fall_Level              int       `db:"Fall_Level" json:"Fall_Level"`
	Owner                   string    `db:"Owner" json:"Owner"`
	FallBezeichnung         *string   `db:"FallBezeichnung" json:"FallBezeichnung"`
	FallInfo                *string   `db:"FallInfo" json:"FallInfo"`
	Ablage                  *string   `db:"Ablage" json:"Ablage"`
	Name                    string    `db:"Name" json:"Name"`
	HSatz                   float64   `db:"HSatz" json:"HSatz"`
	FKPreis                 float64   `db:"FKPreis" json:"FKPreis"`
	KmPreis                 float64   `db:"KmPreis" json:"KmPreis"`
	Interessewert           float64   `db:"Interessewert" json:"Interessewert"`
	Status                  string    `db:"Status" json:"Status"`
	Sprache                 int       `db:"Sprache" json:"Sprache"`
	TypRechnung             int       `db:"TypRechnung" json:"TypRechnung"`
	Waehrung                string    `db:"Waehrung" json:"Waehrung"`
	Intern                  int       `db:"Intern" json:"Intern"`
	Fibu                    int       `db:"Fibu" json:"Fibu"`
	Formular                int       `db:"Formular" json:"Formular"`
	MWStCode                int       `db:"MWStCode" json:"MWStCode"`
	Gesperrt                int       `db:"Gesperrt" json:"Gesperrt"`
	Abgerechnet             int       `db:"Abgerechnet" json:"Abgerechnet"`
	FixHonorar              int       `db:"FixHonorar" json:"FixHonorar"`
	FallBeginn              time.Time `db:"FallBeginn" json:"FallBeginn"`
	Standort                int       `db:"Standort" json:"Standort"`
	SRechFaktInt            int       `db:"SRechFaktInt" json:"SRechFaktInt"`
	Budget                  float64   `db:"Budget" json:"Budget"`
	ProjectMode             int       `db:"ProjectMode" json:"ProjectMode"`
	FallNo_Alt              *string   `db:"FallNo_Alt" json:"FallNo_Alt"`
	INS_DATE                time.Time `db:"INS_DATE" json:"INS_DATE"`
	UPD_DATE                time.Time `db:"UPD_DATE" json:"UPD_DATE"`
	INS_SABE                string    `db:"INS_SABE" json:"INS_SABE"`
	UPD_SABE                string    `db:"UPD_SABE" json:"UPD_SABE"`
	INS_SERVER              string    `db:"INS_SERVER" json:"INS_SERVER"`
	SubprojectsInheritOwner int       `db:"SubprojectsInheritOwner" json:"SubprojectsInheritOwner"`
	EBilling                *string   `db:"eBilling" json:"eBilling"`
}

type Subjekt struct {
	Mandant     int       `db:"Mandant" json:"Mandant"`
	AdrNr       int       `db:"AdrNr" json:"AdrNr"`
	Status      string    `db:"Status" json:"Status"`
	Art         string    `db:"Art" json:"Art"`
	Anrede      *string   `db:"Anrede" json:"Anrede"`
	Titel       *string   `db:"Titel" json:"Titel"`
	Funktion    *string   `db:"Funktion" json:"Funktion"`
	Name1       string    `db:"Name1" json:"Name1"`
	Name2       *string   `db:"Name2" json:"Name2"`
	Sprache     int       `db:"Sprache" json:"Sprache"`
	Geschlecht  int       `db:"Geschlecht" json:"Geschlecht"`
	Briefanrede string    `db:"Briefanrede" json:"Briefanrede"`
	Notizen     *string   `db:"Notizen" json:"Notizen"`
	INS_DATE    time.Time `db:"INS_DATE" json:"INS_DATE"`
	UPD_DATE    time.Time `db:"UPD_DATE" json:"UPD_DATE"`
	INS_SABE    string    `db:"INS_SABE" json:"INS_SABE"`
	UPD_SABE    string    `db:"UPD_SABE" json:"UPD_SABE"`
	INS_SERVER  string    `db:"INS_SERVER" json:"INS_SERVER"`
}

type SubjektStruct struct {
	Mandant1     int     `db:"Mandant_1" json:"Mandant_1"`
	AdrNr1       int     `db:"AdrNr_1" json:"AdrNr_1"`
	Art          int     `db:"Art" json:"Art"`
	Code         int     `db:"Code" json:"Code"`
	Mandant2     int     `db:"Mandant_2" json:"Mandant_2"`
	Nr2          int     `db:"Nr_2" json:"Nr_2"`
	Nr3          int     `db:"Nr_3" json:"Nr_3"`
	Nr4          int     `db:"Nr_4" json:"Nr_4"`
	Beschreibung *string `db:"Beschreibung" json:"Beschreibung"`
	INS_SERVER   string  `db:"INS_SERVER" json:"INS_SERVER"`
}
