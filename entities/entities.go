package entities

type User struct {
	ID         uint   `gorm:"primaryKey" json:"id"`
	Name       string `gorm:"not null" json:"name"`
	Surname    string `json:"surname"`
	Patronymic string `json:"patronymic"`
}

type Book struct {
	ID          uint      `gorm:"primaryKey" json:"id"`
	ISBN        uint      `gorm:"not null;unique" json:"isbn"`
	Title       string    `gorm:"not null" json:"title"`
	PublisherID uint      `gorm:"not null" json:"-"`
	Publisher   Publisher `gorm:"OnDelete:CASCADE" json:"publisher"`
	Edition     uint      `gorm:"not null" json:"edition"`
	Description string    `json:"description"`
	Author      string    `gorm:"not null" json:"author"`
	Year        int       `json:"year"`
}

type BookItems struct {
	ID         uint   `json:"id"`
	ItemsCount uint   `json:"items_count"`
	Takers     []User `json:"takers"`
}

type BookEntity struct {
	ID             uint `gorm:"primaryKey"`
	Book           Book `gorm:"OnDelete:CASCADE"`
	BookID         uint `gorm:"not null"`
	IsTaken        bool
	Taker          User
	TakerID        uint
	Place          string
	PreviousTakers []User `gorm:"many2many:book_entity_users"`
}

type Publisher struct {
	ID   uint   `gorm:"primaryKey" json:"id"`
	Name string `gorm:"not null" json:"name"`
	Addr string `json:"addr"`
	Code uint   `gorm:"not null;unique" json:"code"`
}
