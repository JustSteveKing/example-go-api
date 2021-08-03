package database

type Client struct {
	ID    int    `json:"id",gorm:"primaryKey"`
	Uuid  string `json:"uuid"`
	Name  string `json:"name"`
	Email string `json:"email"`
}
