package palindromeCheck

import(
    "strings"
    )
    
func reverse(s string) (result string) {
  for _,v := range s {
    result = string(v) + result
  }
  return 
}

//IsPalindrome ...
func IsPalindrome(candidate string) bool {
    if len(candidate) <= 1 {
        return true;
    }
    evaluationString := strings.Replace(strings.ToLower( candidate), " ", "", -1)
    evaluationString = strings.Replace(evaluationString, "'", "", -1)
    evaluationString = strings.Replace(evaluationString, ",", "", -1)
    evaluationString = strings.Replace(evaluationString, ".", "", -1)
    evaluationString = strings.Replace(evaluationString, "!", "", -1)
    evaluationString = strings.Replace(evaluationString, "?", "", -1)
    reverseString := reverse(evaluationString);
    return reverseString == evaluationString;
} 
