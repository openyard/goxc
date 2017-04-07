package mdt

type AKH0x0002 struct {
}

func (a *AKH0x0002) Functions() []string {
	return []string{
		//"Stellwert_1B",
		//"Stellwert_1By",
		"Temperaturmesswert_2By",
		"Vorlauftemperatur_2By",
		"Sperren_1B",
		"Status_Stellwert_1B",
		"Status_Stellwert_1By",
		"Komfortverlängerung_1B",
		"Zwangsstellung_1B",
		"Taupunktalarm_1B",
		"PWM_Ausgang_Kühlen_für_4_Rohr_1By",
		"Sollwert_Komfort_2By",
		"Sollwertverschiebung_2By",
		"Aktueller_Sollwert_2By",
		"Betriebsartvorwahl_1By",
		"DPT_HVAC_Status_1By",
		"DPT_RHCC_Status_2By",
		"Betriebsart_Komfort_1B",
		"Betriebsart_Nacht_1B",
		"Betriebsart_Frost_Hitzeschutz_1B",
		"Frostalarm_1B",
		"Hitzealarm_1B",
		"Sollwertverschiebung_1B",
		"Diagnosetext_14By",
	}
}

func (a *AKH0x0002) Prefix() string {
	return "H"
}