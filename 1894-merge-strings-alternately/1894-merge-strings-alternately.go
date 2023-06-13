func mergeAlternately(word1 string, word2 string) string {
   
    arr1 := []rune(word1)
    arr2 := []rune(word2)
   
    var merged []rune
    leng := len(word1)
    
    if len(word2) > leng {
        leng = len(word2)
    }

    merged = make([]rune, 0, leng*2)

    for i := 0; i <= leng; i++{
        if i < len(word1) {
			merged = append(merged, arr1[i])
		}
		if i < len(word2) {
			merged = append(merged, arr2[i])
		}
    }

    result := string(merged)
    return result
}