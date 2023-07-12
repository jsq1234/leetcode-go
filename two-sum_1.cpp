/*

[2,7,11,15]
9;

[3,2,4]
6;

[3,3]
6;

*/
class Solution {
public:
    vector<int> twoSum(vector<int>& nums, int target) {
     unordered_map<int,int> map;
        int N = nums.size();
        vector<int> v(2);
        for(int i=0; i<N; i++)
        {
            int key = target - nums[i];
            if(map.find(key) != map.end()) 
            {
                v[0] = map[key];
                v[1] = i; 
                break;
            }
            map[nums[i]] = i;
        }
        
        return v;    
    }
};
