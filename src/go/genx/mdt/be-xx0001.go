package mdt

type BExx0001 struct {
}

func (a *BExx0001) Functions() []string {
	return []string{
		"Schalter_1B",
		//"Jalousie_1B",
		//"Wert_senden_1By",
		//"Zwangsführung_senden_2B",
		//"Zähler_zurücksetzen_1B",
		//"Dimmen_ein_aus_1B",
		//"kurze_Taste_1B",
		//"kurze_Taste_1By",
		"Wert_für_Umschaltung_1B",
		//"Stop_Lamellen_1B",
		//"Dimmen_4B",
		"Szene_1By",
		//"Wert_für_Umschaltung_1B",
		//"Wert_für_Richtungswechsel_1B",
		//"lange_Taste_1B",
		//"lange_Taste_1By",
		"Zähler_4By",
		"Sperrobjekt_1B",
		"Eingangslogik_1A_1B",
		"Eingangslogik_1B_1B",
		"Ausgangslogik_1_1B",
		"Ausgangslogik_1_Szene_1By",
	}
}

func (a *BExx0001) Prefix() string {
	return "FK"
}
