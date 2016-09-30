# Pass Management System

### @hooph00p

## Usage

#### Requirements

- The latest version of [Go](golang.org/dl).
- A recent version of [Git](https://git-scm.com/downloads).

To get a copy of the source code, clone this repository into your `$GOPATH`.

If you have go installed, it's _much_ easier to call `go get`:

```
$ go get github.com/hooph00p/transport
```

From there, you'll need to install the dependencies -- I've included a Makefile so you can just `make` it in your directory.

From there, you can install and run the command:

```
$ go install
$ transport
```

Build, and run the executable:

```
$ go build
$ ./transport
```

...or call `go run` from the working directory.

```
$ go run main.go
```


If ports `:8080` and `:8081` are free on your machine, you should see be able to open up `http://localhost:8080` to see my web application.


## Approach

#### Three Sprints

## 1. Vertical Development

**Goal: Allow for users to login and create prepaid passes.**

- [x] Create Views
  - [x] Login
  - [x] Account Page

- [x] Show passes
- [x] Allow Creation of Passes

#### User Stories:

- [x] Joe wants to register.
- [x] Joe wants to log in.
- [x] Joe wants to add a monthly commuter pass.
- [x] Joe wants to see how many passes he has.
- [ ] Joe wants to add funds to a prepaid pass.

## 2. Lateral Development

**Goal: Allow for users to login and manage all types of passes and discounts.**

- [ ] Add Passwords
- [ ] Add Remaining Discounts
- [ ] Add Remaining Pass Types
- [ ] Add Usages

#### User Stories:

- Joe wants to add a monthly pass for the Rail.
- Joe wants to add money to his Subway pass.
- Joe wants to use the pass for the Rail.

## 3. Polish

**Goal: Clean everything up and prepare for presentation to client.**

- [ ] Tackle Backlog
- [ ] Expand Testing
- [ ] Expand User Stories
- [ ] Brand User Interface
