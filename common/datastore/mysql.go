package datastore

import (
	"database/sql"

	"github.com/shunsukw/go-chat/models/socialmedia"

	"github.com/shunsukw/go-chat/models"

	"log"

	// mysql ...
	_ "github.com/go-sql-driver/mysql"
)

// MySQLDatastore ...
type MySQLDatastore struct {
	*sql.DB
}

// NewMySQLDatastore ...
func NewMySQLDatastore(dataSourceName string) (*MySQLDatastore, error) {
	connection, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	return &MySQLDatastore{
		DB: connection,
	}, nil
}

// CreateUser ...
func (m *MySQLDatastore) CreateUser(user *models.User) error {
	tx, err := m.Begin()
	if err != nil {
		log.Print(err)
	}

	defer tx.Rollback()

	stmt, err := tx.Prepare("INSERT INTO user(uuid, username, first_name, last_name, email, password_hash) VALUES (?,?,?,?,?,?)")
	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(user.UUID, user.Username, user.FirstName, user.LastName, user.Email, user.PasswordHash)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

// GetUser ...
func (m *MySQLDatastore) GetUser(username string) (*models.User, error) {
	stmt, err := m.Prepare("SELECT uuid, username, first_name, last_name, email, password_hash, UNIX_TIMESTAMP(created_ts), UNIX_TIMESTAMP(updated_ts) FROM user WHERE username = ?")
	if err != nil {
		log.Print(err)
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(username)
	u := models.User{}
	err = row.Scan(&u.UUID, &u.Username, &u.FirstName, &u.LastName, &u.Email, &u.PasswordHash, &u.TimestampCreated, &u.TimestampModified)
	return &u, err
}

// Close ...
func (m *MySQLDatastore) Close() {
	m.Close()
}

func (m *MySQLDatastore) FetchPosts(owner string) ([]socialmedia.Post, error) {
	posts := make([]socialmedia.Post, 0)
	stmt, err := m.Prepare("select p.uuid, p.title, p.body, p.mood, UNIX_TIMESTAMP(p.created_ts), UNIX_TIMESTAMP(p.updated_ts), up.profile_image_path, u.username from post p, user u, user_profile up where p.uuid = u.uuid and p.uuid = up.uuid and (p.uuid = ? or p.uuid in (select friend_uuid from friend_relation where owner_uuid=?) ) order by p.created_ts desc")
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	rows, err := stmt.Query(owner, owner)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		p := socialmedia.Post{}
		err := rows.Scan(&p.UUID, &p.Caption, &p.MessageBody, &p.RawMoodValue, &p.TimeCreatedUnixTS, &p.TimeModifiedUnixTS, &p.ProfileImagePath, &p.Username)
		if err != nil {
			return nil, err
		}

		return posts, nil
	}
}
