package pkg

type City struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Region     string `json:"region"`
	District   string `json:"district"`
	Population int    `json:"population"`
	Foundation int    `json:"foundation"`
}

type CityPopulation struct {
	Population *int `json:"population"`
}
