package config

// "fmt"
type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.
func albumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	// var albums []Album

	// rows, err := db_conn.Query("SELECT * FROM album WHERE artist = ?", name)
	// if err != nil {
	//     return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	// }
	// defer rows.Close()
	// // Loop through rows, using Scan to assign column data to struct fields.
	// for rows.Next() {
	//     var alb Album
	//     if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
	//         return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	//     }
	//     albums = append(albums, alb)
	// }
	// if err := rows.Err(); err != nil {
	//     return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	// }
	var albums []Album
	alb := Album{1, "Hello", "World", 1.00}
	albums = append(albums, alb)
	return albums, nil
}
