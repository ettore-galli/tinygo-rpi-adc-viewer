# UNIT TESTING

## 0. Concept

- Tinygo test does not work
- Hardware code is not (yet) tested
- Only logic is tested
- Machine or other hardware libraries are needed for the tests to compile
- They are replaced by a dummy
- The idea is [./create-extensible-goroot.sh]:
    > Create a dummy goroot that links to go root
    > Once, write a dummy replacement for - e.g. - machine
    > Add dummy libraries to the dummy go root
- Run the tests

## 1. Create an extensible go root

```bash
./create-extensible-goroot.sh
```

## 2. Customize dummy libraries

If needed, customize dummy packages and classes, then add them
to the extensible goroot creation section

## 3. make it the go root

```bash
export GOROOT=/tmp/dummy-go-root
```

## 4. Run the tests

```bash
go test -v ./...
```
