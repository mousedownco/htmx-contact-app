Hypermedia Systems - Contact App
================================

# A Go Implementation

This repository contains a Go implementation of the example [`Contacts App` project][0] used in the [Hypermedia Systems book][1].

Example code for select chapters are available in separate branches:

 * Branch [Ch03-A_Web_1.0_Application](../../tree/Ch03-A_Web_1.0_Application) - implements [Chapter 03 - A Web 1.0 Application][3]
 * Branch [Ch05-Htmx_Patterns](../../tree/Ch05-Htmx_Patterns) - implements [Chapter 05 - Htmx Patterns][5]
 * Branch [Ch06-More_Htmx_Patterns](../../tree/Ch06-More_Htmx_Patterns) - implements [Chapter 06 - More Htmx Patterns][6]

# Implementation Notes

This isn't the most idiomatic Go project as it tries to align itself as closely as possible with the [book's example project][0].

The view elements are based on [Jon Calhoun's article "Creating the V in MVC"][100]. Routing and "Flash" messages are implemented using [Gorilla Mux][101] and [Gorilla Sessions][102] respectively.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"
[1]: https://hypermedia.systems/ "Hypermedia Systems book"
[3]: https://hypermedia.systems/a-web-1-0-application/ "Chapter 03 - A Web 1.0 Application"
[5]: https://hypermedia.systems/htmx-in-action/ "Chapter 05 - Htmx Patterns"
[6]: https://hypermedia.systems/more-htmx-patterns "Chapter 06 - More Htmx Patterns"
[100]: https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/ "Creating the V in MVC"
[101]: https://github.com/gorilla/mux "Gorilla Mux"
[102]: https://github.com/gorilla/sessions "Gorilla Sessions"