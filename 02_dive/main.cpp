// puzzle: https://adventofcode.com/2021/day/2
// input: https://adventofcode.com/2021/day/2/input
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

const string DELM = " ";

int part01(const char* path) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    int hori = 0;
    int depth = 0;
    string line;
    while (getline(ifs, line)) {
        size_t pos = line.find(DELM);
        int num = stoi(line.substr(pos));
        switch (line[0])
        {
        case 'f':
            hori += num;
            break;
        case 'u':
            depth -= num;
            break;
        case 'd':
            depth += num;
            break;
        }
    }
    ifs.close();
    return hori * depth;
}

int part02(const char* path) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    int aim = 0;
    int hori = 0;
    int depth = 0;
    string line;
    while (getline(ifs, line)) {
        size_t pos = line.find(DELM);
        int num = stoi(line.substr(pos));
        switch (line[0])
        {
        case 'f':
            hori += num;
            depth += aim * num;
            break;
        case 'u':
            aim -= num;
            break;
        case 'd':
            aim += num;
            break;
        }
    }
    ifs.close();
    return hori * depth;
}

int main(void)
{
    const char* course1 = "../example";
    cout << part01(course1) << endl;
    cout << part02(course1) << endl;

    const char* course2 = "../input";
    cout << part01(course2) << endl;
    cout << part02(course2) << endl;
    return 0;
}
