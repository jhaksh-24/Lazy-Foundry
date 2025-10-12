package anvil

import (
	"testing"
)

func TestisRpcURL(t *testing.T) {
	want := true
	gotru1 := isRpcURL("https://mainnet.infura.io/v3/1234567890abcdef")

	if want != gotru1 {
		t.Errorf("Check for isRpcURL failed")
	}

	gotru2 := isRpcURL("https://my-custom-node.example.com/rpc")

	if want != gotru2 {
		t.Errorf("Check for isRpcURL failed")
	}

	gotru3 := isRpcURL("https://rpc.ankr.com/eth")

	if want != gotru3 {
		t.Errorf("Check for isRpcURL failed")
	}

	gotru4 := isRpcURL("http://localhost:8545")

	if want != gotru4 {
		t.Errorf("Check for isRpcURL failed")
	}

	gotru5 := isRpcURL("http://127.0.0.1:7545")

	if want != gotru5 {
		t.Errorf("Check for isRpcURL failed")
	}

	gotru6 != isRpcURL("ftp://rpc.example.com")

	if want != gotru6 {
		t.Errorf("Check for isRpcURL failed... This one should fail")
	}

	gotru7 != isRpcURL("mainnet.infura.io/v3/123456")

	if want != gotru7 {
		t.Errorf("Check for isRpcURL failed... This one should fail")
	}

	gotru8 != isRpcURL("://invalid-url")

	if want != gotru8 {
		t.Errorf("Check for isRpcURL failed... This one should fail")
	}

	gotru9 != isRpcURL("http//missing-colon.com")

	if want != gotru9 {
		t.Errorf("Check for isRpcURL failed... This one should fail")
	}

	gotru10 != isRpcURL("randomstring")

	if want != gotru10 {
		t.Errorf("Check for isRpcURL failed... This one should fail")
	}
}
