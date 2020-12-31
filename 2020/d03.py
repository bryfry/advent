import numpy


def read_input():
    with open("input/d03p1.txt") as f:
        return f.readlines()


def parse(line):
    line = line.strip()
    tree_list = []
    for c in line:
        if c == ".":
            tree_list.append(0)
        if c == "#":
            tree_list.append(1)

    return tree_list


def aboreal_count(x_rate, y_rate, tree_field):
    count = 0
    for i in range(0, len(tree_field)):
        y = i * y_rate
        x = (i * x_rate) % len(tree_field[0])
        if y > len(tree_field):
            break
        count += tree_field[y][x]
    return count


input_lines = read_input()
parsed_lines = [parse(line) for line in input_lines]
print("p1:", aboreal_count(3, 1, parsed_lines))
p2_rates = [(1, 1), (3, 1), (5, 1), (7, 1), (1, 2)]
p2_counts = [aboreal_count(d[0], d[1], parsed_lines) for d in p2_rates]
print("p2:", numpy.prod(p2_counts))
