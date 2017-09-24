int strStr(char* haystack, char* needle) {
    if(strstr(haystack,needle)){
        return (int)(strstr(haystack,needle)-haystack);
    }else{
        return -1;
    }
}
