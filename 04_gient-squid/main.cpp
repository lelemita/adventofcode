// puzzle: https://adventofcode.com/2021/day/3
#include <iostream>
#include <fstream>
#include <sstream>
#include <string>
#include <vector>
#include <algorithm>
using namespace std;

vector<string> split(string str, char Delimiter) {
    istringstream iss(str);
    string buffer;
    vector<string> result;
    while (getline(iss, buffer, Delimiter)) {
        if(buffer.length() == 0) {
            continue;
        }
        result.push_back(buffer);
    }
    return result;
}

int part01(const char* path, int playerNum) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }
    // 부르는 숫자 읽기
    string line;
    getline(ifs, line);
    vector<string> strNums = split(line, ',');
    int nums[strNums.size()];
    for (int i = 0; i < strNums.size(); i++)
    {
        int num = stoi(strNums[i]);
        nums[i] = num;
    }
    
    // 게임 판 읽기
    int board[playerNum][25];
    int totalSum[playerNum] = {};
    int player = 0;
    while(getline(ifs, line)) {
        vector<string> strPan;
        for (int i = 0; i < 5; i++)
        {
            getline(ifs, line);
            vector<string> temp = split(line, ' ');
            strPan.insert(strPan.end(), temp.begin(), temp.end());
        }
        for (int i = 0; i < strPan.size(); i++)
        {
            int num = stoi(strPan[i]);
            board[player][i] = num;
            totalSum[player] += num;
        }
        player += 1;
    }
    ifs.close();

    // 빙고게임
    int where[player][2][5] = {};
    int callSum[player] = {};
    int lastNum = -1;
    int winner = -1;
    for (int n = 0; n < strNums.size(); n++)
    {
        bool isBingo = false;
        for (int p = 0; p < player; p++){
            // 마킹 & 빙고인지 확인
            int idx = -1;
            for (int w = 0; w < 25; w++) {
                if (board[p][w] == nums[n]) { 
                    idx = w;
                    break;
                }
            }
            if (idx == -1) { continue; }
            callSum[p] += nums[n];
            int i = idx/5;
            int j = idx%5;
            if (where[p][0][i] == 4 || where[p][1][j] == 4) {
                isBingo = true;
                winner = p;
                lastNum = nums[n];
                break;
            }
            where[p][0][i] += 1;
            where[p][1][j] += 1;
        }
        if (isBingo) {
            break;
        }
    }
    
    return (totalSum[winner] - callSum[winner]) * lastNum;
}

int part02(const char* path, int playerNum) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }
    // 부르는 숫자 읽기
    string line;
    getline(ifs, line);
    vector<string> strNums = split(line, ',');
    int nums[strNums.size()];
    for (int i = 0; i < strNums.size(); i++)
    {
        int num = stoi(strNums[i]);
        nums[i] = num;
    }
    
    // 게임 판 읽기
    int board[playerNum][25] = {};
    int totalSum[playerNum] = {};
    int player = 0;
    while(getline(ifs, line)) {
        vector<string> strPan;
        for (int i = 0; i < 5; i++)
        {
            getline(ifs, line);
            vector<string> temp = split(line, ' ');
            strPan.insert(strPan.end(), temp.begin(), temp.end());
        }
        for (int i = 0; i < strPan.size(); i++)
        {
            int num = stoi(strPan[i]);
            board[player][i] = num;
            totalSum[player] += num;
        }
        player += 1;
    }
    ifs.close();

    // 빙고게임
    int where[player][2][5] = {};
    int callSum[player] = {};
    int lastNum = -1;
    int loser = -1;
    bool isWin[playerNum] = {};
    for (int n = 0; n < strNums.size(); n++)
    {
        for (int p = 0; p < player; p++){
            if (isWin[p]) { continue; }
            // 마킹 & 빙고인지 확인
            int idx = -1;
            for (int w = 0; w < 25; w++) {
                if (board[p][w] == nums[n]) { 
                    idx = w;
                    break;
                }
            }
            if (idx == -1) { continue; }
            callSum[p] += nums[n];
            int i = idx/5;
            int j = idx%5;
            if (where[p][0][i] == 4 || where[p][1][j] == 4) {
                isWin[p] = true;
            }
            where[p][0][i] += 1;
            where[p][1][j] += 1;
        }
        bool isEnd = true;
        for (int i = 0; i < playerNum; i++)
        {
            if (isWin[i] == false) {
                isEnd = false;
                loser = i;
                break;
            }
        }
        if (isEnd) {
            lastNum = nums[n];
            break;
        }
    }
    return (totalSum[loser] - callSum[loser]) * lastNum;
}

int main(void)
{
    const char* course1 = "../example";
    cout << part01(course1, 3) << endl;
    cout << part02(course1, 3) << endl;

    const char* course2 = "../input";
    cout << part01(course2, 100) << endl;
    cout << part02(course2, 100) << endl;
    return 0;
}
