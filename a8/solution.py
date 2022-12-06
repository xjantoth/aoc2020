#!/usr/bin/env python3
import argparse
import pathlib
import time
import sys

parser = argparse.ArgumentParser()
parser.add_argument('file', type=pathlib.Path)
args = parser.parse_args()

with open(args.file) as f:
    lines = [line.rstrip('\n') for line in f]

def determine_accumulator(xlen):
    c = accumulator = 0
    x = []
    while True:
        if c >= xlen:
            print("Program terminated normally")
            print(accumulator)
            #sys.exit(0)
            return accumulator, True
        print(lines[c], c)

        o = lines[c].split(" ")[0]
        v = lines[c].split(" ")[1]

        if o == "nop":
            c += 1
        if o == "acc":
            accumulator += int(v)
            c += 1
        if o == "jmp":
            c += int(v)
        # creating list of indexes that will be repeting if paterns repeats
        x.append(c)
        # checks if freshly addes index already present 
        # in the x list of already existing list of indexes
        # if so, then print accumulator and exit :)
        if x[-1] in x[:-1]:
            print(accumulator)
            #sys.exit(1)
            # Creates Infinite loop obviously
            return accumulator, False
    #time.sleep(0.0004)

total_len = len(lines)
print(determine_accumulator(total_len))
















    #x.append(lines[c].split(" ")[1])
    #print(lines[c])

    #for i in range(len(x)):
    #    if i + 1 == len(x):
    #        break
    #    if int(x[i]) + int(x[i+1]) == 0:
    #        print(f"result: {accumulator}")
    #        #print(x)
    #        #sys.exit(1)



