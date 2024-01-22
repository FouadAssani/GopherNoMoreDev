![Logo](GopherNoMoreDevLogo.webp)

# GopherNoMoreDev
In Go's realm, where Restful APIs flow, A subtle dance of code, unseen yet bright. Whispers of a mind, in shadows glow, Crafting bridges in the digital night.  With a wink and a smile, the coders weave, A hidden jest in every line they pen. In this ballet of bytes, they dare to believe, In unseen hands guiding them now and then.


Welcome to GoPhernoMoreDev - where we write Go code and sometimes, it even works!

## Table of Contents
- [Getting Started](#getting-started)
- [Building - Not Just for Carpenters](#building---not-just-for-carpenters)
- [Testing - Because We Care](#testing---because-we-care)
- [Formatting - Not Just Your College Essays](#formatting---not-just-your-college-essays)
- [Dependencies - We All Have Them](#dependencies---we-all-have-them)
- [Code Generation - Like Magic, But Real](#code-generation---like-magic-but-real)
- [OpenAPI - Because Open Is Better](#openapi---because-open-is-better)
- [Docker - Ship It Like FedEx](#docker---ship-it-like-fedex)
- [Cleanup - Not Just For Your Room](#cleanup---not-just-for-your-room)

## Getting Started

First, ensure Go is installed. It's like the foundation of your house; without it, everything falls apart. Download it from [here](https://golang.org/dl/) if you haven't already.

## Building - Not Just for Carpenters

Run the following command and watch the magic happen:

```bash
make build
```

This will build the entire project, including the parts you forgot existed.

## Testing - Because We Care

We have tests! Yes, believe it or not, we like to check if our code actually works.

- For unit tests, execute:
  ```bash
  make test
  ```

- Integration tests (because more testing equals more fun):
  ```bash
  make test-integration
  ```

- End-to-End Tests, for the ultimate "will it blend?" experience:
  ```bash
  make test-e2e
  ```

## Formatting - Not Just Your College Essays

Our code is pretty, and yours should be too:

```bash
make format
make lint
```

It's like spellcheck, but for code.

## Dependencies - We All Have Them

To install our plethora of dependencies:

```bash
make deps
```

To update them (because staying current is cool):

```bash
make tidy
```

## Code Generation - Like Magic, But Real

For when you're too lazy to write code:

```bash
make generate
```

## OpenAPI - Because Open Is Better

For the fans of OpenAPI:

```bash
make openapi
```

## Docker - Ship It Like FedEx

To containerize the application (because containers are hip):

- Build it:
  ```bash
  make docker-build
  ```

- Run it with Docker Compose (because why do it manually?):
  ```bash
  make docker-compose-run
  ```

## Cleanup - Not Just For Your Room

To clean up the mess (i.e., binaries, caches, etc.):

```bash
make clean
```

## Join the Fun

Contributions are welcome - just make sure your code is more functional than funny.

Remember, a good README is like a good joke - if you have to explain it, it’s not that good. Happy coding!
