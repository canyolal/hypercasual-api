package data

import (
	"testing"

	"github.com/canyolal/hypercasual-inventories/internal/assert"
)

func TestEmailExistById(t *testing.T) {
	tests := []struct {
		name   string
		userID int64
		want   bool
	}{
		{
			name:   "Valid ID",
			userID: 1,
			want:   true,
		},
		{
			name:   "Zero ID",
			userID: 0,
			want:   false,
		},
		{
			name:   "Non-Existing ID",
			userID: 2,
			want:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := EmailModel{db}

			exists, err := m.Exists(tt.userID)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}

func TestEmailExistByMail(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "Non-existent Email",
			email: "test@test.com",
			want:  false,
		},
		{
			name:  "Existent E-mail",
			email: "selami@sahin.com",
			want:  true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := EmailModel{db}

			exists, err := m.ExistsByEmail(tt.email)

			assert.Equal(t, exists, tt.want)
			assert.NilError(t, err)
		})
	}
}

func TestEmailDelete(t *testing.T) {
	tests := []struct {
		name  string
		email string
		want  bool
	}{
		{
			name:  "Non-Existent Email",
			email: "test@test.com",
		},
		{
			name:  "Existent E-mail",
			email: "selami@sahin.com",
		},
		{
			name:  "Empty E-mail",
			email: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db := newTestDB(t)

			m := EmailModel{db}

			err := m.Delete(tt.email)

			assert.NilError(t, err)
		})
	}
}
