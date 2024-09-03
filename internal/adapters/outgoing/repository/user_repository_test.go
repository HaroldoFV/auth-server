package database

import (
	"auth-server/internal/configs"
	"auth-server/internal/domain"
	"database/sql"
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"log"
	"os"
	"path/filepath"
	"testing"
)

type UserRepositoryTestSuite struct {
	suite.Suite
	DB         *sql.DB
	Repository *UserRepository
	Config     *configs.Conf
}

func (suite *UserRepositoryTestSuite) SetupSuite() {
	dir, _ := os.Getwd()

	config, err := configs.LoadConfig(dir)
	if err != nil {
		rootDir := filepath.Join(dir, "..", "..")
		config, err = configs.LoadConfig(rootDir)
		if err != nil {
			fmt.Println("Error loading configurations:", err)
			panic(err)
		}
	}
	suite.Config = config

	connectionString := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		config.TESTDBHost,
		config.TESTDBPort,
		config.TESTDBUser,
		config.TESTDBPassword,
		config.TESTDBName,
	)

	db, err := sql.Open(config.DBDriver, connectionString)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}

	suite.DB = db
	suite.Repository = NewUserRepository(db)

	_, err = suite.DB.Exec(`
	CREATE TABLE IF NOT EXISTS users
	(
		id         UUID PRIMARY KEY,
		name       VARCHAR(50)        NOT NULL,
		email      VARCHAR(50) UNIQUE NOT NULL,
		password   VARCHAR(255)       NOT NULL,
		status     VARCHAR(15)        NOT NULL,
		created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
	);
	`)
	if err != nil {
		log.Fatalf("Error creating table: %v", err)
	}
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	_, err := suite.DB.Exec("DELETE FROM users;")
	if err != nil {
		log.Fatal(err)
	}
}

func (suite *UserRepositoryTestSuite) TestCreateUser() {
	product, err := domain.NewUser("John Doe", "j@j.com", "12345678")
	assert.NoError(suite.T(), err)

	err = suite.Repository.Create(product)
	assert.NoError(suite.T(), err)

	var count int
	err = suite.DB.QueryRow("SELECT COUNT(*) FROM users WHERE id = $1", product.GetID()).Scan(&count)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), 1, count)
}

func (suite *UserRepositoryTestSuite) TestCreateUserWithInvalidData() {
	user, err := domain.NewUser("", "", "")
	assert.Error(suite.T(), err)

	if user != nil {
		err = suite.Repository.Create(user)
		assert.Error(suite.T(), err)
	}
}

func TestUserRepositoryTestSuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
