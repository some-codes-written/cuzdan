package users

import (
	"errors"
	"time"

	"immortality/pkg/common"
	"immortality/pkg/domain/users/user_models"

	"gorm.io/gorm"
)

type UserStore struct {
	common.StoreBase
}

func NewUserStore() *UserStore {
	store := new(UserStore)
	store.Connect()
	return store
}

func (s *UserStore) VerifyCredential(email string, password string) (bool, error) {

	var user *user_models.User
	var res *gorm.DB

	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res = tx.Where("email = ?", email).Table(user_models.USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		var credentials *user_models.Credential
		res = tx.Where("user_id = ?", user.ID).Table(user_models.CREDENTIALS).First(&credentials)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("credentials not found")
		}
		if credentials.Password != password {
			return errors.New("wrong password")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) GetUser(id uint) (*user_models.User, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {

		res := tx.First(&user_models.User{}, id).Table(user_models.USERS)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("User not found")
		}
		res.First(&user)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}

func (s *UserStore) GetUserByEmail(email string) (*user_models.User, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("email = ?", email).Table(user_models.USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("User not found")
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}

func (s *UserStore) UserIfExistsByEmail(email string) (bool, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("email = ?", email).Table(user_models.USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) UserIfExistsByGsm(gsm string) (bool, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("gsm = ?", gsm).Table(user_models.USERS).First(&user)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("user not found")
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}

func (s *UserStore) CreateUser(model *user_models.User) (*user_models.User, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Create(&model)
		if res.Error != nil {
			return res.Error
		}
		user, _ = s.GetUser(model.ID)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}

func (s *UserStore) GetUsers() ([]user_models.User, error) {
	var users []user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(user_models.USERS).Find(&users)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return users, nil
}

func (s *UserStore) GenerateToken(model *user_models.User) (*user_models.UserToken, error) {
	token := common.GenerateToken()
	tokenModel := &user_models.UserToken{
		Token:          token,
		UserID:         model.ID,
		ExpirationDate: time.Now().Add(time.Hour * 24 * 7),
		IsActive:       true,
	}
	tokenModel.SetDefaultsviaCreation()
	err := s.Create(tokenModel)
	if err != nil {
		return nil, err
	}
	return tokenModel, nil
}

func (s *UserStore) GetToken(token string) (*user_models.UserToken, error) {
	var tokenModel *user_models.UserToken
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("token = ?", token).Where("is_active = ?", 1).Table("user_tokens").First(&tokenModel)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return tokenModel, nil
}

func (s *UserStore) GetTokens() ([]*user_models.UserToken, error) {
	var tokens []*user_models.UserToken
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table("user_tokens").Where("is_active = ?", 1).Find(&tokens)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return tokens, nil
}

func (s *UserStore) ExpireToken(token string) (*user_models.UserToken, error) {
	var tokenModel *user_models.UserToken
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("token = ?", token).Table("user_tokens").First(&tokenModel)
		if res.Error != nil {
			return res.Error
		}
		tokenModel.ExpirationDate = time.Now()
		tokenModel.IsActive = false
		tokenModel.SetDefaultsviaUpdation()
		res = tx.Save(&tokenModel)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return tokenModel, nil
}

func (s *UserStore) CreateCredential(model *user_models.Credential) (*user_models.Credential, error) {
	var credential *user_models.Credential
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		model.SetDefaultsviaCreation()
		res := tx.Create(&model)
		if res.Error != nil {
			return res.Error
		}
		credential, _ = s.GetCredential(model.ID)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return credential, nil
}

func (s *UserStore) GetCredential(id uint) (*user_models.Credential, error) {
	var credential *user_models.Credential
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(user_models.CREDENTIALS).First(&credential)
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return errors.New("Credential not found")
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return credential, nil
}

func (s *UserStore) GetCredentials() ([]user_models.Credential, error) {
	var credentials []user_models.Credential
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Table(user_models.CREDENTIALS).Find(&credentials)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return credentials, nil
}

func (s *UserStore) GetCredentialsByUserId(userId uint) ([]user_models.Credential, error) {
	var credentials []user_models.Credential
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("user_id = ?", userId).Table(user_models.CREDENTIALS).Find(&credentials)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return credentials, nil
}

func (s *UserStore) UpdateCredential(model *user_models.Credential) (*user_models.Credential, error) {
	var credential *user_models.Credential
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		model.SetDefaultsviaCreation()
		res := tx.Save(&model)
		if res.Error != nil {
			return res.Error
		}
		credential, _ = s.GetCredential(model.ID)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return credential, nil
}

func (s *UserStore) DeleteCredential(id uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(user_models.CREDENTIALS).Delete(&user_models.Credential{})
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return txres
	}
	return nil
}

func (s *UserStore) DeleteCredentialsByUserId(userId uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("user_id = ?", userId).Table(user_models.CREDENTIALS).Delete(&user_models.Credential{})
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return txres
	}
	return nil
}

func (s *UserStore) UpdateUser(model *user_models.User) (*user_models.User, error) {
	var user *user_models.User
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		model.SetDefaultsviaUpdation()
		res := tx.Save(&model)
		if res.Error != nil {
			return res.Error
		}
		user, _ = s.GetUser(model.ID)
		return nil
	})
	if txres != nil {
		return nil, txres
	}
	return user, nil
}

func (s *UserStore) DeleteUser(id uint) error {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Where("id = ?", id).Table(user_models.USERS).Delete(&user_models.User{})
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return txres
	}
	return nil
}

func (s *UserStore) TokenExists(token string) (bool, error) {
	var exists bool
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res := tx.Model(user_models.UserToken{}).Select("count(*) > 0").Where("token = ?", token).Where("is_active", 1).Find(&exists)
		if res.Error != nil {
			return res.Error
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return exists, nil
}

func (s *UserStore) ExpireAllTokens() (bool, error) {
	txres := s.Db.Transaction(func(tx *gorm.DB) error {
		res, err := s.GetTokens()
		if err != nil {
			return err
		}
		for _, token := range res {
			token.ExpirationDate = time.Now()
			token.IsActive = false
			token.SetDefaultsviaUpdation()
			res := tx.Save(&token)
			if res.Error != nil {
				return res.Error
			}
		}
		return nil
	})
	if txres != nil {
		return false, txres
	}
	return true, nil
}
