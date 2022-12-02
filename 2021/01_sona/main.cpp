// puzzle: https://adventofcode.com/2021/day/1
// input: https://adventofcode.com/2021/day/1/input
#include <iostream>
#include <fstream>
#include <array>
using namespace std;

const int MAX_INT = 2147483647;

int part01(const char* path) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    int count = 0;
    int before = MAX_INT;
    char line[4] = {};
    while (ifs.getline(line, 5)) {
        int num = stoi(line);
        if (num > before) {
            count += 1;
        }
        before = num;
    }
    ifs.close();
    return count;
}

int part02(const char* path, const int len) {
    ifstream ifs(path, ios::in);
    if (ifs.fail()) {
        cout << "input file error" << endl;
        return -1;
    }

    int count = 0;
    char line[4] = {};
    int nums[len];
    int idx = 0;
    while (ifs.getline(line, 5)) {
        nums[idx++] = stoi(line);
    }
    ifs.close();

    int beforeSum = MAX_INT;
    for (int i = 1; i < len-1; i++)
    {
        int sum = nums[i-1] + nums[i] + nums[i+1];
        if (sum > beforeSum) {
            count += 1;
        }
        beforeSum = sum;
    }
    return count;
}

int main(void)
{
    const char* path1 = "../example";
    cout << part01(path1) << endl;
    cout << part02(path1, 10) << endl;

    const char* path2 = "../input";
    cout << part01(path2) << endl;
    cout << part02(path2, 2000) << endl;
    return 0;
}
