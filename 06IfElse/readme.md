# **If Else Statements**

Go's if statements are like its for loops; the expression need not be surrounded by parentheses ( ) but the braces { } are required.

```sh
func max(x int, y int) int {
    if x < y{
        return y
    }
	return x
}
func main() {
	a := max(2, 1)
	fmt.Println(a) //happens to be 2 for this example
}
```
if statement can start with a short statement to execute before the condition.

Variables declared by the statement are only in *scope until the end of the if*.

**(Try using i in the last return statement.)*

```sh
func youngMestemberInFamilyToVote(age1 float64, age2 float64, ageLimit float64) string {
	if i := math.Min(float64(age1), float64(age2)); i > ageLimit {
		return "Youngest family member can Vote"
	}
	return "Sorry youngest member of your family can't vote"
}
func main() {
	fmt.Println(youngMestemberInFamilyToVote(16, 25, 18))
}
```

**Variables declared inside an if short statement are also available inside any of the else blocks.*

## There is no ternary if in Go, so you’ll need to use a full if statement even for basic conditions.
:frowning:

Try some exercise,
[Go flowcontrol exercise](https://go.dev/tour/flowcontrol/8)