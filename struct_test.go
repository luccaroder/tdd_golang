package tddgo

import (
	"testing"

	"gotest.tools/v3/assert"
)

func TestWallet(t *testing.T) {
	type errorTestCases struct {
		description string
		deposit     Bitcoin
		withdraw    Bitcoin
		want        Bitcoin
		wantErr     error
	}

	for _, tt := range []errorTestCases{
		{
			description: "Withdraw insufficient funds",
			deposit:     Bitcoin(20),
			withdraw:    Bitcoin(100),
			want:        Bitcoin(20),
			wantErr:     ErrInsufficientFunds,
		},
		{
			description: "Withdraw",
			deposit:     Bitcoin(20),
			withdraw:    Bitcoin(2),
			want:        Bitcoin(18),
			wantErr:     nil,
		},
		{
			description: "Ballance",
			deposit:     Bitcoin(2),
			withdraw:    Bitcoin(0),
			want:        Bitcoin(2),
		},
	} {
		t.Run(tt.description, func(t *testing.T) {
			wallet := Wallet{}
			wallet.Deposit(tt.deposit)
			err := wallet.Withdraw(tt.withdraw)
			confirmBallance(t, wallet, tt.want)
			confirmErr(t, err, tt.wantErr)
		})
	}
}

func confirmBallance(t *testing.T, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Ballance()
	assert.Equal(t, got, want)
}

func confirmErr(t *testing.T, got error, want error) {
	t.Helper()
	assert.Equal(t, got, want)
}
