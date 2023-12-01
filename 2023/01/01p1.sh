#!/bin/bash

for line in $(cat "input.txt")
do
    for (( i=0; i<${#line}; i++ )); do
        chr="${line:$i:1}"
        if [[ $chr =~ ^[0-9]+$ ]]; then
            if ! declare -p num_first &>/dev/null; then
                num_first="$chr"
            fi
            num_second="$chr"
        fi
    done
    amount="$num_first$num_second"
    amount=$((amount+0))
    sum=$((sum+amount))
    unset num_first
done
echo "$sum"