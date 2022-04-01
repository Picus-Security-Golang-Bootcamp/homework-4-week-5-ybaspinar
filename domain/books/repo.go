package books

import (
	"gorm.io/gorm"
)

type BooksRepository struct {
	db *gorm.DB
}

func NewBooksepository(db *gorm.DB) *BooksRepository {
	return &BooksRepository{db: db}
}
func (b *BooksRepository) Migrations() {
	b.db.AutoMigrate(&Book{})
	b.db.AutoMigrate(&Author{})

}
func (b *BooksRepository) InsertData(booklist []Book) {
	//Checks if author exist if not add to db
	for _, book := range booklist {
		b.db.Where(map[string]interface{}{"authorname": book.Author.Authorname}).FirstOrCreate(&book.Author)
	}
	//Checks if book exist if not add to db
	for _, book := range booklist {
		b.db.FirstOrCreate(&book, map[string]interface{}{"book_id": book.BookID})
		println(book.Author.AuthorID)
		b.db.Table("books").Where(map[string]interface{}{"book_id": book.BookID}).
			Updates(map[string]interface{}{"author_id": book.Author.AuthorID})
	}
}

//ListAll gets all the entries from database
func (b *BooksRepository) ListAll() []Book {
	var books []Book
	b.db.Find(&books)
	return books
}

//Search returns books with given title CASE SENSETIVE
func (b *BooksRepository) Search(key string) []Book {
	var books []Book
	b.db.Where("booktitle LIKE ?", key).First(&books)
	return books
}

//GetById returns books with given ID
func (b *BooksRepository) GetById(key int) []Book {
	var books []Book
	b.db.Where("book_id = ?", key).First(&books)
	return books
}

//NOT WORKING
func (b *BooksRepository) GetBooksWithAuthor(key string) []Book {
	var books []Book
	b.db.
		Joins("LEFT JOIN books ON books.author_id = authors.author_id").
		Where("authorname LIKE ?", key).Find(&books)
	return books
}

//NOT WORKING
func (b *BooksRepository) GetAuthorWithBooks(key string) []Book {
	var books []Book
	b.db.Table("books").Where("authorname LIKE ?", key).Find(&books)
	return books
}

//Delete soft deletes given ID
func (b *BooksRepository) Delete(key int) {
	b.db.Table("books").Where(map[string]interface{}{"book_id": key}).
		Delete(map[string]interface{}{"book_id": key})
}

//PermaDelete deletes given ID
func (b *BooksRepository) PermaDelete(key int) {
	b.db.Unscoped().Table("books").Where(map[string]interface{}{"book_id": key}).
		Delete(map[string]interface{}{"book_id": key})
}

//FindDeleted returns deleted books with given string
func (b *BooksRepository) FindDeleted(key string) []Book {
	var books []Book
	b.db.Unscoped().Where("booktitle LIKE ?", key).Find(&books)
	return books
}
