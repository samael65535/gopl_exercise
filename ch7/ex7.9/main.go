package main
/*
练习 7.9：
使用html/template包 (§4.6) 替代printTracks将tracks展示成一个HTML表格。将这个解决方案用在前一个练习中，让每次点击一个列的头部产生一个HTTP请求来排序这个表格。
*/
import (
	"time"
	"sort"
	"html/template"
	"log"
	"net/http"
	"io"
)

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

var tracks = []*Track{
	{"Ready 2 Go", "Martin Solveig", "Smash", 2011, length("4m24s")},
	{"Go", "Delilah", "From the Roots Up", 2012, length("3m38s")},
	{"Go", "Moby", "Moby", 1992, length("3m37s")},
	{"Go Ahead", "Alicia Keys", "As I Am", 2007, length("4m36s")},
}


func length(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		panic(s)
	}
	return d
}


type byArtist []*Track
func (x byArtist) Len() int           { return len(x) }
func (x byArtist) Less(i, j int) bool { return x[i].Artist < x[j].Artist }
func (x byArtist) Swap(i, j int)      { x[i], x[j] = x[j], x[i] }


const templ = `
<table>
<tr style='text-align: left' >
  <th><a href='?sort=title'>Title</a></th>
  <th><a href='?sort=artist'>Artist</a></th>
  <th><a href='?sort=album'>Album</a></th>
  <th><a href='?sort=year'>Year</a></th>
  <th><a href='?sort=length'>Length</a></th>
</tr>
{{range .}}
<tr style='text-align: left'>
  <td style="padding-right: 30px;">{{.Title}}</td>
  <td style="padding-right: 30px;">{{.Artist}}</td>
  <td style="padding-right: 30px;">{{.Album}}</td>
  <td style="padding-right: 30px;">{{.Year}}</td>
  <td>{{.Length}}</td>
</tr>
{{end}}
</table>
`

func printTracks(tracks []*Track, w io.Writer) {
	result := template.Must(template.New("tracklist").Parse(templ))
	if err := result.Execute(w, tracks); err != nil {
		log.Fatal(err)
	}
}


func main() {
	http.HandleFunc("/", func (w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		var compFunc func(x, y *Track) bool
		var cs customSort

		if len(r.Form["sort"]) == 0 {
			cs = customSort{tracks, compFunc}
		}

		for _, sortBy := range r.Form["sort"] {
			switch sortBy {
			case "title":
				compFunc = compTitle
				break;
			case "artist":
				compFunc = compArtist
				break;
			case "album":
				compFunc = compAlbum
				break;
			case "year":
				compFunc = compYear
				break;
			case "length":
				compFunc = compLength
				break;
			default:
				compFunc = compTitle
				break;
			}
			cs = customSort{tracks, compFunc}
			sort.Stable(cs)
		}
		printTracks(tracks, w)
	})
	http.ListenAndServe("localhost:8001", nil)
}

type customSort struct {
	t    []*Track
	less func(x, y *Track) bool
}

func (x customSort) Len() int { return len(x.t)}
func (x customSort) Less(i, j int) bool { return x.less(x.t[i], x.t[j]) }
func (x customSort) Swap(i, j int)      { x.t[i], x.t[j] = x.t[j], x.t[i] }


func compTitle(x, y *Track) bool{
	return x.Title < y.Title
}

func compAlbum(x, y *Track) bool{
	return x.Album < y.Album
}

func compLength(x, y *Track) bool {
	return x.Length.Seconds() < y.Length.Seconds()
}

func compYear(x, y *Track) bool {
	return x.Year < y.Year
}

func compArtist(x, y *Track) bool {
	return x.Artist < y.Artist
}
