// https://leetcode.com/contest/weekly-contest-205/problems/minimum-deletion-cost-to-avoid-repeating-letters/
class Solution {
public:
    int minCost(string s, vector<int>& cost) {
        int totalCost = 0;
        int maxValue = cost[0];
        for(int i = 0; i < cost.size() - 1; i++) {
            if(s[i] == s[i+1]) {
                totalCost += min(maxValue, cost[i+1]);
                maxValue = max(maxValue, cost[i+1]);
            }
            else {
                maxValue = cost[i+1];
            }
        }
        return totalCost;
    }
};

