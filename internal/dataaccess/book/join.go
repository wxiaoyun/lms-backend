package book

const (
	JoinLoan        = "JOIN loans ON loans.book_id = books.id"
	JoinReservation = "JOIN reservations ON reservations.book_id = books.id"
)
