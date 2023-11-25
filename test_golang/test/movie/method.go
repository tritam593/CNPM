package movie
import (
	"encoding/json"

	"github.com/gorilla/mux"

	"math/rand"
	"net/http"
	"strconv"
)

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}
var movies []Movie

var Movies = movies
func GetMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func DeleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	param := mux.Vars(r)
	for index, item := range movies {

		if item.ID  == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(movies)
}

func GetMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for _, item := range movies {

		if item.ID  == param["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
}

func CreateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var movie Movie
	_ = json.NewDecoder(r.Body).Decode(&movie)
	movie.ID = strconv.Itoa(rand.Intn(100000000))
	movies = append(movies, movie)

	json.NewEncoder(w).Encode(movie)
}

func UpdateMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	param := mux.Vars(r)
	for index, item := range movies {
		if item.ID  == param["id"] {
			movies = append(movies[:index], movies[index+1:]...)

			var movie Movie
			_ = json.NewDecoder(r.Body).Decode(&movie)
			movie.ID = param["id"]
			movies = append(movies, movie)
			json.NewEncoder(w).Encode(movie)
			return
		}
	}
}

func Initial(){
	movies = append(movies, Movie{ID: "1", Isbn: "12345", Title: "Movie one",
		Director: &Director{FirstName: "Tam", LastName: "Tri"}})

	movies = append(movies, Movie{ID: "2", Isbn: "22345", Title: "Movie two",
		Director: &Director{FirstName: "Tam", LastName: "Tri"}})
}

// var movies  = movie.Movies


// func main() {
// 	r := mux.NewRouter()
// 	movie.Initial()
// 	r.HandleFunc("/movies", movie.GetMovies).Methods("GET")
// 	r.HandleFunc("/movies/{id}", movie.GetMovie).Methods("GET")
// 	r.HandleFunc("/movies", movie.CreateMovie).Methods("POST")
// 	r.HandleFunc("/movies/{id}", movie.UpdateMovie).Methods("PUT")
// 	r.HandleFunc("/movies/{id}", movie.DeleteMovie).Methods("DELETE")

// 	fmt.Println("Starting at port 127.0.0.1:8000")
// 	log.Fatal(http.ListenAndServe("127.0.0.1:8000", r))

// }