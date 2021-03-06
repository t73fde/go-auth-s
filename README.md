# go-auth-s
Simple HTTP Basic Authentication Web Server.

Provides a simple web server that acts as an authentication server, mostly for
development purposes. This is the Go version. There is also a Python version.

To avoid re-implementation of user authentication based on user/password
scheme, application could delegate this to a web server using HTTP Basic
Authentication. This abstracts away authentications details. Application just
need an appropriate URI to configure authentication.

This simple web server is for development purposes. It provides a web server
that authenticates all access to any ressource (identified by a path):

* If there is no user name and password given, it returns 401 (Authentication
  required).
* If the user name or the password is an empty string, it returns 403
  (Forbidden).
* If the user name starts with a `x`, it returns 403 (Forbidden).
* If the user name starts with a `q` and the password is different to the user
  name, it returns 403 (Forbidden).
* Otherwise it returns 200 (OK)

The server listens on port 9876. This can changed by using the `-p` command
line parameter, eg. `./go-auth-s -p 1234` for listening on port 1234.

My students are encouraged to use this server. A simple motivation is given by
my blog post [HTTP Basic Auth als Infrastruktur zur
Authentifizierung](https://t73f.de/blog/2019/basic_auth_infrastruktur/) (in
german).
