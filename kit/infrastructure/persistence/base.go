package persistence

type Base struct {
	ID        []byte `gorm:"type:binary(16);primary_key"`
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
}
