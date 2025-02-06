package bot

type User struct {
	Id       uint    `gorm:"id,primarykey"`
	ChatId   int64   `gorm:"chat_id"`
	Name     *string `gorm:"name"`
	IsActive bool    `gorm:"is_active,not null"`
}

type UsersService struct {
	DbConnection *DbConnection
}

func NewUsersService(dbConnection *DbConnection) *UsersService {
	return &UsersService{
		DbConnection: dbConnection,
	}
}

func (s *UsersService) RegisterUser(chatId int64, name *string) (*User, error) {
	var user User
	err := s.DbConnection.Db.Where("chat_id = ?", chatId).First(&user).Error

	if err != nil {
		user = User{
			ChatId:   chatId,
			Name:     name,
			IsActive: true,
		}
		if err := s.DbConnection.Db.Create(&user).Error; err != nil {
			return nil, err
		}
	}

	return &user, nil
}

func (s *UsersService) GetAllUsers() ([]*User, error) {
	var users []*User
	if err := s.DbConnection.Db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (s *UsersService) GetById(id int64) (*User, error) {
	var user User
	if err := s.DbConnection.Db.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersService) GetByChatId(chatId int64) (*User, error) {
	var user User
	if err := s.DbConnection.Db.Where("chat_id = ?", chatId).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (s *UsersService) CheckChatIsRegistered(chatId int64) bool {
	_, err := s.GetByChatId(chatId)
	return err == nil
}
