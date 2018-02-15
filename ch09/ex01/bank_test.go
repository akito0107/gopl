package main

import "testing"

func TestWithdraw(t *testing.T) {

	cases := []struct {
		name     string
		depos    int
		withdraw int
		out      bool
	}{
		{
			name:     "basic",
			depos:    100,
			withdraw: 50,
			out:      true,
		},
		{
			name:     "negative",
			depos:    100,
			withdraw: 200,
			out:      false,
		},
		{
			name:     "same",
			depos:    100,
			withdraw: 100,
			out:      true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			Deposit(c.depos)
			if act := Withdraw(c.withdraw); act != c.out {
				t.Errorf("must be %t, but %t\n", c.out, act)
			}
			Clear()
		})
	}
}
