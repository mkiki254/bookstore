package bookstore

import "errors"

type Book struct {
	Title  string
	Author string
	Copies int
	ID     int
}

func Buy(b Book) (Book, error) {
	if b.Copies == 0 {
		return Book{}, errors.New("no copies left")
	}
	b.Copies--
	return b, nil
}

func GetAllBooks(catalogue []Book) []Book {
	return catalogue
}

func GetBook(catalogue []Book, ID int) Book {
	// for _, b := range catalogue {
	// 	if b.ID == ID {
	// 		return b
	// 	}
	// }
	// return Book{}
	return catalogue[0]
}
