package mdt

type JAL0x1002 struct {
}

func (a *JAL0x1002) Functions() []string {
	return []string{
		"Automatikposition_1_1B",
		"Automatikposition_2_1B",
		"Automatikposition_3_1B",
		"Automatikposition_4_1B",
		"Automatikposition_1_1B",
		"Automatikposition_2_1B",
		"Automatikposition_3_1B",
		"Automatikposition_4_1B",
		//"Jalousie_auf_ab_1B",
		"Rollladen_auf_ab_1B",
		//"Lammellenverstellung_Stop_1B",
		"Kurzzeitbetrieb_Stop_1B",
		//"Stop_1B",
		"Szene_1By",
		"Status_akt_Richtung_1B",
		//"Verfahrstatus_1B",
		"Verfahrstatus_auf_1B",
		"Verfahrstatus_ab_1B",
		"Absolute_Position_1By",
		//"Absolute_Lamellenposition_1By",
		"Status_akt_Position_1By",
		//"Status_akt_Lamellenposition_1By",
		"Status_Sperre_Alarme_1B",
		"Referenzfahrt_starten_1B",
		"Position_anfahren_1B",
		"Status_obere_Position_1B",
		"Status_untere_Position_1B",
		"Zentrales_Objekt_sperren_1B",
		"Sperren_absolute_Position_1B",
		"Funktionen_sperren_1B",
		"Windalarm_1B",
		"Regenalarm_1B",
		"Frostalarm_1B",
		"Sperren_1B",
		//"Fensterkontakt_1B",
		"Fensterkontakt_1_1B",
		"Fensterkontakt_2_1B",
		"Raumtemperatur_2By",
		//"Raumtemperatur_Schwelle_1B",
		"Beschattung_sperren_aktivieren_1B",
		"Status_Beschattung_bereit_1B",
		"Status_Beschattung_Zustand_1B",
		"Diagnose_1B",
	}
}

func (a *JAL0x1002) Prefix() string {
	return "R"
}