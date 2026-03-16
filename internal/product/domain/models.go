package domain

type ItemTypeDto struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Type      int     `json:"type"`
	Price     float64 `json:"price"`
	Image     string  `json:"image"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}

type ItemDto struct {
	Price float64 `json:"price"`
	Type  int     `json:"type"`
}
