package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main(){
	http.Handle("/",http.FileServer(http.Dir("polymer")))
	http.HandleFunc("/api/film",user)
	fmt.Println("web berjalan dengan port 8221")
	http.ListenAndServe(":8221",nil)
}

type film struct {
	Nama string `json:"nama_film"`
	Rate string `json:"rate_film"`
	Foto string `json:"foto_film"`
	Link string `json:"link_film"`
}

var data_film = []film{
	{
		Nama :"Avatar : The Way Of Water",
		Rate :"8.0/10",
		Foto : "img/gambar1.jpg",
		Link : "https://www.imdb.com/title/tt1630029/",
	},
	{
		Nama :"Pengabdi Setan 2",
		Rate :"7.0/10",
		Foto : "img/gambar2.jpg",
		Link : "https://www.imdb.com/title/tt16915972/",
	},
	{
		Nama :"The Raid 2 : Berandal",
		Rate :"7.9/10",
		Foto : "img/gambar3.png",
		Link : "https://www.imdb.com/title/tt2265171/",
	},
	{
		Nama :"KKN di Desa Penari",
		Rate :"6.0/10",
		Foto : "img/gambar4.png",
		Link : "https://www.imdb.com/title/tt11013610/",
	},
	{
		Nama :"Blank Adam",
		Rate :"6.5/10",
		Foto : "img/dontmissend_credits_3881613.jpg",
		Link : "https://www.imdb.com/title/tt6443346/",
	},
	{
		Nama :"Perfect Strangers",
		Rate :"7.5/10",
		Foto : "img/PERFECT-STRANGER.jpg",
		Link : "https://www.imdb.com/title/tt17497250/",
	},

}

func user(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")

	if r.Method == http.MethodGet{
		result, err := json.Marshal(data_film)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(result)
		return
	}
}
