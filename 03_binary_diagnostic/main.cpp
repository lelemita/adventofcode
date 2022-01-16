// puzzle: https://adventofcode.com/2021/day/3
#include <iostream>
#include <fstream>
#include <string>
#include <cmath>
#include <algorithm>
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

int part02(const char* path, int len) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    string lines[len];
    string line;
    int i = 0;
    while(getline(ifs, line)) {
        lines[i++] = line;
    }

    int begin = 0;
    int end = len;
    int jdx = 0;
    int count = 0;
    while (count++ < 7) {
        int middle = (begin + end) / 2;
        sort(&lines[begin], &lines[end]);  
        // cout << "------- " << jdx << " : " << begin << "-" << end << endl;
        // for (int k = begin; k < end; k++) 
        // {
        //     cout << k << " : " << lines[k] << endl;
        // }
        int save = begin;
        // while (lines[begin++][jdx] == '0') {}
        while (true) {
            char ch = lines[begin][jdx];
            if (ch == '1') {
                break;
            }
            begin += 1;
        }

        if (begin <= middle) {
            // cout << " 대세는 1" << endl;
            // begin -= 1;
        } else {
            // cout << " Win Zero " << begin << " : " << middle << endl;
            end = begin;
            begin = save;
        }
        if (end - begin == 0) {
            break;
        }
        jdx = (jdx == line.length() -1)? 0 : ++jdx;
    }
    return binToDecimal(lines[begin]);
}

int main(void)
{
    const char* course1 = "../example";
    cout << part01(course1) << endl;
    cout << part02(course1, 12) << endl;

    // const char* course2 = "../input";
    // cout << part01(course2) << endl;
    // cout << part02(course2) << endl;
    return 0;
}
