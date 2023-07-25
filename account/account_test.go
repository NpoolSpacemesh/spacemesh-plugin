package account

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAccount(t *testing.T) {
	priStr := "e0aef1ce781f28ed8a88f5a3fd87e4ffe399f3750b3fee640e2606c5e2922217f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	pubStr := "f7d6c5e814c89faf8a2866e65000f31253f5c42568ca3cf4f0c6e70cbc878559"
	stestStr := "stest1qqqqqqru6pcet8crur4qw52ac73w3mamj0mk0vqu5zg3w"

	acc, err := CreateAccountFromHexPri(priStr)
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, acc.Pub, pubStr)
	assert.Equal(t, acc.Pri, priStr)
	assert.Equal(t, acc.GetAddress(TestHRP), stestStr)

	acc1, err := CreateAccount()
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, true, strings.HasPrefix(acc1.GetAddress(TestHRP).String(), TestHRP))

	acc2, err := CreateAccount()
	if !assert.Nil(t, err) {
		return
	}
	assert.Equal(t, true, strings.HasPrefix(acc2.GetAddress(MainHRP).String(), MainHRP))
}
