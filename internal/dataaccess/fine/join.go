package fine

const (
	JoinBook = "JOIN books ON loans.book_id = books.id"
	JoinLoan = "JOIN loans ON loans.id = fines.loan_id"
	JoinUser = "JOIN users ON users.id = loans.user_id"
)
