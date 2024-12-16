package main

import (
  "fmt"
  "github.com/google/uuid"
)

type Book struct {
  ID string
  Title string
  Author string
  Genre string
  Year int
}

var books []Book
var bookMap = make(map[string]Book)

type LibraryImpl struct {
  books []Book
  bookMap map[string]Book
}

type Library interface {
  AddBook(book Book) string
  GetBookByID(id string) (Book, bool)
  SearchByName(name string) (Book, bool)
} 

func GenerateID() string {
  id := uuid.New()
  return id.String()
}

func (lib *LibraryImpl) AddBook(book Book) string {
  book.ID = GenerateID()
  lib.books = append(lib.books, book)
  lib.bookMap[book.ID] = book
  return book.ID
}

func (lib *LibraryImpl) GetByID(id string) (Book, bool) {
  book, found := lib.bookMap[id]
  return book, found
}

func (lib *LibraryImpl) SearchByName(name string) (Book, bool) {
  for _, book := range lib.books {
    if book.Title == name {
      return book, true
    }
  }
  return Book{}, false
}

func main() {
  library := LibraryImpl{
    books: []Book{},
    bookMap: make(map[string]Book),
  }
  // Добавление книг
  id1 := library.AddBook(Book{Title: "First", Author: "Ann", Genre : "Horror", Year: 2023})
  id2 := library.AddBook(Book{Title: "Second", Author: "Bob", Genre : "Classic", Year: 2024})
  id3 := library.AddBook(Book{Title: "Third", Author: "Cooper", Genre : "Humor", Year: 2029})

  fmt.Println("Added IDs:", id1, id2, id3)

  // Поиск по ID
  book, found := library.GetByID(id1)
  if found {
    fmt.Println("Found by ID:", book)
  } else {
    fmt.Println("Not Found")
  } 

  // Поиск по имени
  bookByName, found := library.SearchByName("WeDontHaveIt")
  //bookByName, found := library.SearchByName("Second")
  if found {
    fmt.Println("Found by name:", bookByName)
  } else {
    fmt.Println("Not found")
  }
}
