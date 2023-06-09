package db

import (
	"database/sql"
	"fmt"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

// albumsByArtist queries for albums that have the specified artist name.
func AlbumsAll() ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// Parameterized query
	rows, err := db_conn.Query("SELECT * FROM album")
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist: %v", err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		// Scan automatically converts values from database columns into appropriate Go types based on the destination variables
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArtist: %v", err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist: %v", err)
	}

	return albums, nil
}

// albumsByArtist queries for albums that have the specified artist name.
func AlbumsByArtist(name string) ([]Album, error) {
	// An albums slice to hold data from returned rows.
	var albums []Album

	// Parameterized query
	rows, err := db_conn.Query("SELECT * FROM album WHERE artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}
	defer rows.Close()

	// Loop through rows, using Scan to assign column data to struct fields.
	for rows.Next() {
		var alb Album
		// Scan automatically converts values from database columns into appropriate Go types based on the destination variables
		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
		}
		albums = append(albums, alb)
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("AlbumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

// albumByID queries for the album with the specified ID.
func AlbumByID(id int64) (Album, error) {
	// An album to hold data from the returned row.
	var alb Album

	row := db_conn.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

// AddAlbum adds the specified album to the database,
// returning the album ID of the new entry
func AddAlbum(alb Album) (int64, error) {
	result, err := db_conn.Exec("INSERT INTO album (title, artist, price) VALUES (?, ?, ?)", alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("addAlbum: %v", err)
	}
	return id, nil
}

// UpdateAlbumByID updates the specified album in the database by its ID
func UpdateAlbumByID(id int64, alb Album) (int64, error) {
	var album Album
	row := db_conn.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("albumsById %d: no such album", id)
		}
		fmt.Printf("albumsById %d: %v", id, err)
	}
	fmt.Printf("originalRow: %v\n", album)

	result, err := db_conn.Exec("UPDATE album SET title = ?, artist = ?, price = ? WHERE id = ?", alb.Title, alb.Artist, alb.Price, id)
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("updateAlbum: %v", err)
	}
	fmt.Printf("newRow: %v\n", alb)
	fmt.Printf("rowsAffected: %v\n", rowsAffected)

	return id, nil
}

// DeleteAlbumByID updates the specified album in the database by its ID
func DeleteAlbumByID(id int64) (int64, error) {
	var album Album
	row := db_conn.QueryRow("SELECT * FROM album WHERE id = ?", id)
	if err := row.Scan(&album.ID, &album.Title, &album.Artist, &album.Price); err != nil {
		if err == sql.ErrNoRows {
			fmt.Printf("albumsById %d: no such album", id)
		}
		fmt.Printf("albumsById %d: %v", id, err)
	}
	fmt.Printf("originalRow: %v\n", album)

	result, err := db_conn.Exec("DELETE FROM album WHERE id = ?", id)
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return 0, fmt.Errorf("deleteAlbum: %v", err)
	}
	fmt.Printf("rowsAffected: %v\n", rowsAffected)

	return id, nil
}
