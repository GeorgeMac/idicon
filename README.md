Identicons in Go
================

Use the following to get the libs:

`go get github.com/GeorgeMac/idicon`

Use the following to get the `idicon` command

`go get github.com/GeorgeMac/idicon/cmd/idicon`

The command is mostly just a proof of concept.
It outputs both a couple of complementary RGBA colours and a printed representation of the identicon.

For example `idicon GeorgeMac` produces the following:

```
Colours [{110 161 255 255}, {255 177 113 255}]
-++++-
++++++
------
++--++
++--++
--++--
```

You can alter the dimensions using the `-w` and `-h` options for width and height respectively.

e.g. `idicon -w 10 -h 10 GeorgeMac` gives:

```
Colours [{110 161 255 255}, {255 177 113 255}]
+-+----+-+
----++----
--+----+--
-+------+-
---++++---
-+------+-
+--------+
++++++++++
-++++++++-
-++----++-
```
