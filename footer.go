package footer

import (
	"bytes"
	"strconv"
)

//Bonduary structure used to help assemble the footer's
//bonduary with the begining of the footer and its ending
type Bonduary struct {
	begin, end []int
}

//Around used to help assemble the footer's main body
//that will be merged with the bonduary
type Around struct {
	pages []int
}

//the main function of this package
func footer(currentPage int, totalPages int, bonduaries int, around int) string {
	aroundCurPag := aroundCurrent(currentPage, around)
	bonduary := bonduary(totalPages, bonduaries)

	if totalPages == 1 {
		return "1"
	}

	return createFooterString(aroundCurPag, bonduary)
}

//footerHiddenPages will fill the blank spaces with "..."
//and convert the pages from []int to []string
func footerHiddenPages(pages []int) []string {
	var hiddenPagesArray []string

	for i := 0; i <= (len(pages) - 2); i++ {
		hiddenPagesArray = append(hiddenPagesArray, strconv.Itoa(pages[i]))
		if (pages[i+1] - pages[i]) > 1 {
			hiddenPagesArray = append(hiddenPagesArray, "...")
		}
	}
	hiddenPagesArray = append(hiddenPagesArray, strconv.Itoa(pages[len(pages)-1]))
	return hiddenPagesArray
}

//createFooterString merge the initial footer's structures into a []int
//and after that generate the string with the pages and the "..." between spaces
func createFooterString(aroundStruct Around, bonduaryStruct Bonduary) string {
	var footerStr bytes.Buffer

	bonduaryStruct.begin = removeRepeated(aroundStruct, bonduaryStruct.begin)
	bonduaryStruct.end = removeRepeated(aroundStruct, bonduaryStruct.end)

	footerArray := append(bonduaryStruct.begin, aroundStruct.pages...)
	if bonduaryStruct.begin[0] > aroundStruct.pages[0] {
		footerArray = append(aroundStruct.pages, bonduaryStruct.begin...)
	}
	footerArray = append(footerArray, bonduaryStruct.end...)

	footerStr.WriteString(footerHiddenPages(footerArray)[0])
	for _, pageNum := range footerHiddenPages(footerArray)[1:] {
		footerStr.WriteString(" ")
		footerStr.WriteString(pageNum)
	}

	return footerStr.String()
}

//removeRepeated removes duplicated itens from a slice in [] format
//returning a slice with unique itens
func removeRepeated(aroundArray Around, bonduary []int) []int {
	var unique []int

	for i := range aroundArray.pages {
		for y := range bonduary {
			if aroundArray.pages[i] == bonduary[y] {
				bonduary[y] = 0
			}
		}
	}

	for _, value := range bonduary {
		if value != 0 {
			unique = append(unique, value)
		}
	}

	return unique
}

//aroundCurrent adds the pages around the current pages,
//left and right sides
func aroundCurrent(currentPage int, around int) Around {
	var returnAround Around

	begin := currentPage - around
	if around > currentPage {
		begin = around - currentPage
	}

	end := currentPage + around

	for i := begin; i <= end; i++ {
		returnAround.pages = append(returnAround.pages, i)
	}

	return returnAround
}

//bonduary function calculates the extremes of the footer string
//in bouth sides, begining and end of it
func bonduary(totalPages int, bonduaries int) Bonduary {
	var bonduaryReturn Bonduary

	for i := 1; i <= bonduaries; i++ {
		bonduaryReturn.begin = append(bonduaryReturn.begin, i)
	}

	for i := (totalPages - bonduaries) + 1; i <= totalPages; i++ {
		bonduaryReturn.end = append(bonduaryReturn.end, i)
	}

	return bonduaryReturn
}
