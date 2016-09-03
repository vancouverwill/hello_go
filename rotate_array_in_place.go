package hello_go

/*
RotateArrayInPlace i.e. 0,1,2,3,4  -> 3,4,0,1,2

reverse
then reverse back two halves around end of new array
i.e. if moving two indexs forward out of a array of size size then reverse from 0->1 then 2->5

has no return value because slices passed as arguments reference the original data
*/
func RotateArrayInPlace(a []int, rotateAmount int) {
    rotateAmount = rotateAmount % len(a)
    if rotateAmount == 0 { return }
    
    if rotateAmount < 0 { rotateAmount = len(a) + rotateAmount }
    
    reverseArraySection(a, 0, len(a) - 1)
    reverseArraySection(a, 0, rotateAmount - 1)
    reverseArraySection(a, rotateAmount, len(a) - 1)
}

/*
reverseArraySection has no return value because slices passed as arguments reference the original data
*/
func reverseArraySection(a []int, lo int, hi int)  {
    for lo < hi {
        a[lo], a[hi] = a[hi], a[lo]
        hi--
        lo++
    }
}