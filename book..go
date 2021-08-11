package main

import "fmt"

type Book struct {
	ID            int
	Title         string
	Author        string
	YearPublished int
}

//func (book Book) String() string {
//	return fmt.Sprintf(
//		"Title:\t\t%q\n"+
//			"Author:\t\t%q\n"+
//			"Published\t\t%v\n", book.Title, book.Author, book.YearPublished)
//}

func (book Book) String() string {
	return fmt.Sprintf(
		"\tTitle:\t%q\tID:%v\n", book.Title, book.ID)
}

var books = []Book{
	Book{
		ID:            1,
		Title:         "Harry Potter 1",
		Author:        "J K Rownling",
		YearPublished: 2001,
	},
	Book{
		ID:            2,
		Title:         "Harry Potter 2",
		Author:        "J K Rownling",
		YearPublished: 2004,
	},
	Book{
		ID:            3,
		Title:         "Harry Potter 3",
		Author:        "J K Rownling",
		YearPublished: 2006,
	},
	Book{
		ID:            4,
		Title:         "Harry Potter 4",
		Author:        "J K Rownling",
		YearPublished: 2008,
	},
	Book{
		ID:            5,
		Title:         "Harry Potter 5",
		Author:        "J K Rownling",
		YearPublished: 2010,
	},
	Book{
		ID:            6,
		Title:         "Harry Potter 6",
		Author:        "J K Rownling",
		YearPublished: 2012,
	},
	Book{
		ID:            7,
		Title:         "Harry Potter 7",
		Author:        "J K Rownling",
		YearPublished: 2014,
	},
}
