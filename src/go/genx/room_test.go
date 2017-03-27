package genx_test

import (
	"bytes"
	"fmt"
	"testing"

	"go/genx"
)

func TestLightsMDT(t *testing.T) {
	r := genx.NewRoom("R.00.01", 2, 0, 0, 0, 0)
	var b bytes.Buffer
	r.Generate(&b, genx.MDT)
	fmt.Printf(b.String())
	// Output:
	// L_R.00.01_01_Schalten_1B 1/0/1
	// L_R.00.01_01_Treppenlicht_1B 1/0/2
	// L_R.00.01_01_Service erforderlich_1B 1/0/3
	// L_R.00.01_01_Schaltimpuls_1B 1/0/4
	// L_R.00.01_01_Sperren_1B 1/0/5
	// L_R.00.01_01_Zeit_bis_nächsten_Service_2By_4By 1/0/6
	// L_R.00.01_01_Betriebsstunden_RM_2By_4By 1/0/7
	// L_R.00.01_01_Treppenlicht_mit_Zeit_1By 1/0/8
	// L_R.00.01_01_Vorwarnen_1B 1/0/9
	// L_R.00.01_01_Betriebsstunden_Reset_1B 1/0/10
	// L_R.00.01_01_Servicemeldung_Reset_1B 1/0/11
	// L_R.00.01_01_Zwangsführung_2B 1/0/12
	// L_R.00.01_01_Priorität_1B 1/0/13
	// L_R.00.01_01_Szene_1By 1/0/14
	// L_R.00.01_01_Status_1B 1/0/15
	// L_R.00.01_01_Status_invertiert_1B 1/0/16
	// L_R.00.01_01_Schwellwertschalter_1By_2By 1/0/17
	// L_R.00.01_01_Logik_1_1B 1/0/18
	// L_R.00.01_01_Logik_2_1B 1/0/19
	// L_R.00.01_02_Schalten_1B 1/0/1
	// L_R.00.01_02_Treppenlicht_1B 1/0/2
	// L_R.00.01_02_Service erforderlich_1B 1/0/3
	// L_R.00.01_02_Schaltimpuls_1B 1/0/4
	// L_R.00.01_02_Sperren_1B 1/0/5
	// L_R.00.01_02_Zeit_bis_nächsten_Service_2By_4By 1/0/6
	// L_R.00.01_02_Betriebsstunden_RM_2By_4By 1/0/7
	// L_R.00.01_02_Treppenlicht_mit_Zeit_1By 1/0/8
	// L_R.00.01_02_Vorwarnen_1B 1/0/9
	// L_R.00.01_02_Betriebsstunden_Reset_1B 1/0/10
	// L_R.00.01_02_Servicemeldung_Reset_1B 1/0/11
	// L_R.00.01_02_Zwangsführung_2B 1/0/12
	// L_R.00.01_02_Priorität_1B 1/0/13
	// L_R.00.01_02_Szene_1By 1/0/14
	// L_R.00.01_02_Status_1B 1/0/15
	// L_R.00.01_02_Status_invertiert_1B 1/0/16
	// L_R.00.01_02_Schwellwertschalter_1By_2By 1/0/17
	// L_R.00.01_02_Logik_1_1B 1/0/18
	// L_R.00.01_02_Logik_2_1B 1/0/19
}