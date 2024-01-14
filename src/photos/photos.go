package photos

import "encoding/json"

type Photo struct {
	AlbumId      int    `json:"albumId"`
	Id           int    `json:"id"`
	Title        string `json:"title"`
	Url          string `json:"url"`
	ThumbnailUrl string `json:"thumbnailUrl"`
}

func JSONParse(data []byte) []Photo {
	photos := []Photo{}

	if err := json.Unmarshal(data, &photos); err != nil {
		panic(err)
	}
	return photos
}
