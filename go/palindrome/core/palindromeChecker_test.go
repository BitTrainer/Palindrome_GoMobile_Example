package palindromeCheck
import(
    "testing"
)

func TestAnEmptyStringIsAPalindrome(test *testing.T){
    input := ""
    actual := IsPalindrome(input)
    if (!actual) {
        test.Log("Empty string does not evaluate as being a palindrome.")
        test.Fail();
    }
}

func TestOneCharacterStringIsAPalindrome(test *testing.T){
    input := "a"
    actual := IsPalindrome(input)
    if (!actual) {
        test.Log("One character string does not evaluate as being a palindrome.")
        test.Fail();
    }
}

func TestPalindromeWithMarkersIsAPalindrome(test *testing.T){
    input := "maddam i'm addam."
    actual := IsPalindrome(input)
    if (!actual) {
        test.Log("Palindrome with markers does not evaluate as being a palindrome.")
        test.Fail();
    }
}

func TestPalindromeWithCaseIsAPalindrome(test *testing.T){
    input := "Maddam I'm Addam."
    actual := IsPalindrome(input)
    if (!actual) {
        test.Log("Palindrome with case does not evaluate as being a palindrome.")
        test.Fail();
    }
}