#!/bin/zsh

csmith --arrays --no-bitfields --no-checksum --comma-operators --embedded-assigns --compound-assignment --consts --divs --no-pre-incr-operator --no-pre-decr-operator --no-post-incr-operator --no-post-decr-operator --no-jumps --no-longlong --int8 --uint8 --float --main --no-math64 --no-inline-function --muls --no-safe-math --no-packed-struct --no-paranoid --no-pointers --no-structs --no-unions  --no-volatiles --no-volatile-pointers --no-const-pointers --no-builtins --no-return-structs --no-arg-structs --no-return-unions --no-arg-unions --take-no-union-field-addr --no-vol-struct-union-fields --no-const-struct-union-fields --no-dangling-global-pointers --no-return-dead-pointer --no-union-read-type-sensitive --no-hash-value-printf --no-signed-char-index > "test.c"

#for i in {1..10}
#do
#    csmith $flags > "$i.c"
#done