bool isMatch(char* s, char* p) {
    char *star = NULL;
    char *tmp = s;
    while(*s != '\0'){
        if(*p == *s || *p == '?'){
            s++;
            p++;
        }else if(*p == '*'){
            star = p++;
            tmp = s;
        }else if(star){
            s = ++tmp;
            p = star + 1;
        }
        else 
            return 0;
    }
    while(*p == '*') ++p;
    return !*p;
}
