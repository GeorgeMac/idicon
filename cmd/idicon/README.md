`idicon` command line tool
=========================

The command is mostly just a proof of concept.
It doesn’t currently support all the configurable parameters provided by the library.

### Printing to terminal

For example `idicon print GeorgeMac` produces an ASCII representation like the following:

```
-++++-
++++++
------
++--++
++--++
--++--
```

You can alter the dimensions using the `-w` and `-h` options for width and height respectively.

e.g. `idicon print -w 10 -h 10 GeorgeMac` gives:

```
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

### SVG

`idicon svg -w 5 -h 5 GeorgeMac`

This command will produce something similar to the following:

```xml
<svg width="250" height="250">
    <rect x="0" y="0" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="50" y="0" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="100" y="0" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="150" y="0" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="200" y="0" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="0" y="50" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="50" y="50" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="100" y="50" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="150" y="50" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="200" y="50" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="0" y="100" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="50" y="100" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="100" y="100" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="150" y="100" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="200" y="100" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="0" y="150" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="50" y="150" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="100" y="150" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="150" y="150" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="200" y="150" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="0" y="200" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="50" y="200" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="100" y="200" width="50" height="50" style="fill:#76ffc4"></rect>
    <rect x="150" y="200" width="50" height="50" style="fill:#f28074"></rect>
    <rect x="200" y="200" width="50" height="50" style="fill:#76ffc4"></rect>
</svg>
```

### HTML

`idicon html -w 5 -h 5 GeorgeMac`

Same as above but wrapped in `<html><body></body></html>`!
