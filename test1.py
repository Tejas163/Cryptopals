import sys
import os
int n

def num(n):
    a=2000
    b=3200
    for i in range(b):
        if n%7==0 and n%5!=0:
            print n
        else:
            pass

def main():
    input=sys.argv
    print('enter a number',input)
    f=num(input)