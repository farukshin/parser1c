#!/bin/bash
function myfunc {
    for var in $(find ./example/ -name '*.log' -type f)
    do
    ./parser_tj_1c  --input=$var --format=json > /dev/null
    echo "done $var"
    done
}

time myfunc