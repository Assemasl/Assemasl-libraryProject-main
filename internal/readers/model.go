package readers

type Reader struct {
	Id       int    `json:"id"`
	FullName string `json:"full_name"`
}

type UpdateBody struct {
	FullName string `json:"full_name"`
}
