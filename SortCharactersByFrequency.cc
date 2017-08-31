class Solution {
public:
    static bool freq_sort(const pair<int,string> & lhs, const pair<int,string> & rhs)
    {
        return lhs.first>rhs.first;
    }
    string frequencySort(string s) {
        vector<pair<int,string> > freq(128);
        for(auto & c:s)
        {
            freq[c].first++;
            freq[c].second = c;
        }
        sort(freq.begin(),freq.end(),Solution::freq_sort);
        string ret;
        for(auto elem:freq)
        {
            for(int i=0; i<elem.first;++i)
                ret.append(elem.second);
        }
        return ret;
    }
};
