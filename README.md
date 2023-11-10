# Chapter 06 - More Htmx Patterns

Go implementation of the [example Contact App][0] code in Chapter 06 of the [Hypermedia Systems][1] book.

> [!NOTE]  
> See the [main][2] branch for an index of all implemented chapters.

# Bulk Delete

The book describes a [Bulk Delete][8] feature that uses the the `DELETE` method to send a form encoded request containing the IDs of the contacts to delete. This is not a standards-compliant use of `DELETE`. [Go's `Request.ParseForm()`][9] method will initialize the form to an empty value, so creating a similar implementation is not feasible using the standard Go libraries.

The implementation in this repository has been modified to use `hx-post` with `/contacts/delete` instead. The handler is implemented to handle `POST` requests and the remaining functionality aligns with what is described in the book.

# Running

The following command will start the server that runs the `Contacts.App` application.

```shell
go run ./cmd/main.go 
```

Once the server is running, the app can be accessed in your browser at [http://localhost:8080/](http://localhost:8080/). Use `ctrl+c` to stop the server.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"
[1]: https://hypermedia.systems/ "Hypermedia Systems book"
[2]: https://github.com/mousedownco/htmx-contact-app "htmx-contact-app main"
[8]: https://hypermedia.systems/more-htmx-patterns/#_bulk_delete "Bulk Delete"
[9]: https://pkg.go.dev/net/http#Request.ParseForm "net/http - Request.ParseForm()"