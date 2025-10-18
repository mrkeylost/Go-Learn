package notes

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"
)

// struct tages : yang dipake buat json tags
type Note struct {
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// contructor
func New(title, content string) (Note, error) {

	if len(title) < 1 || len(content) < 1 {
		return Note{}, errors.New("Invalid Data")
	}

	return Note{
		Title:     title,
		Content:   content,
		CreatedAt: time.Now(),
	}, nil
}

func (data Note) Display() {
	fmt.Printf("Your Note titled %v with this content\n\n%v\n", data.Title, data.Content)
}

func (data Note) Save() error {
	//var fileName string

	fileName := strings.ReplaceAll(data.Title, " ", "_")
	fileName = strings.ToLower(data.Title) + ".json"

	json, err := json.Marshal(data)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}
