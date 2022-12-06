package data

import (
	"testing"

	"github.com/canyolal/hypercasual-inventories/internal/assert"

	_ "github.com/lib/pq"
)

func TestPublisherInsert(t *testing.T) {
	tests := []struct {
		name          string
		publisherName string
		publisherLink string
		want          string
	}{
		{
			name:          "New Valid Publisher",
			publisherName: "Test",
			publisherLink: "https://test.com",
		},
		{
			name:          "Link Not Provided",
			publisherName: "sample test",
		},
		{
			name:          "Name Not Provided",
			publisherLink: "https://testtest.com",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := PublisherModel{db}

			publisher := &Publisher{
				Name: tt.publisherName,
				Link: tt.publisherLink,
			}

			err := m.Insert(publisher)

			assert.NilError(t, err)
		})
	}
}
