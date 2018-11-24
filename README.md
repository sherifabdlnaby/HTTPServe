## HTTP Hello World Server
A very basic HTTP Server using net package TCP listener for the sole purpose of printing Hello World on a browser.

This is an example to study Golang and is by no mean a competitor to the [net/http](https://golang.org/pkg/net/http/) or to any other non-GO Web Servers in the game (e.g Apache or Tomcat) 🤷🏻‍♂️.
and yea it technically doesn't even qualify for being an HTTP server since it doesn't come close to implementing [HTTP Specifications and Drafts](https://www.w3.org/Protocols/Specs.html) 🤦🏻‍♂️

However, It is a good domenstration for using and combining Go's interfaces (e.g `io.Reader/Writer`, `net.Conn`) to build something.