# GoDemo
Go App Demo

INSTALL VSC FOR GO

1.- Install VSC from: https://code.visualstudio.com/download

2.- Install Go Extension:

![Extension Image](https://raw.githubusercontent.com/danieluribee/GoDemo/master/images/Screen%20Shot%202020-03-26%20at%2010.06.03.png)

3.- Update Go Tools:

Open the Command Palette

![Command Palette](https://raw.githubusercontent.com/danieluribee/GoDemo/master/images/Screen%20Shot%202020-03-26%20at%2010.15.18.png)

Type goinstall and select update tools

![GoInstall](https://raw.githubusercontent.com/danieluribee/GoDemo/master/images/Screen%20Shot%202020-03-26%20at%2010.16.03.png)

Select all and click ok

![Select All](https://raw.githubusercontent.com/danieluribee/GoDemo/master/images/Screen%20Shot%202020-03-26%20at%2010.16.42.png)

--------------------------------------------------------------------------------------------------------------------------

BASICS OF GO

Go is a statically typed, compiled programming language designed at Google by Robert Griesemer, Rob Pike, and Ken Thompson. Go is syntactically similar to C, but with memory safety, garbage collection, structural typing, and CSP-style concurrency. The language is often referred to as "Golang" because of its domain name, golang.org, but the proper name is Go.

1.- General form of a function<br/>

```
func function_name( [Parameter list] ) [return_types]{
	body
}
```

<br/>
2.- In the functions the type comes after the parameter name<br/>

```
func add(a int, b int) int {
	return a+b
}
```

<br/>
3.- Go manage two types of variables, package and function level (different scope)<br/>

```
var node, angular bool  //package level

func main(){
	var x int  //function level
}
```

<br/>
4.- Arrays<br/>

```
var array1 [size] int

var array2 = []float32{10.2, 2.1, 3.5}

printf(array2[2])
```

<br/>
5.- To define a structure, you must use type and struct statements. The struct statement defines a new data type, with more than one member for your program. type statement binds a name with the type which is struct in our case.<br/>

```
type Vertex struct {
	X int
	Y int
}

func main() {
	fmt.Println(Vertex{1, 2})
}
```

<br/>
6.- Map<br/>

```
CREATE MAP
m := make(map[string]int)

ADD VALUE
m["Answer"] = 42

EDIT VALUE
m["Answer"] = 48

DELETE VALUE
delete(m, "Answer")

LOOK IF A KEY/VALUE IS PRESENT
v, ok := m["Answer"]
fmt.Println("The value:", v, "Present?", ok)
```
<br/>

7.-Slice

Unlike arrays, slices are typed only by the elements they contain (not the number of elements). To create an empty slice with non-zero length, use the builtin make. Here we make a slice of strings of length 3 (initially zero-valued).

EMPTY SLICE
s := make([]string, 3)

ADD VALUES
s[0] = "a"
s[1] = "b"
s[2] = "c"

COPY SLICE
c := make([]string, len(s))
copy(c, s)

SLICE OPERATORS:

Using already Existing Slice: It is also be allowed to create a slice from the given slice. For creating a slice from the given slice first you need to specify the lower and upper bound, which means slice can take elements from the given slice starting from the lower bound to the upper bound. It does not include the elements above from the upper bound. As shown in the below example:

CREATING SLICE
sliceDemo := []int{90, 60, 40, 50, 34, 49, 30} 

CREATING SLICES USING OPERATORS [High:Low]
var my_slice_1 = oRignAl_slice[1:5] 
my_slice_2 := oRignAl_slice[0:] 
my_slice_3 := oRignAl_slice[:6] 
my_slice_4 := oRignAl_slice[:] 
my_slice_5 := my_slice_3[2:4] 

PRINT RESULTS
fmt.Println("Original Slice:", oRignAl_slice) 
fmt.Println("New Slice 1:", my_slice_1) 
fmt.Println("New Slice 2:", my_slice_2) 
fmt.Println("New Slice 3:", my_slice_3) 
fmt.Println("New Slice 4:", my_slice_4) 
fmt.Println("New Slice 5:", my_slice_5) 

OUTPUT
Original Slice: [90 60 40 50 34 49 30]
New Slice 1: [60 40 50 34]
New Slice 2: [90 60 40 50 34 49 30]
New Slice 3: [90 60 40 50 34 49]
New Slice 4: [90 60 40 50 34 49 30]
New Slice 5: [40 50]

<br/>
8.- Syntax of a for loop<br/>

```
for [condition | (init; condition; increment) | Range]
{
	statement(x)
}
```

USE OF ":="<br/>
:= is known as the short declaration operator.
It is used to declare and initialize the variables only inside the functions.
Here, variables has only local scope as they are only declared inside functions.
Declaration and initialzation of the variables must be done at the same time.
There is no need to put type. If you, it will give error.

The Range keyword is used in for loop to iterate over items of an array, slice, channel or map. With array and slices, it returns the index of the item as integer. With maps, it returns the key of the next key-value pair. Range either returns one value or two. If only one value is used on the left of a range expression, it is the 1st value in the following table.

TRADITIONAL EXAMPLE

```
package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)
}
```

GOLANG EXAMPLES

EXAMPLE 1

```
strDict := map[string]string{"Japan": "Tokyo", "China": "Beijing", "Canada": "Ottawa"}
for index, element := range strDict {
	fmt.Println("Index :", index, " Element :", element)
}
```

EXAMPLE 2

```
for key := range strDict {
	fmt.Println(key)
}
```

EXAMPLE 3

```
for _, value := range strDict {
  fmt.Println(value)
}
```
