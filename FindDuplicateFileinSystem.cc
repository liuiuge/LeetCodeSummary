

#include <vector>
#include <iostream>
#include <string>
#include <unordered_map>
#include <sstream>

using std::vector;
using std::stringstream;
using std::string;
using std::unordered_map;

class Solution {
public:
    vector<vector<string>> findDuplicate(vector<string>& paths) {
        vector<vector<string>> result;
        unordered_map<string, vector<string>> files;        
        for(auto path:paths)
        {
            stringstream ss(path);
            string root;
            string s;
            getline(ss,root,' ');
            while(getline(ss,s,' '))
            {
                string filename = root + '/' + s.substr(0,s.find('('));
                string filecont = s.substr(s.find('(')+1,s.find('(')-s.find(')')-1);
                files[filecont].push_back(filename);
            }
        }
        for(auto name:files)
        {
            if(name.second.size()>1)
                result.push_back(name.second);
        }
        return result;
    }
};
