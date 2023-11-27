package todo

import (
	errDomain "github.com/YukiOnishi1129/go-docker-firebase-restapi/domain/error"
	"github.com/oklog/ulid/v2"
	"math/rand"
	"time"
	"unicode/utf8"
)

//ビジネスルールを実装

type Todo struct {
	id string
	//userID  string
	title       string
	description string
}

func newTodo(id string, title string, description string) (*Todo, error) {
	//titleのバリデーション
	if utf8.RuneCountInString(title) <= titleLengthMin || utf8.RuneCountInString(title) >= titleLengthMax {
		return nil, errDomain.NewError("タイトルの値が不正です")
	}

	//contentのバリデーション
	if utf8.RuneCountInString(description) <= descriptionLengthMin || utf8.RuneCountInString(description) >= descriptionLengthMax {
		return nil, errDomain.NewError("説明の値が不正です")
	}

	return &Todo{
		id:          id,
		title:       title,
		description: description,
	}, nil
}

func NewTodo(
	title string,
	description string) (*Todo, error) {
	t := time.Now()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	id := ulid.MustNew(ulid.Timestamp(t), entropy)
	return newTodo(
		id.String(),
		title,
		description)
}

func UpdateTodo(
	id string,
	title string,
	description string) (*Todo, error) {
	return newTodo(
		id,
		title,
		description)
}

func (t *Todo) ID() string {
	return t.id
}

func (t *Todo) Title() string {
	return t.title
}

func (t *Todo) Description() string {
	return t.description
}

const (
	// titleの最大値/最小値
	titleLengthMin = 1
	titleLengthMax = 20

	//descriptionの最大値/最小値
	descriptionLengthMin = 1
	descriptionLengthMax = 1000
)
