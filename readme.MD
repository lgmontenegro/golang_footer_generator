#Footer Pagination Exercise#

This exercise consists in to create the base for a footer where you inform the current page, the number of pages, the number of pages on each side of current pages and the number of pages at the footer's border.
"..." will be used to fulfill the lack of pages

Ex: 
Current page is 4
Total of 10 pages
2 pages surrounding the current page each side
A border with 2 pages in the beginning and the end

So, the footer will be:
"1 2 3 4 5 6 ... 9 10"

Another example:
Current page is 100 of 1000 pages
no pages surrounding the current page
4 pages in the borders, each side

We will get "1 2 3 4 ... 100 ... 997 998 999 1000"

To test it:
1. Download using: ```go get github.com/lgmontenegro/golang_footer_generator```
2. Get do directory (maybe ```<go worksapce>/src/github/lgmontenegro/golang_footer_generator```)
3. Run ```go test```
4. It's Done!

Enjoy editing the ```github.com/lgmontenegro/golang_footer_generator/footer_test.go``` file!

Thank you for the opportunity