// puzzle: https://adventofcode.com/2021/day/25
// input: https://adventofcode.com/2021/day/25/input
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

const string DELM = " ";

int part01(const char* path, int fileLen) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }
    
    string line;
    getline(ifs, line);
    int lineSize = line.size();
    cout << lineSize << endl;
    char current[fileLen * lineSize];
    char next[fileLen * lineSize];
    for (int i = 0; i < lineSize-2; i++)
    {
        current[i] = line[i];
    }
    
    cout << current << endl;

    // string current[fileLen];
    // string line;
    // int idx;
    // while (getline(ifs, line)) {
    //     current[idx++] = line;
    // }

    


    return 0;
}


int main(void)
{
    const char* course1 = "../example";
    cout << part01(course1, 9) << endl;
    // cout << part02(course1) << endl;

    // const char* course2 = "../input";
    // cout << part01(course2) << endl;
    // cout << part02(course2) << endl;
    return 0;
}
