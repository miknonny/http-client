package examples

/*
"current_user_url": "https://api.github.com/user",
  "authorizations_url": "https://api.github.com/authorizations",
*/

type Endpoints struct {
	CurrentUser       string `json:"current_user_url"`
	AuthorizationsUrl string `json:"authorizations_url"`
	RepositoryUrl     string `json:"repository_url"`
}

func GetEndPoints() (*Endpoints, error) {
	response, err := httpClient.Get("https://api.github.com", nil)
	if err != nil {
		// Deal with this error.
		return nil, err
	}

	// fmt.Printf("Status Code: %d\n", response.StatusCode())
	// fmt.Printf("Status: %s\n", response.Status())
	// fmt.Printf("Status: %s\n", response.String())

	var endpoints Endpoints
	if err := response.UnmarshalJson(&endpoints); err != nil {
		// Deal with this error.
		return nil, err
	}

	return &endpoints, nil
}
