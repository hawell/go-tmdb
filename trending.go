package tmdb

import (
	"fmt"
)

// GetTrending get the daily or weekly trending items.
// https://developers.themoviedb.org/3/trending/get-trending
func (tmdb *TMDb) GetTrending(media_type, time_window string, page int) (*MoviePagedResults, error) {
	var nowPlaying MoviePagedResults
	uri := fmt.Sprintf("%s/trending/%v/%v?page=%d&api_key=%s", baseURL, media_type, time_window, page, tmdb.apiKey)
	result, err := getTmdb(uri, &nowPlaying)
	return result.(*MoviePagedResults), err
}
