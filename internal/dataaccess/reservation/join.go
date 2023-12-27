package reservation

const (
	JoinBookCopy = "JOIN book_copies ON reservations.book_copy_id = book_copies.id"
	JoinBook     = "JOIN books ON book_copies.book_id = books.id"
	JoinUser     = "JOIN users ON reservations.user_id = users.id"
)
