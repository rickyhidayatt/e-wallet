package repository

import (
	"database/sql/driver"
	"e-wallet/model"
	"e-wallet/utils"
	"errors"
	"testing"
	"time"

	"log"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var dummyUsers = []model.User{
	{
		Id: utils.GenerateId(),
		Name: "DummyUser01",
		Email: "dummy01-email@mail.com",
		PhoneNumber: "098712345678",
		Password: "Dum01*123",
		Address: "Jl. Dummy No. 01",
		BirthDate: time.Date(1990, time.June, 1, 1, 1, 1, 1, time.UTC),
		ProfilePicture: "DummyImage01",
		CreatedAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
		UpdateAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
	},
	{
		Id: utils.GenerateId(),
		Name: "DummyUser02",
		Email: "dummy02-email@mail.com",
		PhoneNumber: "098743218765",
		Password: "Dum02*789",
		Address: "Jl. Dummy No. 02",
		BirthDate: time.Date(1990, time.June, 1, 1, 1, 1, 1, time.UTC),
		ProfilePicture: "DummyImage02",
		CreatedAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
		UpdateAt: time.Date(2010, time.December, 1, 1, 1, 1, 1, time.UTC),
	},
}

type UserRepositoryTestSuite struct {
	suite.Suite
	mockDb *sqlx.DB
	mockSql sqlmock.Sqlmock
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	mockDb, mockSql, err := sqlmock.New()
	sqlxDB := sqlx.NewDb(mockDb, "sqlmock")
	if err != nil {
		log.Fatalln("Error when opening db connection")
	}
	suite.mockDb = sqlxDB
	suite.mockSql = mockSql
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	suite.mockDb.Close()
}

func (suite *UserRepositoryTestSuite) TestGetUserById_Success() {
	dummyUser := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDb)
	actual, err := repo.GetUserById(dummyUser.Id)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has not passed unit test

func (suite *UserRepositoryTestSuite) TestGetUserById_Failed() {
	dummyUser := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id).WillReturnError(errors.New("failed"))
	repo := NewUserRepository(suite.mockDb)
	_, err := repo.GetUserById(dummyUser.Id)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("Query: could not match actual sql: \"SELECT * FROM users WHERE id=$1\" with expected regexp \"SELECT * FROM users WHERE id=$1\""), err)
} // This method has passed unit test

func (suite *UserRepositoryTestSuite) TestFindByEmail_Success() {
	dummyUser := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE email = $1").WithArgs(dummyUser.Email).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDb)
	actual, err := repo.FindByEmail(dummyUser.Email)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has not passed unit test

func (suite *UserRepositoryTestSuite) TestFindByEmail_Failed() {
	dummyUser := dummyUsers[0]
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE email = $1").WithArgs(dummyUser.Email).WillReturnError(errors.New("failed"))
	repo := NewUserRepository(suite.mockDb)
	_, err := repo.FindByEmail(dummyUser.Email)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("Query: could not match actual sql: \"SELECT * FROM users WHERE email = $1\" with expected regexp \"SELECT * FROM users WHERE email = $1\""), err)
} // This method has passed unit test


func (suite *UserRepositoryTestSuite) TestSaveUser_Success() {
	dummyUser := dummyUsers[0]
	query := `INSERT INTO users \(id, name, email, phone_number, password, address, birth_date, profile_picture, created_at, update_at\)
	VALUES \(\?, \?, \?, \?, \?, \?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt,).WillReturnResult(sqlmock.NewResult(1, 1))
	repo := NewUserRepository(suite.mockDb)
	err := repo.SaveUser(&dummyUser)
	assert.Nil(suite.T(), err)
} // This method has passed unit test

func (suite *UserRepositoryTestSuite) TestSaveUser_Failed() {
	dummyUser := dummyUsers[0]
	query := `INSERT INTO users \(id, name, email, phone_number, password, address, birth_date, profile_picture, created_at, update_at\)
	VALUES \(\?, \?, \?, \?, \?, \?, \?, \?, \?, \?\)`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt,).WillReturnError(errors.New("failed"))
	repo := NewUserRepository(suite.mockDb)
	err := repo.SaveUser(&dummyUser)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("failed"), err)
} // This method has passed unit test

func (suite *UserRepositoryTestSuite) TestUpdate_Success() {
	dummyUser := dummyUsers[0]
	query := `UPDATE users SET name=\?, email=\?, phone_number=\?, password=\?, address=\?, birth_date=\?, profile_picture=\?, update_at=\? WHERE id=\?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.UpdateAt, dummyUser.Id).WillReturnResult(driver.RowsAffected(1))
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDb)
	actual, err := repo.Update(&dummyUser)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has not passed unit test

func (suite *UserRepositoryTestSuite) TestUpdate_Failed() {
	dummyUser := dummyUsers[0]
	query := `UPDATE users SET name=\?, email=\?, phone_number=\?, password=\?, address=\?, birth_date=\?, profile_picture=\?, update_at=\? WHERE id=\?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.UpdateAt, dummyUser.Id).WillReturnResult(driver.RowsAffected(1))
	rows := sqlmock.NewRows([]string{"id", "name", "email", "phone_number", "password", "address", "birth_date", "profile_picture", "created_at", "update_at"})
	rows.AddRow(dummyUser.Id, dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.CreatedAt, dummyUser.UpdateAt)
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id).WillReturnRows(rows)
	repo := NewUserRepository(suite.mockDb)
	_, err := repo.Update(&dummyUser)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("Query: could not match actual sql: \"SELECT * FROM users WHERE id=$1\" with expected regexp \"SELECT * FROM users WHERE id=$1\""), err)
} // This method has passed unit test

func (suite *UserRepositoryTestSuite) TestSaveAvatar_Success() {
	dummyUser := dummyUsers[0]
	query := `UPDATE users SET name=\?, email=\?, phone_number=\?, password=\?, address=\?, birth_date=\?, profile_picture=\?, update_at=\? WHERE id=\?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.UpdateAt, dummyUser.Id).WillReturnResult(driver.RowsAffected(1))
	(sqlmock.NewResult(1, 1))
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id)
	repo := NewUserRepository(suite.mockDb)
	actual, err := repo.SaveAvatar(&dummyUser)
	assert.Nil(suite.T(), err)
	assert.NotNil(suite.T(), actual)
} // This method has not passed unit test

func (suite *UserRepositoryTestSuite) TestSaveAvatar_Failed() {
	dummyUser := dummyUsers[0]
	query := `UPDATE users SET name=\?, email=\?, phone_number=\?, password=\?, address=\?, birth_date=\?, profile_picture=\?, update_at=\? WHERE id=\?`
	suite.mockSql.ExpectExec(query).WithArgs(dummyUser.Name, dummyUser.Email, dummyUser.PhoneNumber, dummyUser.Password, dummyUser.Address, dummyUser.BirthDate, dummyUser.ProfilePicture, dummyUser.UpdateAt, dummyUser.Id).WillReturnResult(driver.RowsAffected(1))
	suite.mockSql.ExpectQuery("SELECT * FROM users WHERE id=$1").WithArgs(dummyUser.Id)
	repo := NewUserRepository(suite.mockDb)
	_, err := repo.SaveAvatar(&dummyUser)
	assert.Error(suite.T(), err)
	assert.Equal(suite.T(), errors.New("Query: could not match actual sql: \"SELECT * FROM users WHERE id=$1\" with expected regexp \"SELECT * FROM users WHERE id=$1\""), err)
} // This method has passed unit test

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}