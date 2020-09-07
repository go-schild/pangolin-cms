package cms

type WebURL struct {
	URL  string
	Page Page
}

type WebURLSet map[string]WebURL

func loadWebURLSet(url string) (WebURLSet, error) {
	var set WebURLSet
	err := unmarshal(url, &set)
	if err != nil {
		return nil, err
	}
	return set, nil
}
