package main

import "fmt"

//SECTION 1: 🌻Stack & Pointers🌻
//===============================================================
//NOTES :
/*
	⭐️ in golang every goroutine would have its own stack

	⭐️lets say a goroutine wants to call a function...
		what would happen is that an isolated memory space...
		would be created on the stack and assigned....
		to that funcion call which would be called ...
		that functions's "frame boundary"

	⭐️a function call provided with some data arguments...
		would require a transition of data...
		from the caller function's stack frame...
		to the newly called function's stack frame.

	⭐️in golang this transition of data between
		two stack frames is always done "by value"...
		meaning The value will be copied,...
		passed and placed into the new frame.

	⭐️if we do wish to "share a value" however between 2 function calls...
		and their stack frames. We would use something called...
		a "pointer" to that value.

	⭐️"Pointers" allow us to share a value with a function so...
		the function can read and write to that value...
		even though the value does not exist directly inside its own frame.

	⭐️A pointer variable would store the address of...
		the value it is pointing to allowing the executing...
		functionn frame to have indirect memory access to...
		that value outside of it's own boundry frame.

	⭐️For every type that is declared....
		either by you or the language itself,...
		a compliment pointer type would also be declared.
*/
//===============================================================
func main() {
	fmt.Println("<<Stack & Pointers>>\n=======================")
	stackFrameLocation := "|*main* frame stack|"

	/*lets initianlize some variables
	and see their addreses in the main function
	stack frame*/
	a := 22 //int

	//🌼PASS BY VALUE🌼
	fmt.Println("\t\t--PASS BY VALUE--")

	fmt.Println("!BEFORE CALLING *zeroDataByValue*!")
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of a:[", &a, "], value of a:[", a, "]")
	fmt.Println("------------")

	/* lets call a function
	pass our variables "by value" to it
	see it create a new stack frame and store
	copy of our value in that frame */
	fmt.Println("\n!CALLING *zeroDataByValue*!")
	zeroDataByValue(a)

	/* as we see the variable a is not affected by changes made
	to its value in the function call amd the function fooReciveByValue
	has no access to this location in memory considering its outside its frame
	*/
	fmt.Println("\n!AFTER CALLING *zeroDataByValue*!")
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of a:[", &a, "], value of a:[", a, "]")
	fmt.Println("------------")

	//🌼SHARE WITH POINTERS🌼
	fmt.Println("\n\n\t\t--SHARE WITH POINTERS--")
	/*
		lets first create a pointer variable
		pointing to our variable a
		we would expect this pointer2a to have :
		1- been stored in the stack frame of main
		2- a value equal to the address of the variable its pointing to
	*/
	var pointer2a *int = &a

	fmt.Println("!BEFORE CALLING *zeroDataBySharing*!")
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of a:[", &a, "], value of a:[", a, "],")
	fmt.Println("addr of pointer2a:[", &pointer2a, "], value of pointer2a:[", pointer2a,
		"], value of where pointer2a is pointing:[", *pointer2a, "]")
	fmt.Println("------------")

	/*
		we now wish to actualy share the same memory
		with our function call instead of just creating a
		a copy of its value so we use the help of
		pointer variables
	*/
	fmt.Println("\n!CALLING *zeroDataBySharing*!")
	zeroDataBySharing(pointer2a)

	/*
		as we see throughout this method we indeed
		allowed a function indirect access
		to a varible stored outside its own frame and therefor
		it was able to manipulate our data
	*/
	fmt.Println("\n!AFTER CALLING *zeroDataBySharing*!")
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of a:[", &a, "], value of a:[", a, "],")
	fmt.Println("addr of pointer2a:[", &pointer2a, "], value of pointer2a:[", pointer2a,
		"], value of where pointer2a is pointing:[", *pointer2a, "]")
	fmt.Println("------------")
}

//go:noinline
func zeroDataByValue(a int) {
	stackFrameLocation := "|*zeroDataByValue* frame stack|"

	/* the variable "a"	in this function would have the same value as the
	variable "a" in the main function but it would have a
	different address meaning it is actually stored sepreatly
	*/
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of a:[", &a, "], value of a:[", a, "]")
	fmt.Println("------------")

	/*changing this variable
	would not affect the original "a"
	considering they are stored on different parts of the memory
	unrelated to eachother
	*/
	a = 0
}

//go:noinline
func zeroDataBySharing(pointer2a *int) {
	stackFrameLocation := "|*zeroDataBySharing* frame stack|"

	/* as we discused all transitions of data between 2 stack
	frames are done "by value" even when using pointers! as we
	can see the addresse of the two "pointer2a" variables is indeed
	different and they are independant variables stored separately*/
	fmt.Println(stackFrameLocation)
	fmt.Println("addr of pointer2a:[", &pointer2a, "], value of pointer2a:[", pointer2a,
		"], value of where pointer2a is pointing:[", *pointer2a, "]")
	fmt.Println("------------")

	/*considering the "by value" nature of the function call data transfer
	both the "pointer2a"s would have the same value
	meaning they point to the same exact memory spot and therefor
	these 2 function can both manipulate and access the data inside that
	corresponding address and they "share" it together and if one changes the content of this
	shared pointed data it would infact be seen in the other one as well
	*/
	*pointer2a = 0
}

//===============================================================

//SECTION 2: 🌺Heap & EscapeAnalysis🌺
// ========================================
//NOTES :
/*
	🌸the data of our code in the go language get stored...
	either in the stack or the heap

	🌸the go compiler is designed so that it prefers storage of...
	data on the stack over the storage on the heap considering...
	the higher perfomance the stack memory has

	🌸but in some certain scenarios it does however...
	decide to alocate its memory on the heap and...
	these scenarios would be :

	🌸-1: when the data would be to large to fit in the stack

	🌸-2: when the size of the data is not known in advanced

	🌸-3: in the case of interface type variables value assignments

	🌸-4: when the value created is being "shared up" the stack

	🌸"sharing up" a value would be the scenario where a called function...
	tries to returns a pointer to a value which would be stored...
	on its own stack frame. considering after the return statment the stack...
	frame of a function call would be considered invalid then the storage of...
	that value on the stack would create integrity issues in our program...
	and there for the value would "escape"/be stored in the heap.

	🌸another thing to watch out for is the concept of "inline"ing...
	a function call, which is an optimization techniche used by...
	the compiler.

	🌸when the compiler sees a small function being called...
	it might deside to just place the instructions of the called function...
	directly inside the caller function, effectivly saving the overload...
	of performing a function call.

	🌸this "inline"ing strategy could also effect the place ...
	a value is assigned to and the memory allocations.

	🌸the algorythm used by the compiler to determine whether a...
	value is indeed being shared up or not is called "escape analysis"...
	and if the compiler has any doubt that the value could...
	possibly be refrence in any direct or indirect way higher up...
	the stack it won't hesitate storing it in the heap.
*/

// ===============================================================
//becnh marks and test code yet to be added ... :3
