package utils

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestGetOutboundIP(t *testing.T) {
	ip := GetOutboundIP()
	t.Log(ip)
	// require ip like :*.*.*.*
	require := require.New(t)
	regexIP := `^\d{1,3}\.\d{1,3}\.\d{1,3}\.\d{1,3}$`
	require.Regexp(regexIP, ip)
}
