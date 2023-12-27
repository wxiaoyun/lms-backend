package bookmark

const (
	JoinUser = "JOIN users on bookmarks.user_id = users.id"
	JoinBook = "JOIN books on bookmarks.book_id = books.id"
)
