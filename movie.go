package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"path"
	"regexp"
	"strconv"
)

// Movie represent a movie
type Movie struct {
	ID          int      `json:"id,string"`
	Title       string   `json:"title"`
	Description string   `json:"desc"`
	Genre       string   `json:"genre"`
	Artists     []string `json:"artists"`
	Director    string   `json:"director"`
	Rating      float64  `json:"-"`
	ReleaseDate string   `json:",omitempty"`
}

type movieRequest struct {
	Title string `json:"title"`
}

var movies []Movie
var reMovieID *regexp.Regexp

func handleMovies(w http.ResponseWriter, r *http.Request) {
	var req movieRequest
	encoder := json.NewEncoder(w)
	decoder := json.NewDecoder(r.Body)

	err := decoder.Decode(&req)
	if err != nil {
		encoder.Encode(&movies)
		return
	}

	if movieResponse, ok := searchMovie(req.Title); ok {
		encoder.Encode(&movieResponse)
	} else {
		http.Error(w, fmt.Sprintf("Not found movie with title '%s'", req.Title), http.StatusBadRequest)
	}
}

func searchMovie(title string) (*Movie, bool) {
	for _, m := range movies {
		if m.Title == title {
			return &m, true
		}
	}
	return nil, false
}

func handleMovieFromID(w http.ResponseWriter, r *http.Request) {
	values := reMovieID.FindStringSubmatch(r.URL.Path)
	if len(values) < 1 {
		http.Error(w, fmt.Sprintf("Bad request. Not valid ID (%v) in request '%s'", path.Base(r.URL.Path), r.URL.Path), http.StatusBadRequest)
		return
	}
	id, err := strconv.Atoi(values[1])
	if err != nil {
		http.Error(w, fmt.Sprintf("Bad request. Non numeric ID (%v) in request '%s'", values[1], r.URL.Path), http.StatusBadRequest)
		return
	}
	if id <= 0 || id > len(movies) {
		http.Error(w, fmt.Sprintf("Bad request. Out of range ID (%d) in request '%s'", id, r.URL.Path), http.StatusBadRequest)
		return
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(&movies[id-1])
}

func init() {
	reMovieID, _ = regexp.Compile("/movies/([0-9]+)")
	movies = []Movie{
		{
			ID:          1,
			Title:       "Kagemusha",
			Description: "When a warlord dies, a peasant thief is called upon to impersonate him, in order to protect his clan from appearing weak and vulnerable. But he finds himself haunted by the warlord’s spirit as well as his own ambitions.",
			Genre:       "Action & Adventure",
			Artists:     []string{"Jinpachi Nezu", "Kenichi Hagiwara", "Tatsuya Nakadai", "Tsutomu Yamazaki"},
			Director:    "Akira Kurosawa",
			Rating:      10.0,
			ReleaseDate: "1980-04-26",
		},
		{
			ID:          2,
			Title:       "Seven Samurai",
			Description: "A veteran samurai, who has fallen on hard times, answers a village's request for protection from bandits. He gathers 6 other samurai to help him, and they teach the townspeople how to defend themselves, and they supply the samurai with three small meals a day. The film culminates in a giant battle when 40 bandits attack the village.",
			Genre:       "Action & Adventure",
			Artists:     []string{"Bokuzen Hidari", "Daisuke Katô", "Eijirô Tôno", "Isao Kimura", "Kamatari Fujiwara", "Keiko Tsushima", "Kokuten Kôdô", "Minoru Chiaki", "Seiji Miyaguchi", "Takashi Shimura", "Toshirô Mifune", "Yoshio Inaba", "Yoshio Kosugi", "Yoshio Tsuchiya", "Yukiko Shimazaki"},
			Director:    "Akira Kurosawa",
			Rating:      10.0,
			ReleaseDate: "1956-11-19",
		},
		{
			ID:          3,
			Title:       "13 Assassins",
			Description: "Cult director Takashi Miike (Ichi the Killer, Audition) delivers a bravado period action film set at the end of Japan’s feudal era. 13 Assassins - a “masterful exercise in cinematic butchery” (New York Post) - is centered around a group of elite samurai who are secretly enlisted to bring down a sadistic lord to prevent him from ascending to the throne and plunging the country into a war torn future.",
			Genre:       "Action & Adventure",
			Artists:     []string{"Kôji Yakusho", "Takashi Miike", "Yusuke Iseya"},
			Director:    "Takashi Miike",
			Rating:      9.5,
			ReleaseDate: "2011-07-05T07:00:00Z",
		},
		{
			ID:          4,
			Title:       "007: Skyfall",
			Description: "When Bond's latest assignment goes gravely wrong and agents around the world are exposed, MI6 is attacked forcing M to relocate the agency. These events cause her authority and position to be challenged by Gareth Mallory (Ralph Fiennes), the new Chairman of the Intelligence and Security Committee. With MI6 now compromised from both inside and out, M is left with one ally she can trust: Bond. 007 takes to the shadows - aided only by field agent, Eve (Naomie Harris) - following a trail to the mysterious Silva (Javier Bardem), whose lethal and hidden motives have yet to reveal themselves.",
			Genre:       "Action & Adventure",
			Artists:     []string{"Albert Finney", "Ben Whishaw", "Bérénice Marlohe", "Daniel Craig", "Helen McCrory", "Javier Bardem", "Judi Dench", "Naomie Harris", "Ola Rapace", "Ralph Fiennes", "Rory Kinnear", "Tonia Sotiropoulou"},
			Director:    "Sam Mendes",
			Rating:      8.5,
			ReleaseDate: "2012-11-09",
		},
		{
			ID:          5,
			Title:       "The Godfather",
			Description: "The story spans the years from 1945 to 1955 and chronicles the fictional Italian-American Corleone crime family. When organized crime family patriarch Vito Corleone barely survives an attempt on his life, his youngest son, Michael, steps in to take care of the would-be killers, launching a campaign of bloody revenge.",
			Genre:       "Action & Adventure",
			Artists:     []string{"Abe Vigoda", "Al Lettieri", "Al Martino", "Al Pacino", "Diane Keaton", "Gianni Russo", "James Caan", "John Cazale", "John Marley", "Marlon Brando", "Richard Conte", "Richard S. Castellano", "Robert Duvall", "Rudy Bond", "Simonetta Stefanelli", "Sterling Hayden", "Talia Shire", "Tony Giorgio", "Victor Rendina"},
			Director:    "Francis Ford Coppola",
			Rating:      9.5,
			ReleaseDate: "1972-03-24",
		},
	}
}
