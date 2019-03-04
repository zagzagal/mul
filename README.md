# mul
A cli London Mulligan simulator

Installation
```
go get -u github.com/zagzagal/mul
```

```
Usage of ./mul:
Usage: ./mul <deckstring> <logic>

  -deck int
        Deck size (default 60)
  -hand int
        Hand size (default 7)
  -mul int
        Min after muligan hand size (default 5)
  -perm int
        number of trys (default 10000)
  -q    Quiet mode
  -t int
        Number of Threads (default 1)
  -v    Debug/verbose mode

Example: ./mul '4a4b' 'a&b'
        ./mul -perm 100000 -mul 5 '4a4b4c4d' '(a & b & (c ^ d) | ((a ^ b) & c & d)'
```
