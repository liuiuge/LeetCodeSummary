int maxArea(int* height, int heightSize) {
    int head,tail, left,right, maxi, vol;
    head = 0, tail = heightSize-1;
    maxi = 0;
    while(head!=tail){
        left = height[head];
        right = height[tail];
        if( left < right ){
            vol = left*(tail-head);
            if( vol>maxi )
                maxi = vol;
            head++;
        }
        else{
            vol = right*(tail-head);
            if( vol>maxi )
                maxi = vol;
            tail--;
        }
    }
    return maxi;
}
