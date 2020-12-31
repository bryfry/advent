class Bunch:
    def __init__(self, **kwds):
        self.__dict__.update(kwds)


def read_input():
    with open("input/d02p1.txt") as f:
        return f.readlines()


def parse(line):
    c = line.split(":")
    pw = c[1].lstrip()
    req = c[0].lstrip()

    s = req.split(" ")
    letter = s[1]
    counts = s[0]

    d = counts.split("-")
    pos1 = int(d[0])
    pos2 = int(d[1])

    return Bunch(
        pw=pw,
        letter=letter,
        pos1=pos1,
        pos2=pos2,
    )


def valid_p1(p):
    letter_count = p.pw.count(p.letter)
    if letter_count >= p.pos1 and letter_count <= p.pos2:
        return 1
    return 0


def valid_p2(p):
    pos1_valid = p.pw[p.pos1 - 1] == p.letter
    pos2_valid = p.pw[p.pos2 - 1] == p.letter
    if pos1_valid ^ pos2_valid:
        return 1
    return 0


input_lines = read_input()
parsed_lines = [parse(line.strip()) for line in input_lines]
valid_lines_p1 = [valid_p1(parsed_line) for parsed_line in parsed_lines]
valid_lines_p2 = [valid_p2(parsed_line) for parsed_line in parsed_lines]
print("p1:", sum(valid_lines_p1))
print("p2:", sum(valid_lines_p2))
