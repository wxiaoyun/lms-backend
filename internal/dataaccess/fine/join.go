package fine

const (
	JoinUser     = "JOIN users ON loans.user_id = users.id"
	JoinLoan     = "JOIN loans ON fines.loan_id = loans.id"
	JoinBookCopy = "JOIN book_copies ON loans.book_copy_id = book_copies.id"
	JoinBook     = "JOIN books ON book_copies.book_id = books.id"
)
