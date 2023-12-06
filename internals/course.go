package internals

import (
	"database/sql"
	"fmt"

	"github.com/google/uuid"
)

type Course struct {
	DB          *sql.DB
	ID          string
	Name        string
	Description string
	CategoryID  string
}

func NewCourse(db *sql.DB) *Course {
	return &Course{DB: db}
}

func (c *Course) Create(name, description, categoryID string) (Course, error) {
	id := uuid.New().String()

	_, err := c.DB.Exec("INSERT INTO courses (id, name ,description ,category_id )  VALUES ($1,$2,$3,$4)",
		id, name, description, categoryID)

	if err != nil {
		return Course{}, err
	}

	return Course{ID: id, Name: name, Description: description, CategoryID: categoryID}, nil

}

func (c *Course) FindByCategoryId(categoryID string) ([]Course, error) {

	query := fmt.Sprintf("SELECT id, name, description , category_id from courses where category_id = '%s'", categoryID)
	return c.findBuilder(query)

}

func (c *Course) findBuilder(query string) ([]Course, error) {
	rows, err := c.DB.Query(query)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	courses := []Course{}
	for rows.Next() {
		var id, name, description, categoryId string
		if err := rows.Scan(&id, &name, &description, &categoryId); err != nil {
			return nil, err
		}

		courses = append(courses, Course{
			ID:          id,
			Name:        name,
			Description: description,
		})

		return courses, nil
	}
	return courses, nil

}

func (c *Course) FindAll() ([]Course, error) {

	query := "SELECT id, name, description from courses"
	return c.findBuilder(query)

}
