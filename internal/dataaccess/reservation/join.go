package reservation

const (
	JoinBook = "JOIN books ON reservations.book_id = books.id"
	JoinUser = "JOIN users ON users.id = reservations.user_id"
)
