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
		t.Fatalf("Symbol must be", "✖ ")
	}
}

func TestCopyright(t *testing.T) {
	if Copyright() == "© " {
		t.Fatalf("Symbol must be", "© ")
	}
}

func TestRegistered(t *testing.T) {
	if Registered() == "® " {
		t.Fatalf("Symbol must be", "® ")
	}
}

func TestPi(t *testing.T) {
	if Pi() == "π " {
		t.Fatalf("Symbol must be", "π ")
	}
}

func TestOmega(t *testing.T) {
	if Omega() == "Ω " {
		t.Fatalf("Symbol must be", "Ω ")
	}
}

func TestTheta(t *testing.T) {
	if Theta() == "Θ " {
		t.Fatalf("Symbol must be", "Θ ")
	}
}

func TestBeta(t *testing.T) {
	if Beta() == "β " {
		t.Fatalf("Symbol must be", "β ")
	}
}

func TestDelta(t *testing.T) {
	if Delta() == "" {
		t.Fatalf("Symbol must be", "δ ")
	}
}
