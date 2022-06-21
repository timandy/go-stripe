package session

import (
	"testing"

	assert "github.com/stretchr/testify/require"
	stripe "github.com/timandy/go-stripe/v72"
	_ "github.com/timandy/go-stripe/v72/testing"
)

func TestBillingPortalSessionNew(t *testing.T) {
	session, err := New(&stripe.BillingPortalSessionParams{
		Customer:  stripe.String("cus_123"),
		ReturnURL: stripe.String("https://stripe.com/return"),
	})
	assert.Nil(t, err)
	assert.NotNil(t, session)
}
