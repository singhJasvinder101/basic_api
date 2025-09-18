package sqlite

// Concrete

import (
	"database/sql"
	"fmt"
	"log/slog"

	_ "github.com/mattn/go-sqlite3"
	"github.com/singhJasvinder101/basic-api/internal/config"
	"github.com/singhJasvinder101/basic-api/internal/types"
)

type Sqlite struct {
	Db *sql.DB
}

// constructor for Sqlite
func New(cfg *config.Config) (*Sqlite, error) {
	db, err := sql.Open("sqlite3", cfg.StoragePath)
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(`create table if not exists students(
		id integer primary key autoincrement,
		name text,
		email text unique,
		age integer
	)`)

	if err != nil {
		return nil, err
	}

	return &Sqlite{
		Db: db,
	}, nil
}

func (s *Sqlite) CreateStudent(name string, email string, age int) (int64, error) {
	stmt, err := s.Db.Prepare("insert into students (name, email, age) values (?, ?, ?)")
	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	res, err := stmt.Exec(name, email, age)
	if err != nil {
		return 0, err
	}

	lastInsertedId, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return lastInsertedId, nil
}

func (s *Sqlite) GetStudentById(id int64) (types.Student, error) {
	stmt, err := s.Db.Prepare("select id, name, email, age from students where id = ?")
	if err != nil {
		return types.Student{}, err
	}

	defer stmt.Close()

	var student types.Student

	err = stmt.QueryRow(id).Scan(&student.ID, &student.Name, &student.Email, &student.Age)
	if err != nil {
		if err == sql.ErrNoRows {
			slog.Error("student not found", slog.Int64("id", id))
			return types.Student{}, fmt.Errorf("student with id %d not found", id)
		}
		return types.Student{}, err
	}

	return student, nil
}


