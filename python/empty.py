#!/usr/bin/env python

"""A simple python script template.
"""

import sys
import argparse


def main(arguments):
    parser = argparse.ArgumentParser()
    parser.add_argument('infile', nargs='?',
                        type=argparse.FileType('r'), default=sys.stdin)
    parser.add_argument('outfile', nargs='?',
                        type=argparse.FileType('w'), default=sys.stdout)
    args = parser.parse_args(arguments)
    print(args)

if __name__ == '__main__':
    sys.exit(main(sys.argv[1:]))
