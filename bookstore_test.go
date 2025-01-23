package bookstore_test

import (
	"bookstore"
	"testing"
)

func TestBook(t *testing.T){
	t.Parallel()
	_ = bookstore.Book{
		Title: "Spark Joy",
		Author: "Marie Kondo",
		Copies: 2,
	}
}

func TestBuy(t *Testing.T){
	t.Parallel()
	b := bookstore.Book{
		Title: "Nicholas Chuckleby",
		Author: "Charles Dickens",
		Copies: 8,
	}
	tc
}