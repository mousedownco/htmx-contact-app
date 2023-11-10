# Chapter 07 - A Dynamic Archive UI

Go implementation of the [example Contact App][0] code in Chapter 07 of the [Hypermedia Systems][1] book.

> [!NOTE]  
> See the [main][2] branch for an index of all implemented chapters.

# Auto-Download Implementation

The Auto-Download functionality described in the last section of this chapter is commented out in the [archive_ui.gohtml](templates/contacts/archive_ui.gohtml) file. This is the only use of the `_hyperscript` for this chapter.

# Running

The following command will start the server that runs the `Contacts.App` application.

```shell
go run ./cmd/main.go 
```

Once the server is running, the app can be accessed in your browser at [http://localhost:8080/](http://localhost:8080/). Use `ctrl+c` to stop the server.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"
[1]: https://hypermedia.systems/ "Hypermedia Systems book"
[2]: https://github.com/mousedownco/htmx-contact-app "htmx-contact-app main"