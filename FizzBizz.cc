

#include <iostream>
#include <vector>
#include <string>
using std::vector;
using std::to_string;

class Solution {
public:
    vector<string> fizzBuzz(int n) {
        vector<string> fb;
        for(int i=1; i<n+1; ++i)
        {
            if(!(i%5) && !(i%3))
                fb.push_back("FizzBuzz");
            else if(!(i%5))
                fb.push_back("Buzz");
            else if(!(i%3))
                fb.push_back("Fizz");
            else
            {
                /*
                ostringstream os;
                os<<i;
                string str1 = os.str();
                fb.push_back(str1);
                */
                fb.push_back(to_string(i));
            }
        }
        return fb;
    }
};
