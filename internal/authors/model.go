package authors

type Author struct {
	Id             int    `json:"id"`
	FullName       string `json:"full_name"`
	Specialization string `json:"specialization"`
}

type updateAuthorBody struct {
	FullName       string `json:"full_name"`
	Specialization string `json:"specialization"`
}
