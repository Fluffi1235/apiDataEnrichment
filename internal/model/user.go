package model

type User struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	SurName    string `json:"surName"`
	Patronymic string `json:"patronymic"`
	Age        int    `json:"age"`
	Gender     string `json:"gender"`
	Country    string `json:"country"`
}
