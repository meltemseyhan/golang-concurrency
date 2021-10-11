package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

func (b Book) String() string {
	return fmt.Sprintf("ID:\t\t%d\nTitle:\t\t%q\nAuthor:\t\t%q\nPublished:\t%v\n", b.ID, b.Title, b.Author, b.YearPublished)

}

var books = []Book{
	Book{
		ID:            1,
		Title:         "The Hitchhiker's Guide to the Galaxy",
		Author:        "Dougles Adams",
		YearPublished: 1979,
	},
	Book{
		ID:            2,
		Title:         "The Hobbit",
		Author:        "J.R.R. Tolkien",
		YearPublished: 1937,
	},
	Book{
		ID:            3,
		Title:         "A tale of Two Cities",
		Author:        "Charles Dickens",
		YearPublished: 1859,
	},
	Book{
		ID:            4,
		Title:         "Adı: Aylin",
		Author:        "Ayşe Kulin",
		YearPublished: 1997,
	},

	Book{
		ID:            5,
		Title:         "Geniş Zamanlar",
		Author:        "Ayşe Kulin",
		YearPublished: 1998,
	},

	Book{
		ID:            6,
		Title:         "Foto Sabah Resimleri",
		Author:        "Ayşe Kulin",
		YearPublished: 1998,
	},

	Book{
		ID:            7,
		Title:         "Sevdalinka",
		Author:        "Ayşe Kulin",
		YearPublished: 1999,
	},

	Book{
		ID:            8,
		Title:         "Köprü",
		Author:        "Ayşe Kulin",
		YearPublished: 2001,
	},

	Book{
		ID:            9,
		Title:         "Nefes Nefese",
		Author:        "Ayşe Kulin",
		YearPublished: 2002,
	},

	Book{
		ID:            10,
		Title:         "Gece Sesleri",
		Author:        "Ayşe Kulin",
		YearPublished: 2004,
	},
}
