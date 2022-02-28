package state

import "testing"

func TestViewState(t *testing.T) {
    wanted := "[Korn Kylling Rev HS ---\\ \\__/ _________________/--- ]"
    state := ViewState();
    if state != wanted {
        t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}

func TestPutInBoat(t *testing.T) {
	wanted := "[Korn Rev ---\\ \\_Kylling HS_/ _________________/--- ]"
	state := PutInBoat("west", "Kylling");
	if state != wanted {
	t.Errorf("Feil, fikk %q, ønsket %q.", state, wanted)
	}
}
