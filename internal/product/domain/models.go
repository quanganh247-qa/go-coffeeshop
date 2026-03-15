package domain

type ItemTypeDto struct {
	Name  string  `json:"name"`
	Type  int     `json:"type"`
	Price float64 `json:"price"`
	Image string  `json:"image"`
}

type ItemDto struct {
	Name      string  `json:"name"`
	Price     float64 `json:"price"`
	Type      int     `json:"type"`
	CreatedAt string  `json:"created_at"`
	UpdatedAt string  `json:"updated_at"`
}
