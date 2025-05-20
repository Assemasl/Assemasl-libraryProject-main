package authors

func GetAuthors() ([]Author, error) {
	return GetAllAuthors()
}

func GetAuthorService(id int) (Author, error) {
	return GetAuthorById(id)
}

func AddAuthorService(author *Author) error {
	return AddAuthor(author)
}

func UpdateAuthorService(readerId int, fullName string, specialization string) error {
	return ChangeAuthorById(readerId, fullName, specialization)
}
