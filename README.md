# Peng Code
A small Go program to convert (a subset of) street slang to English and vice versa.

e.g. arms -> strong

## Usage
Compile the program.

    go build main.go
    
Run the program with command-line arguments:

    ./main "mash man" arms cool bare tourist peng
    
You should get the following output:

    mash man -> gunman
    arms -> strong
    cool -> beast
    bare -> a lot
    tourist -> clueless person
    peng -> good looking

## Testing
To run the matching algorithm tests:

    cd match-finder/
    go test