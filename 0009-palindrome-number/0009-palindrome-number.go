func isPalindrome(x int) bool {
    copyNumber := x
     res := 0
   for copyNumber>0 {
      remainder := copyNumber% 10
      res = (res * 10) + remainder
      copyNumber /= 10
   }
    if x<0 {
        return false
    }else{
     if res == x{    
        return true
     }
    }
    return false
}