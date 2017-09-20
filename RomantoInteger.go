func romanToInt(s string) int {
    var ret,i int = 0,0
    var length = len([]rune(s))
    var number map[byte]int
    number = make(map[byte]int)
    number['I'] = 1
    number['V'] = 5
    number['X'] = 10
    number['L'] = 50
    number['C'] = 100
    number['D'] = 500
    number['M'] = 1000
    for i=length-1;i>=0;i--{ 
        if(i<length-1 && number[s[i]]<number[s[i+1]]){
            ret = ret - number[s[i]]
        }else{
            ret = ret + number[s[i]]
        }
    }
    return ret
}
