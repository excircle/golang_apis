# RESTful API Design Examples In Golang
Repository for Examples of Golang APIs. These examples are designed to be cloned and executed on a terminal and used in conjunction with a web browser.

## Table of Contents

#### API & HTTP Fundamentals

<details><summary>API Dictionary</summary>
<p>

A list of helpful terms regarding APIs

### Acronyms

<b>SPA</b> - Single Page Application </br>
<b>SOA</b> - Service-Oriented Architecture </br>
<b>ROA</b> - Response-Oriented Architecture </br>
<b>OAS</b> - OpenAPI Specification</br>

### Key Words

<b>Multiplexer</b> - An entity that matches a given URL to a registered pattern</br>
<b>Middleware</b> - An entity that matches a given URL to a registered pattern</br>
<b>Closure Function</b> - A closure function returns another function</br>
closure function
<p>
</details>

<details><summary>What Is 'Cross-Origin Resource Sharing?'</summary>
<p>

<b>Cross-Origin Resource Sharing (CORS)</b> is method of authorizing a calling client (I.E your code or application) to receive resources from a given host (targeted API).

In general, your browser does not know which external resources are safe and which are not.

CORS allows your browser to determine if a host is safe by checking if it's origin is listed at the API host (server).

An "origin" is the fully-qualified domain name and port of the client that's making the request.

EXAMPLE ORIGIN: `http://app.example.com:8080/`

If your browser can see its origin on the API host, in the form of an HTTP header response bearing it's fully-qualified name, it will consider the resource safe and allow subsequent HTTP methods to succeed.

The chain of events for this security authorization to occur is as follows:

![alt text](https://static.packt-cdn.com/products/9781788294287/graphics/a4a9f6d7-fbed-4b54-870f-e90f5f7a66b6.png "CORS process")

<ol start="1">
  <li>An HTTP 'OPTIONS' call is made to determine if the client's origin is listed at the API</li>
  <li>A HTTP response header including this data is returned</li>
  <li>If the origin is found in the response header, the browser authorizes subsequent HTTP methods to be carried out</li>
</ol>

</p>
</details>

### Small Projects

<b><a href="https://github.com/excircle/golang_apis/tree/master/books_CRUD_API">Books_CRUD_API</a></b> - Simple CRUD API to demonstrate basic functionality of "github.com/gorilla/mux" routing and basic http handling using the "net/http" packge.

<b><a href="https://github.com/excircle/golang_apis/tree/master/web_latency_API">Web_Latency_API</a></b> - A slightly sophisticated API. This API will parse a list of web URLs, record their latency, and return the URL whose latency is the lowest.