
def print_map(current):
    for line in current:
        for c in line:
            print(c, end='')
        print()


def count_ch(target, array):
    count = 0
    for line in array:
        for c in line:
            if c == target:
                count += 1
    return count


def move(way, current):
    is_stop = True
    next = [[x for x in line] for line in current]
    max_i = len(current) -1
    max_j = len(current[0]) -1
    for i, line in enumerate(current):
        for j, c in enumerate(line):           
            if c == way:
                after_i = i + 1 if i < max_i else 0
                after_j = j + 1 if j < max_j else 0
                after = current[i][after_j] if way == '>' else current[after_i][j]
                if after == '.':
                    is_stop = False
                    next[i][j] = '.'
                    if way == '>':
                        next[i][after_j] = way
                    else:
                        next[after_i][j] = way  
                else:
                    next[i][j] = way
    return is_stop, next

def part_1(file_path):
    step = 0
    current = []
    f = open(file_path, 'r')
    for line in f.readlines():
        current.append([x for x in line.strip()])
    f.close()
    
    while True:
        step += 1  
        is_stop_x, current = move('>', current)
        is_stop_y, current = move('v', current)
        if is_stop_x and is_stop_y:
            break
    # print_map(current)
    return step


if __name__ == "__main__":
    print(part_1("./example"))
    print(part_1("./input"))
    
    # print(part_2("./example"))
    # print(part_2("./input"))
