#!/bin/sh

# Import a module
. /path/to/script

# Define a variable
hi="hello world"

# Print "hello world"
echo "hello world"

# If, then, else
if [ 1 -eq 1 ]; then
    echo ''
elif [ 2 -eq 2 ]; then
    echo ''
else
    echo ''
fi

# Iterate
for a in 1 2 3 4 5; do
    echo $a
done

# Define a function
foo() {
    echo ''
}

# Run a function
foo
