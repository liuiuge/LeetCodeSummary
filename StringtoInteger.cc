int myAtoi(char* str) 
{
    char *p = str;
    int natv = 0;
    unsigned long ans = 0;
    int out;
    int check_signed = 0;
    
    /* ignore the space of begianing */
    while(*p == ' ')
        p++;
    
    /* check signed */
    if(*p == '+'){
        p++;
    }else if(*p == '-'){
        natv = 1;
        p++;
    }
    
    while(*p != '\0'){
        if(*p > '9' || *p < '0'){
            break;
        }
        
        ans = ans*10 + *p-'0';
        p++;
        
        /* over flow */
        if(ans > 0x7fffffff){
            return natv ? 0x80000000 : 0x7fffffff;
        }
    }
   
    out = (int)ans;
    
    return natv ? (~out+1) : out;
}
