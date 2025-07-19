package dictionary

import "errors"

// alias of type
type Dictionary map[string]string

var (
	notFoundError     = errors.New("word not found")
	existsError       = errors.New("word already exists")
	cannotUpdateError = errors.New("cannot update non-existing word")
)

// NewDictionary 함수는 새로운 사전을 생성합니다.
func NewDictionary() Dictionary {
	return Dictionary{}
}

// Search 함수는 사전에서 주어진 키에 해당하는 값을 반환합니다.
func (d Dictionary) Search(word string) (string, error) {
	value, exists := d[word]
	if exists {
		return value, nil
	}
	return "", notFoundError
}

// Add 함수는 사전에 새로운 단어와 정의를 추가합니다.
func (d Dictionary) Add(word, def string) error {
	_, err := d.Search(word)

	if err == notFoundError {
		d[word] = def
		return nil
	}

	return existsError
}

// Update 함수는 사전에서 주어진 단어의 정의를 업데이트합니다.
func (d Dictionary) Update(word, def string) error {
	_, err := d.Search(word)

	switch err {
	case nil:
		d[word] = def
	case notFoundError:
		return cannotUpdateError
	}

	return nil
}

// Delete 함수는 사전에서 주어진 단어를 삭제합니다.
func (d Dictionary) Delete(word string) {
	delete(d, word)
}
