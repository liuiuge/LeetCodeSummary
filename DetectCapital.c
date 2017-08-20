#include <stdio.h>
#include <string.h>
bool detectCapitalUse(char* word) {
    // a-z 97-122
    // A-Z 65-90
    
    int i = 0;
    int len = strlen(word);
    int num_caps = 0;
    
    while(word[i]!='\0') {
        int ascii_of_char = (int)word[i];
        if(ascii_of_char>=65 && ascii_of_char<=90) {
            num_caps++;
        }
        i++;
    }
    
    if(num_caps == len) {
        return true;
    } else if (num_caps == 1) {
        if ((int)word[0]>=65 && (int)word[0]<=90) {
            return true;
        } else {
            return false;
        }
    } else if (num_caps > 1) {
        return false;
    } 
    return true;
    
}
