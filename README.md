```
NAME
  incat - replaces TOKEN in TEMPLATE with STDIN

SYNOPSIS
  incat -token TOKEN -template ""

DESCRIPTION
  It is sometimes desirable to combine two files by inserting one into the middle of another.

OPTIONS
  -template string
        the name of the template file
  -token string
        the string being replaced (default "TOKEN")

EXAMPLE
  $ cat giant.txt | incat -template small.tmpl > combined.txt
```

To see a *very* simple demo:

```
$ cd example && . ./demo.sh
```


