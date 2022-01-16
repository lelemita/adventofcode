// puzzle: https://adventofcode.com/2021/day/3
#include <iostream>
#include <fstream>
#include <string>
#include <list>
#include <cmath>
using namespace std;

int binToDecimal(string line)
{
    int len = line.length();
    int num = 0;
    for (int i = 0; i < len; i++)
    {
        if ((int)line[i] - '0') {
            num += pow(2, len-1-i);
        }
    }
    return num;
}

int part01(const char* path) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    string line;
    getline(ifs, line);
    int lineSize = 1;
    int len = line.length();
    int sums[len] = {};
    do {
        lineSize += 1;
        for (int i = 0; i < len; i++)
        {
            sums[i] += ((int)line[i] - '0');
        } 
    } while (getline(ifs, line));
    ifs.close();

    int gamma = 0;
    int epsilon = 0;
    for (int i = 0; i < len; i++)
    {
        if (sums[i] > lineSize/2) {
            gamma += pow(2, len-1-i);
        } else {
            epsilon += pow(2, len-1-i);
        }
    }
    return gamma * epsilon;
}

int part02(const char* path) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    string line;
    list<string> lines;
    while (getline(ifs, line)) {
        lines.push_back(line);
    }
    ifs.close();

    string oxygen;
    string carbonDioxide;

    int idx = 0;
    while (lines.size() > 1) {
        int sum = 0;
        list<int> onesIdx;
        list<int> sIdx;
        for (auto l : lines)
        {
            sum += (int)line[idx] - '0';
        }
        int target = 0;
        if (sum >= lines.size()/2) {
            target = 1;
        }
        
        
        idx = idx >= line.length() ? 0 : idx+1;
    }
    

    return binToDecimal(oxygen) * binToDecimal(carbonDioxide);
}

int main(void)
{
    const char* course1 = "../example";
    cout << part01(course1) << endl;
    cout << part02(course1) << endl;

    // const char* course2 = "../input";
    // cout << part01(course2) << endl;
    // cout << part02(course2) << endl;
    return 0;
}
