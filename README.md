# Contacts.App - Htmx Patterns

This repository contains a Go implementation of the example `Contacts.App` project used in the [Hypermedia Systems book][1].

> [!NOTE]  
> This branch contains the [updates to the Contact App from chapter 6 "More Htmx Patterns"][7].

# Bulk Delete

The book describes a [Bulk Delete][8] feature that uses the the `DELETE` method to send a form encoded request containing the IDs of the contacts to delete. This is not a standards-compliant use of `DELETE` and [Go's `Request.ParseForm()`][9] method will initialize the to an empty value.

The implementation in this repository has been modified to use `hx-post` to `/contacts/delete` instead. The handler is implemented to handle `POST` requests and the remaining functionality aligns with what is described in the book.

# Running

The following command will start the server that runs the `Contacts.App` application.

```shell
go run ./cmd/main.go 
```

Once the server is running, the app can be accessed in your browser at [http://localhost:8080/](http://localhost:8080/). Use `ctrl+c` to stop the server.

# Go References

This isn't the most idiomatic Go project as it tries to align itself as closely as possible with the [book's example project][0].

The view elements are based on [Jon Calhoun's article "Creating the V in MVC"][3]. Routing and "Flash" messages are implemented using [Gorilla Mux][4] and [Gorilla Sessions][5] respectively.

## Why?

After running into problems getting the book's [Python/Flask based Contact App example][0] working, I decided to just re-implement it in Go.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"

[1]: https://hypermedia.systems/ "Hypermedia Systems book"

[2]: https://hypermedia.systems/a-web-1-0-application/ "Chapter 03 - A Web 1.0 Application"

[3]: https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/ "Creating the V in MVC"

[4]: https://github.com/gorilla/mux "Gorilla Mux"

[5]: https://github.com/gorilla/sessions "Gorilla Sessions"

[6]: https://hypermedia.systems/htmx-in-action/ "Chapter 05 - Htmx Patterns"

[7]: https://hypermedia.systems/more-htmx-patterns "Chapter 06 - More Htmx Patterns"

[8]: https://hypermedia.systems/more-htmx-patterns/#_bulk_delete "Bulk Delete"

[9]: https://pkg.go.dev/net/http#Request.ParseForm "net/http - Request.ParseForm()"