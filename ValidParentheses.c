bool isValid(char* s) {
    char stack[10000];
    int top = -1;
    while(*s != '\0'){
       if (*s == '['  || *s == '{' || *s == '(') {
            stack[++top] =  *s;
        }else if (top < 0){
            return 0;
        }else if (*s == ']') {
            if (stack[top--] != '[') {
                return 0;
            }
        }else if (*s == '}') {
            if (stack[top--] != '{') {
                return 0;
            }
        }else if (*s == ')') {
            if (stack[top--] != '(') {
                return 0;
            }
        } else return 0;
        ++s;
    }
    if(top!=-1)
        return 0;
    return 1;
}
