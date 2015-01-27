# [Resoursea](http://resoursea.com)

A high productivity web framework for quickly writing resource based services fully implementing the REST architectural style.

## Printer

This package is a tool that prints all the routed Resources and Methods from a Resoursea project.

## How to use

Install the Printer package:

~~~
go get github.com/resoursea/printer
~~~

Routing all your Resource tree using Resoursea is as easy as:

~~~ go
router, err := api.NewRouter(Resource{})
~~~

To print all te routed Resources and Methods:

- Import this package `import "github.com/resoursea/printer"`

- Call the method `printer.Router(router)`

## Larn More

[The concept, Samples, Documentation, interfaces and resources to use...](http://resoursea.com)