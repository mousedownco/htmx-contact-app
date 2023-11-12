Hypermedia Systems - Contact App
================================

## A Go Implementation

This repository contains a Go implementation of the example [`Contacts App` project][0] used in the [Hypermedia Systems book][1].

Example code for select chapters are available in separate branches:

 * [Chapter 03][3] - Branch [Ch03-A_Web_1.0_Application](../../tree/Ch03-A_Web_1.0_Application)
 * [Chapter 05][5] - Branch [Ch05-Htmx_Patterns](../../tree/Ch05-Htmx_Patterns)
 * [Chapter 06][6] - Branch [Ch06-More_Htmx_Patterns](../../tree/Ch06-More_Htmx_Patterns)
 * [Chapter 07][7] - Branch [Ch07-A_Dynamic_Archive_UI](../../tree/Ch07-A_Dynamic_Archive_UI)
 * [Chapter 09][9] - Branch [Ch09-Client_Side_Scripting](../../tree/Ch09-Client_Side_Scripting)
 * [Chapter 10][10] - Branch [Ch10-JSON_Data_APIs](../../tree/Ch10-JSON_Data_APIs)

## Implementation Notes

This isn't the most idiomatic Go project as it tries to align itself as closely as possible with the [book's example project][0].

The view elements are based on [Jon Calhoun's article "Creating the V in MVC"][100]. Routing and "Flash" messages are implemented using [Gorilla Mux][101] and [Gorilla Sessions][102] respectively.

[0]: https://github.com/bigskysoftware/contact-app "Contact App"
[1]: https://hypermedia.systems/ "Hypermedia Systems book"
[3]: https://hypermedia.systems/a-web-1-0-application/ "Chapter 03 - A Web 1.0 Application"
[5]: https://hypermedia.systems/htmx-in-action/ "Chapter 05 - Htmx Patterns"
[6]: https://hypermedia.systems/more-htmx-patterns/ "Chapter 06 - More Htmx Patterns"
[7]: https://hypermedia.systems/a-dynamic-archive-ui/ "Chapter 07 - A Dynamic Archive UI"
[9]: https://hypermedia.systems/client-side-scripting/ "Chapter 09 - Client-Side Scripting"
[10]: https://hypermedia.systems/json-data-apis/ "Chapter 10 - JSON Data APIs & Hypermedia-Driven Applications"
[100]: https://www.calhoun.io/intro-to-templates-p4-v-in-mvc/ "Creating the V in MVC"
[101]: https://github.com/gorilla/mux "Gorilla Mux"
[102]: https://github.com/gorilla/sessions "Gorilla Sessions"