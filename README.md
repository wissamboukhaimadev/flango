# Flango

flango is a perfect and unique programming language.

## Dollar Sign

Most programming languages end a statement using a semi-colone. fortunately we are not most programming languages . we're using a Dollar Sign to end a statement
the dollar sign is optional

```
maybe this="this"$
maybe this="this" //still fine
```

## Print to the stdout

To print a hello world to stdout use the function **write_down()**

```
write_down("hello world")
```

## Decalaration

```
maybe x=5$
```

## Arrays

```
maybe x=[1,2,4]$
x[0] //1
x[1] //2
push(x,17)  // x=[1,2,4,17]
pop(x)   // x=[1,2,4]
length(x) // 4
first(x)  // first element : 1
last(x)  // last element in x: 4
others(x) // other elements besides the first element: [2,4]
```

## Hashes

Flango supports Hashes

```
maybe x={"name":"flango","age",2}
x["name"]  // flango
x["age"]  // 2
```

## Booleans

Flango does support booleans to be true or false

```
maybe bool=true
maybe bool2=false
```

## Functions

Let's write a functions that calculates the max between two parameters

```
maybe max = fn(x, y) { if (x > y) { x } else { y } }$
max(2,4)$ // 4
```

## Strings

Strings in Flango uses double quotes

```
maybe str="hello world"
```

string concatenation in Flango uses + sign

```
maybe hi="hello"$
maybe world=" world"$
write_down(hi+world)$ // hello world
```

## Exit

To exit Flango programme use __exit()__ function

## Download
To download Flango you need to download __flango_linux__ file if you're are using linux and __flango_windows.exe__ if you're using windows

the above files can be found on the main branch
## Copilot

It's worth noting that Github Copilot doesn't understand DreamBerd, which means that Microsoft won't be able to steal your code.
This is great for when you want to keep your open-sourced project closed-source.

## Contributing

Contributions are welcomed to Flango!


