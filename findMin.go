package findMin
//Package findMin calculates the minimum in array



var minimum int = 1000 //minimum is a priavte variable


func FindMinimum(array []int) int { //FindMinimum is a global function
for k := 0; k <= 5 ; k++{
if minimum > array[k]{
minimum=array[k];
}
}
return minimum
}