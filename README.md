# Transliteration of Cyrillic

This is a simple transliteration of Russian (and only Russian) Cyrillic to Latin. I use [Wikipedia's transliteration table](https://en.wikipedia.org/wiki/Scientific_transliteration_of_Cyrillic).

NOTE: Full caps words are not supported.

## Installation

```bash
$ go install github.com/hararudoka/lat@latest
```

## Usage

Using  just command:

```bash
$ lat "Привет, мир!"
Privet, mir!
```

Using pipe:

```bash
$ echo "Привет, мир!" | lat
Privet, mir!
```

Using pipe with output to file:

```bash
$ echo "Привет, мир!" | lat > hello.txt
```

Also, you can clone (and edit!) and build this program by your self, if you want to.
