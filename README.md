# Chapter 05 - Htmx Patterns

Go implementation of the [example Contact App][0] code in Chapter 05 of the [Hypermedia Systems][1] book.

> [!NOTE]  
> See the [main][2] branch for an index of all implemented chapters.

## Pagination Examples

Chapter 5 illustrates three pagination techniques. The uncommented code in [index.gohtml](templates/contacts/index.gohtml) runs the "Infinite Scroll" implementation.  To really appreciate the technique, the browser window should be made very small or the contacts in [contacts.json](contacts.json) should be increased to the point where all contacts will not fit in the browser window.

The "Basic Pagination" and "Load More Button" techniques are implemented in [index.gohtml](templates/contacts/index.gohtml. To enable them, simply comment out the "Infinite Scroll" code and uncomment the code for the desired technique.

## Running

The following command will start the server that runs the application.

```shell
go run ./cmd/main.go 
```
Once the server is running, the app can be accessed in your browser at http://localhost:8080/. Use `ctrl+c` to stop the server.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"
[1]: https://hypermedia.systems/ "Hypermedia Systems book"
[2]: https://github.com/mousedownco/htmx-contact-app "htmx-contact-app main"