# Contacts.App - Web 1.0

This repository contains a Go implementation of the example `Contacts.App` project used in the [Hypermedia Systems book][1].

> [!NOTE]  
> This branch contains the [Web 1.0 Application from Chapter 3][2] of the book.

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
