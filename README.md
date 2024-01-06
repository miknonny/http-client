# http-client

A simple and lightweight HTTP client implemented using go.

Why are we creating our own httpClient.

- ease of doing telemetry.
- ownership and proper understanding of code. nothing can break our code.
- changes are made only in one place which is our library.
- Not having to deal with closing connections. as the library closes open connections.
- we can use the httpClient as the base for another api.
- only one instance of a httpclient exists in our entire code base. we can not run out of memery fast.

- to check code coverage for our library - go test ./... -cover

Main.Minor.Patch
Main - Breaking changes to already existing api. if 0 it means its not stable.
Minor- added more api but existing api works as usual.
Patch - fixed existing api or modified implementation details of existing api.

\ Master Branch - this branch contains our stable code.
$ git checkout -b feature/httpClient - this contains features that we are adding to our branch.

$ git add .

$ git commit -m "initial version"

$ git push -u origin feature/httpClient // this indicates we are adding a feature to our httpClient lib.

$ git tag v0.1.0

$ git tag // displays the current tag of our commit

$ git push origin v0.1.0
