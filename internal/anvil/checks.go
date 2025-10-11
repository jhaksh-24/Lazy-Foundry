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

func isChainId(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}

func isGasLimit(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

func isGasPrice(s string) bool {
	_, err := strconv.ParseUint(s, 10, 64)
	return err == nil
}

func isPrivateKey(s string) bool {
	s = strings.TrimPrefix(s, "0x")
	match, _ := regexp.MatchString("^[0-9a-fA-F]{64}$", s)
	return match
}
