package anvil

import (
	"net/url"
	"regexp"
	"strconv"
	"strings"
)

func isRpcURL(s string) bool {
	u, err := url.ParseRequestURI(s)
	if err != nil {
		return false
	}
	return u.Scheme == "http" || u.Scheme == "https"
}

func isForkURL(s string) bool {
	if !isRpcURL(s) {
		return false
	}
	return strings.Contains(s, "infura") || strings.Contains(s, "alchemy")

}

func isChainID(s string) (bool, int) {
	n, err := strconv.Atoi(s)
	if err != nil {
		return false, 0
	}
	return true, n
}

func isGasLimit(s string) (bool, uint64) {
	gl, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return false, 0
	}
	return err == nil, gl
}

func isGasFee(s string) (bool, uint64) {
	gf, err := strconv.ParseUint(s, 10, 64)
	if err != nil {
		return false, 0
	}
	return err == nil, gf
}

func isPrivateKey(s string) bool {
	s = strings.TrimPrefix(s, "0x")
	match, _ := regexp.MatchString("^[0-9a-fA-F]{64}$", s)
	return match
}
