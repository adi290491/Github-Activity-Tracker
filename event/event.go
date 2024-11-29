package event

import "fmt"

type Event struct {
	Id   string `json:"id"`
	Type string `json:"type"`

	Actor struct {
		Id    int    `json:"id"`
		Login string `json:"login"`
		URL   string `json:"url"`
	} `json:"actor"`

	Repo struct {
		Id   int    `json:"id"`
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"repo"`

	Payload struct {
		RepositoryID int `json:"repository_id"`

		Commits []struct {
			Author struct {
				Email string `json:"email"`
				Name  string `json:"name"`
			} `json:"author"`
			Message  string `json:"message"`
			Distinct bool   `json:"distinct"`
			URL      string `json:"url"`
		} `json:"commits"`
	} `json:"payload,omitempty"`

	Public    bool   `json:"public"`
	CreatedAt string `json:"created_at"`
}

func (e *Event) String() string {
	return fmt.Sprintf("Event: %s\nActor: %s\nRepo: %s\nPayload: %+v\nPublic: %t\nCreated At: %s",
		e.Type, e.Actor.Login, e.Repo.Name, e.Payload, e.Public, e.CreatedAt)
}
