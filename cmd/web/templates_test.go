package main

import (
	"testing"
	"time"

	"github.com/dchupp/snippetbox/internal/assert"
	"github.com/dchupp/snippetbox/ui"
)

func TestHumanDate(t *testing.T) {
	tests := []struct {
		name string
		tm   time.Time
		want string
	}{
		{
			name: "UTC",
			tm:   time.Date(2023, 3, 17, 10, 15, 0, 0, time.UTC),
			want: "17 Mar 2023 at 10:15",
		},
		{
			name: "Empty",
			tm:   time.Time{},
			want: "",
		},
		{
			name: "CET",
			tm:   time.Date(2023, 3, 17, 9, 15, 0, 0, time.FixedZone("CET", 1*60*60)),
			want: "17 Mar 2023 at 09:15",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			hd := ui.HumanDate(tt.tm)

			// Use the new assert.Equal() helper to compare the expected and
			// actual values.
			assert.Equal(t, hd, tt.want)
		})
	}
}
