package mdt

type AKUxx1601 struct {
}

func (a *AKUxx1601) Functions() []string {
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
		"Verfahrstatus_1B",
		"Absolute_Position_1By",
		//"Absolute_Lamellenposition_1By",
		"Status_akt_Position_1By",
		//"Status_akt_Lamellenposition_1By",
		"Akt_Position_gültig_1B",
		"Referenzfahrt_starten_1B",
		"Position_anfahren_1B",
		"Status_obere_Position_1B",
		"Status_untere_Position_1B",
		"Sperren_absolute_Position_1B",
		"Sperren_universal_1B",
		"Windalarm_1B",
		"Regenalarm_1B",
		"Frostalarm_1B",
		"Sperren_1B",
	}
}

func (a *AKUxx1601) Prefix() string {
	return "U"
}