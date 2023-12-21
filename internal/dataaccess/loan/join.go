package loan

const (
	JoinBook = "JOIN books ON loans.book_id = books.id"
	JoinUser = "JOIN users ON users.id = loans.user_id"
)
