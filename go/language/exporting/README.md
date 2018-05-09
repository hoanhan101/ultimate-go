# Encapsulation

In Object oriented programming, we are used to public, private, protected type of access
mechanism. Go is different. Everything in Go is about package.

Package is a self contained unit of code. Every folder in our source tree is a self-contained
user code. We will get deeper into that in package oriented design section. For now, if we are
thinking about a package being that self containing code, firewall that separates code  and
having language supported for it then we can think of encapsulation being associated with the
package itself.

The idea is: anything that is named in a given package can be exported or accessible through
other packages or unexported or not accessible through other packages.
