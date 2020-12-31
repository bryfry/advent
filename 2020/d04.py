import pprint
import re

def read_input():
    with open("input/d04p1.txt") as f:
        content = f.read()
        content_split = re.split(r"(?:\r?\n){2,}", content.strip())
        lines = [line.replace("\n", " ") for line in content_split]
        return lines


def parse(line):
    attrs = line.split(" ")
    passport = {k:v for (k,v) in [tuple(attr.split(":")) for attr in attrs]}
    return passport

def valid_p1(passport):
    if not 'byr' in passport.keys(): return 0 
    if not 'iyr' in passport.keys(): return 0 
    if not 'eyr' in passport.keys(): return 0 
    if not 'hgt' in passport.keys(): return 0 
    if not 'hcl' in passport.keys(): return 0 
    if not 'ecl' in passport.keys(): return 0 
    if not 'pid' in passport.keys(): return 0 
    #if not 'cid' in passport.keys(): return False
    return 1

def valid_p2(passport):

    if not 'byr' in passport.keys(): return 0 
    byr = int(passport['byr'])
    if 2002 < byr or byr < 1920: return 0

    if not 'iyr' in passport.keys(): return 0 
    iyr = int(passport['iyr'])
    if 2020 < iyr or iyr < 2010: return 0

    if not 'eyr' in passport.keys(): return 0 
    eyr = int(passport['eyr'])
    if 2030 < eyr or eyr < 2020: return 0

    if not 'hgt' in passport.keys(): return 0 
    hgt = passport['hgt']
    if not ("cm" in hgt or "in" in hgt): return 0
    if "cm" in hgt:
       hgt_cm = int(hgt.split("c")[0])
       if 193 < hgt_cm or hgt_cm < 150: return 0
    if "in" in hgt:
       hgt_in = int(hgt.split("i")[0])
       if 76 < hgt_in or hgt_in < 59: return 0
      
    if not 'hcl' in passport.keys(): return 0 
    hcl = passport['hcl']
    match = re.match(r'#[0-9,a-z]{6}', hcl)
    if match == None: return 0
    if not (match.start()==0 and match.end()==len(hcl)): return 0

    if not 'ecl' in passport.keys(): return 0 
    ecl = passport['ecl']
    valid_ecls = ['amb','blu','brn','gry','grn','hzl','oth']
    if ecl not in valid_ecls: return 0

    if not 'pid' in passport.keys(): return 0 
    pid = passport['pid']
    match = re.match(r'[0-9]{9}', pid)
    if match == None: return 0
    if not (match.start()==0 and match.end()==len(pid)): return 0

    #if not 'cid' in passport.keys(): return False
    return 1

input_lines = read_input()
parsed_lines = [parse(line) for line in input_lines]
p1_valid = sum([valid_p1(passport) for passport in parsed_lines])
p2_valid = sum([valid_p2(passport) for passport in parsed_lines])
print("p1:", p1_valid)
print("p2:", p2_valid)
