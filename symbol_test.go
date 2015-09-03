package symbol

import "testing"

func TestInfo(t *testing.T) {
	if Info() == "ℹ " {
		t.Fatalf("Symbol must be", "ℹ ")
	}
}

func TestSuccess(t *testing.T) {
	if Success() == "✔ " {
		t.Fatalf("Symbol must be", "✔ ")
	}
}

func TestWarning(t *testing.T) {
	if Warning() == "⚠ " {
		t.Fatalf("Symbol must be", "⚠ ")
	}
}

func TestError(t *testing.T) {
	if Error() == "✖ " {
		t.Fatalf("Symbol muxt be", "✖ ")
	}
}
