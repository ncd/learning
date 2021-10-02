//https://www.hackerrank.com/challenges/climbing-the-leaderboard/problem
vector<int> climbingLeaderboard(vector<int> ranked, vector<int> player) {
    vector<int> grade;
    vector<int> ans;
    if (ranked.size() == 0) {
        return ans;
    }
    grade.emplace_back(ranked[0]);
    for(int i = 0; i < ranked.size()-1; i++) {
        if(ranked[i] != ranked[i+1]) {
            grade.emplace_back(ranked[i+1]);
        }
    }
    int count = grade.size() - 1;
    for(int j = 0; j < player.size(); j++) {
        while(count >=0 && player[j] > grade[count]) {
            count--;
        }
        if(count < 0) {
            ans.emplace_back(1);
        }
        else if(player[j] == grade[count]) {
            ans.emplace_back(count+1);
        }
        else {
            ans.emplace_back(count+2);
        }
    }
    return ans;
}