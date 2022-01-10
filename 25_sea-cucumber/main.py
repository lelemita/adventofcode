
def part_1(file_path):
    step = 0
    current = []
    f = open(file_path, 'r')
    for line in f.readlines():
        current.append([x for x in line.strip()])
    f.close()

    next = [[x for x in line] for line in current]
    max_i = len(current) -1
    max_j = len(current[0]) -1
    for i, line in enumerate(current):
        for j, c in enumerate(line):
            
            # 틀린 상태: 동쪽으로 가는 해삼들이 먼저 움직인다고 함....ㅠㅠㅠㅠㅠ
            
            if c == 'v':
                after_i = i + 1 if i < max_i else 0
                after = current[after_i][j]
                if after == '.':
                    next[after_i][j] = 'v'
                    next[i][j] = '.'
                else:
                    next[i][j] = 'v'
            elif c == '>':
                after_j = j + 1 if j < max_j else 0
                after = current[i][after_j]
                if after == '.':
                    next[i][after_j] = '>'
                    next[i][j] = '.'
                else:
                    next[i][j] = '>'       
                
    for line in next:
        print(line)     
    print(count_ch('>', current))    
    print(count_ch('>', next))    
    print(count_ch('v', current))    
    print(count_ch('v', next))    
    return step

def count_ch(target, array):
    count = 0
    for line in array:
        for c in line:
            if c == target:
                count += 1
    return count


if __name__ == "__main__":
    print(part_1("/home/yeon/dev/advent_2021/25_sea-cucumber/example"))
    # print(part_1("./input"))
    
    # print(part_2("./example"))
    # print(part_2("./input"))
