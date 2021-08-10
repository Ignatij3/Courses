package main

type trio struct {
    n int
    pchar *rune
    next *trio
}    

func main() {
    var s *trio
    s = &trio{7, nil, nil}
    (*s).pchar = new(rune)
    *((*s).pchar) = 'R'
    ((*s).next) = &trio{9, new(rune), s}
    (*(*((*s).next)).pchar) = 'W'
}
