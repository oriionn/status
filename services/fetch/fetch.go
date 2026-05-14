package fetch

import "net/url"

func Fetch(rawUrl string) (Data, error) {
	u, err := url.Parse(rawUrl)
	if err != nil {
		return Data{}, err
	}

	if u.Scheme == "mc" {
		return Minecraft(rawUrl)
	}

	return Http(rawUrl)
}
