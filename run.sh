# go run main.go

# FILE="./out/test_program.out"

# if [ -n "$1" ]; then
#     FILE = $1
# fi

cat $1 | tcc -run ./resources/tinyc/tinyc.c

if [ $? -eq 0 ]; then
    echo "Success"
fi