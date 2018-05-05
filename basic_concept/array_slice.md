https://blog.golang.org/slices

# Array
* fixed-size
* size of array is part of its type
* `var buffer [256]byte`
  * variable buffer, which holds 256 bytes. The type of buffer includes its size
  * An array of with 512 bytes would be of the distinct type [512]byte
  * buffer looks like this in memory: `buffer: byte byte byte ... byte`
  * `len(buffer)`: 256

# Slices
* A slice is not an array, it describes a piece of an array
* `var slice []byte = buffer[100:150]`
  * `[]byte`: slice of bytes
  * aka `var slice = buffer[100:150]` aka `slice := buffer[100:150]`
* Think of a slice like this:

```
type sliceHeader struct {
    Length        int
    ZerothElement *byte
}

slice := sliceHeader{
    Length:        50,
    ZerothElement: &buffer[100],
}
```

* Slice of a slice: `slice2 := slice[5:10]`

```
slice2 := sliceHeader{
    Length:        5,
    ZerothElement: &buffer[105],
}
```

* Slice header e.g: https://golang.org/pkg/bytes/#IndexRune

## Passing slices to function
* Slice is a struct "value" holding a pointer and a length, it is NOT a pointer to a struct
* i.e: when we called `IndexRune`, it was passed a copy of the slice header

```
func SubtractOneFromLength(slice []byte) []byte {
    slice = slice[0 : len(slice)-1]
    return slice
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    newSlice := SubtractOneFromLength(slice)
    fmt.Println("After:  len(slice) =", len(slice))
    fmt.Println("After:  len(newSlice) =", len(newSlice))
}
```

## Pointers to slices: Method receivers
* Pass a pointer to a slice header

```
func PtrSubtractOneFromLength(slicePtr *[]byte) {
    slice := *slicePtr
    *slicePtr = slice[0 : len(slice)-1]
}

func main() {
    fmt.Println("Before: len(slice) =", len(slice))
    PtrSubtractOneFromLength(&slice)
    fmt.Println("After:  len(slice) =", len(slice))
}
```

* Pointer receiver
```
type path []byte

func (p *path) TruncateAtFinalSlash() {
    i := bytes.LastIndex(*p, []byte("/"))
    if i >= 0 {
        *p = (*p)[0:i]
    }
}

func main() {
    pathName := path("/usr/bin/tso") // Conversion from string to path.
    pathName.TruncateAtFinalSlash()
    fmt.Printf("%s\n", pathName)
}
```

## Capacity
* slice header has "capacity"

```
type sliceHeader struct {
    Length        int
    Capacity      int
    ZerothElement *byte
}
```

* Capacity field is equal to the length of the underlying array
* `cap(slice)`

## Make
* to grow the slice beyond its capacity
  * allocating a new array
  * copying the data over
  * modifying the slice to describe the new array

* `slice := make([]int, 10, 15)`: type of slice, initial length, capacity
* `gophers := make([]Gopher, 10)`: shorthand

## Copy
* `copy(newSlice, slice)`
* Only copies what it can

# Append: An example
* Implement a function that take a slice, and append a value behind the slice.
* 
```
func Extend(slice []int, element int) []int {
    n := len(slice)
    if n == cap(slice) {
        newSlice := make([]int, len(slice), 2*len(slice)+1)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[0 : n+1]
    slice[n] = element
    return slice
}
```

* Implement `append` to append the items to the slice
* 
```
func Append(slice []int, items ...int) []int {
    for _, item := range items {
        slice = Extend(slice, item)
    }
    return slice
}

    slice1 := []int{0, 1, 2, 3, 4}
    slice2 := []int{55, 66, 77}
    fmt.Println(slice1)
    slice1 = Append(slice1, slice2...) // The '...' is essential!
    fmt.Println(slice1)
```

* Use `copy` for efficiency

```
// Append appends the elements to the slice.
// Efficient version.
func Append(slice []int, elements ...int) []int {
    n := len(slice)
    total := len(slice) + len(elements)
    if total > cap(slice) {
        // Reallocate. Grow to 1.5 times the new size, so we can still grow.
        newSize := total*3/2 + 1
        newSlice := make([]int, total, newSize)
        copy(newSlice, slice)
        slice = newSlice
    }
    slice = slice[:total]
    copy(slice[n:], elements)
    return slice
}
```

## Build-in append:
* call `append` will return a new slice

## Nil slice
* `[]int(nil)` aka `sliceHeader{}`: the element pointer is nil too

```
sliceHeader{
    Length:        0,
    Capacity:      0,
    ZerothElement: nil,
}
```

* A nil slice has length zero and can be appended to, with allocation

## Strings
* read-only slices of bytes
* no need for a capacity, can't grow them


