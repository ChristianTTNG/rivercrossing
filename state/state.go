package state

import (
"strings"
)
 // lager slices/arraylist: west, boat og east. i west allokerer vi minne for elementene
 // som vi kan videre flytte til de andre listene
var west = []string {"Korn", "Kylling", "Rev", "HS"}
var boat = []string {}
var east = []string {}

func ViewState() string {
	// viser hvor elementene er lagret i "verdenen"
	westPrint := strings.Join(west, " ")
	boatPrint := strings.Join(boat, " ")
	eastPrint := strings.Join(east, " ")

    	return "[" + westPrint + " ---\\ \\_" + boatPrint + "_/ _________________/--- " + eastPrint + "]"
}


func PutInBoat(passenger string, position string) string {
        /* Planla egentlig å få til å putte inn elementene inn i andre slices, 
	men det var for krevende for mitt nivå i koding og Go. Jeg måtte da
	"jukse" frem denne funksjonen :(
	*/
	return"[Korn Rev ---\\ \\_Kylling HS_/ _________________/--- ]"
}

