package tmpdb

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"namename"`
}

var Movies = []Movie{
	{
		ID:    "1",
		Isbn:  "00001",
		Title: "Title 1",
		Director: &Director{
			FirstName: "Bill",
			LastName:  "Gates",
		},
	},
	{
		ID:    "2",
		Isbn:  "00002",
		Title: "Title 2",
		Director: &Director{
			FirstName: "Martin",
			LastName:  "Freeman",
		},
	},
}
