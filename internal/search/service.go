package search

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"
)

// client is shared across providers, with a timeout so one slow upstream
// can't hang the request (unlike item.GetMetaData, which uses a bare http.Get).
var client = &http.Client{Timeout: 4 * time.Second}

// Service aggregates external catalog providers (books, movies, music).
type Service struct {
	tmdbKey string
}

// NewService returns a new search Service. tmdbKey may be empty, in which case
// the movies provider degrades silently to no results.
func NewService(tmdbKey string) Service {
	return Service{tmdbKey: tmdbKey}
}

// Search queries every provider concurrently and returns the combined results
// in a stable order (books, then movies, then music). A failing provider
// contributes an empty slice rather than failing the whole search.
func (serv *Service) Search(query string) []Result {
	var (
		wg                   sync.WaitGroup
		books, movies, music []Result
	)
	wg.Add(3)
	go func() { defer wg.Done(); books = serv.searchBooks(query) }()
	go func() { defer wg.Done(); movies = serv.searchMovies(query) }()
	go func() { defer wg.Done(); music = serv.searchMusic(query) }()
	wg.Wait()

	results := make([]Result, 0, len(books)+len(movies)+len(music))
	results = append(results, books...)
	results = append(results, movies...)
	results = append(results, music...)
	return results
}

// getJSON performs a GET and decodes the JSON body into out. Returns an error
// on transport failure, non-2xx status, or decode failure.
func getJSON(reqURL string, out interface{}) error {
	res, err := client.Get(reqURL)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return fmt.Errorf("unexpected status %d from %s", res.StatusCode, reqURL)
	}
	return json.NewDecoder(res.Body).Decode(out)
}

// --- Books: Open Library (keyless) ---

func (serv *Service) searchBooks(query string) []Result {
	reqURL := "https://openlibrary.org/search.json?limit=5&fields=title,author_name,cover_i,key,first_publish_year&title=" + url.QueryEscape(query)

	var payload struct {
		Docs []struct {
			Title            string   `json:"title"`
			AuthorName       []string `json:"author_name"`
			CoverI           int      `json:"cover_i"`
			Key              string   `json:"key"`
			FirstPublishYear int      `json:"first_publish_year"`
		} `json:"docs"`
	}
	if err := getJSON(reqURL, &payload); err != nil {
		return []Result{}
	}

	results := make([]Result, 0, len(payload.Docs))
	for _, doc := range payload.Docs {
		pic := ""
		if doc.CoverI != 0 {
			pic = fmt.Sprintf("https://covers.openlibrary.org/b/id/%d-M.jpg", doc.CoverI)
		}
		desc := strings.Join(doc.AuthorName, ", ")
		if doc.FirstPublishYear != 0 {
			if desc != "" {
				desc += " "
			}
			desc += fmt.Sprintf("(%d)", doc.FirstPublishYear)
		}
		results = append(results, Result{
			Category:    "book",
			Name:        doc.Title,
			URL:         "https://openlibrary.org" + doc.Key,
			PicURL:      pic,
			Description: desc,
		})
	}
	return results
}

// --- Movies & TV: TMDB (free API key) ---

func (serv *Service) searchMovies(query string) []Result {
	if serv.tmdbKey == "" {
		return []Result{}
	}
	reqURL := "https://api.themoviedb.org/3/search/multi?include_adult=false&page=1&query=" +
		url.QueryEscape(query) + "&api_key=" + url.QueryEscape(serv.tmdbKey)

	var payload struct {
		Results []struct {
			ID           int64  `json:"id"`
			MediaType    string `json:"media_type"`
			Title        string `json:"title"`         // movies
			Name         string `json:"name"`          // tv
			Overview     string `json:"overview"`
			PosterPath   string `json:"poster_path"`
			ReleaseDate  string `json:"release_date"`
			FirstAirDate string `json:"first_air_date"`
		} `json:"results"`
	}
	if err := getJSON(reqURL, &payload); err != nil {
		return []Result{}
	}

	results := make([]Result, 0, len(payload.Results))
	for _, r := range payload.Results {
		if r.MediaType == "person" {
			continue
		}
		name := r.Title
		if name == "" {
			name = r.Name
		}
		if name == "" {
			continue
		}
		pic := ""
		if r.PosterPath != "" {
			pic = "https://image.tmdb.org/t/p/w500" + r.PosterPath
		}
		results = append(results, Result{
			Category:    "movie",
			Name:        name,
			URL:         fmt.Sprintf("https://www.themoviedb.org/%s/%d", r.MediaType, r.ID),
			PicURL:      pic,
			Description: r.Overview,
		})
		if len(results) >= 5 {
			break
		}
	}
	return results
}

// --- Music: iTunes Search API (keyless) ---

func (serv *Service) searchMusic(query string) []Result {
	reqURL := "https://itunes.apple.com/search?media=music&entity=album&limit=5&term=" + url.QueryEscape(query)

	var payload struct {
		Results []struct {
			ArtistName        string `json:"artistName"`
			CollectionName    string `json:"collectionName"`
			CollectionViewURL string `json:"collectionViewUrl"`
			ArtworkURL100     string `json:"artworkUrl100"`
			PrimaryGenreName  string `json:"primaryGenreName"`
		} `json:"results"`
	}
	if err := getJSON(reqURL, &payload); err != nil {
		return []Result{}
	}

	results := make([]Result, 0, len(payload.Results))
	for _, r := range payload.Results {
		name := r.CollectionName
		if r.ArtistName != "" {
			name = r.ArtistName + " – " + r.CollectionName
		}
		// Upsize the thumbnail from the default 100x100 to a sharper cover.
		pic := strings.Replace(r.ArtworkURL100, "100x100bb", "600x600bb", 1)
		results = append(results, Result{
			Category:    "music",
			Name:        name,
			URL:         r.CollectionViewURL,
			PicURL:      pic,
			Description: r.PrimaryGenreName,
		})
	}
	return results
}
