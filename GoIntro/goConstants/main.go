package main

// conts are assigned by compiler time, not in run time
import "fmt"

const (
	// Create a huge number by shifting a 1 bit left 100 places.
	// In other words, the binary number that is 1 followed by 100 zeroes.
	Big = 1 << 100
	// Shift it right again 99 places, so we end up with 1<<1, or 2.
	Small = Big >> 99
)

func needInt(x int) int { return x*10 + 1 }
func needFloat(x float64) float64 {
	return x
}

func main() {
	//example between compile time and run time
	const num1 = 1 // it works so well, because number 1 is already defined
	//const num2 = math.Sin(1.57) // it didn't works because need the run time for execute the function and after assigne it for num2
	const (
		A = iota // it increase automatic, it starts in zero. iota normaly is used to initialize enumerated consts
		B        // untyped consts are good for make some things with number of different types
		C
	)
	var b int16 = 24
	fmt.Printf("%v, %T", A+b, A+b) // you can do it because the compilor auto convert the type of the variable A to int16 becuse it's an untyped const
	// out of print -> 24, int16
	// fmt.Println(A, B, C) //-> 0 1 2
	// fmt.Println(Small) //-> 2
	// fmt.Println(needInt(Small)) //-> 21
	// fmt.Println(needFloat(Small)) //-> float
	// fmt.Println(needFloat(Big)) //-> 1.2676506002282294e+30
}
