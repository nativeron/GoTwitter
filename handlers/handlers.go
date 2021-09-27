package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/nativeron/GoTwitter/middlew"
	"github.com/nativeron/GoTwitter/routers"
	"github.com/rs/cors"
)

/*seteo mi puerto, el handler y pongo a escuchar al servidor*/
func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlew.CheckDB(routers.Register)).Methods("POST")
	router.HandleFunc("/login", middlew.CheckDB(routers.Login)).Methods("POST")
	router.HandleFunc("/seeprofile", middlew.CheckDB(middlew.ValidateJWT(routers.SeeProfile))).Methods("GET")
	router.HandleFunc("/editprofile", middlew.CheckDB(middlew.ValidateJWT(routers.ModifProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlew.CheckDB(middlew.ValidateJWT(routers.PublishTweet))).Methods("POST")
	router.HandleFunc("/gettweets", middlew.CheckDB(middlew.ValidateJWT(routers.GetAllTweets))).Methods("GET")
	router.HandleFunc("/deletetweet", middlew.CheckDB(middlew.ValidateJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadavatar", middlew.CheckDB(middlew.ValidateJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/uploadbanner", middlew.CheckDB(middlew.ValidateJWT(routers.UploadBanner))).Methods("POST")

	router.HandleFunc("/getavatar", middlew.CheckDB(middlew.ValidateJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/getbanner", middlew.CheckDB(middlew.ValidateJWT(routers.GetBanner))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
