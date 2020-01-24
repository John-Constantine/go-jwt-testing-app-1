package model

type Todo struct {
	ID     uint64 `gorm:"primary_key;auto_increment" json:"id"`
	UserID uint64 `gorm:"not null" json:"user_id"`
	Title  string `gorm:"size:255;not null" json:"title"`
}

func (s *Server) CreateTodo(todo *Todo) (*Todo, error) {
	err := s.DB.Debug().Create(&todo).Error
	if err != nil {
		return nil, err
	}
	return todo, nil
}