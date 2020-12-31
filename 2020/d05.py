import pprint
import re

def read_input():
    with open("input/d05p1.txt") as f:
        return f.readlines()

def parse(line):
    line = line.strip()
    row = int(line[:7].replace("B","1").replace("F","0"),2)
    seat = int(line[7:].replace("R","1").replace("L","0"),2)
    return (row * 8 ) + seat

input_lines = read_input()
seat_ids = set(parse(line) for line in input_lines)

print("p1:", max(seat_ids))
missing_seats = set(range(min(seat_ids),max(seat_ids))) - seat_ids
print("p2:", list(missing_seats)[0])
