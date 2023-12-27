package loan

const (
	JoinBookCopy = "JOIN book_copies ON loans.book_copy_id = book_copies.id"
	JoinBook     = "JOIN books ON book_copies.book_id = books.id"
	JoinUser     = "JOIN users ON loans.user_id = users.id"
)
