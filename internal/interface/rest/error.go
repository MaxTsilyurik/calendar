package rest

func BadRequest() {

}

type Error struct {
	Slug    string `json:"slug"`
	Message string `json:"message"`
	Status  int
}
