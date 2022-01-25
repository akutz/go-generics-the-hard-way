#!/usr/bin/env python3

"""
Copyright 2022

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
"""

"""b2md

Transforms Go benchmark output into a markdown table.

    b2md.py FILE
    cat FILE | b2md.py
"""

import argparse
import re
import sys


def intOrFloat(s):
    if "." in s:
        return float(s)
    return int(s)


def twoDecimalFloat(f):
    return float("{:.2f}".format(f))


def noZedZed(f):
    c = str(f)
    if "." in c:
        if c.endswith(".0"):
            return int(c.replace(".0", ""))
        return twoDecimalFloat(float(c))
    return int(c)


def avgNoZedZed(a, b):
    return noZedZed(a / b)


def avgBoxingData(listType, data):
    ops = 0
    nsOp = 0
    bytesOp = 0
    allocsOp = 0

    for e in data[listType]:
        ops += e["ops"]
        nsOp += e["nsOp"]
        bytesOp += e["bytesOp"]
        allocsOp += e["allocsOp"]

    return {
        "ops": avgNoZedZed(ops, len(data[listType])),
        "nsOp": avgNoZedZed(nsOp, len(data[listType])),
        "bytesOp": avgNoZedZed(bytesOp, len(data[listType])),
        "allocsOp": avgNoZedZed(allocsOp, len(data[listType])),
    }


def printBoxingMarkdown(listType, data):
    print(
        "| {} | 1 | {} | {} | {} | {} |".format(
            listType,
            data["ops"],
            data["nsOp"],
            data["bytesOp"],
            data["allocsOp"],
        )
    )


def printBuildTimeMarkdown(data):
    s = "| {} | {} | {} | {} | {} | {} | {} | {} | {} | {} |"

    def printFileType(file_type):
        for i in range(0, 6):
            ft = file_type
            if i > 0:
                ft = ""
            tOps = data["typed"][file_type][i]["ops"]
            gOps = data["generic"][file_type][i]["ops"]
            tNsOp = data["typed"][file_type][i]["nsOp"]
            gNsOp = data["generic"][file_type][i]["nsOp"]
            print(
                s.format(
                    ft,
                    i,
                    noZedZed(tOps),
                    noZedZed(gOps),
                    noZedZed(gOps - tOps),
                    noZedZed(-(((tOps - gOps) / tOps) * 100)),
                    noZedZed(tNsOp),
                    noZedZed(gNsOp),
                    noZedZed(gNsOp - tNsOp),
                    noZedZed(-((tNsOp - gNsOp) / tNsOp) * 100),
                )
            )

    printFileType("pkg")
    printFileType("bin")


def printFileSizeMarkdown(data):
    s = "| {} | {} | {} | {} | {} | {} |"

    def printFileType(file_type):
        for i in range(0, 6):
            ft = file_type
            if i > 0:
                ft = ""
            t = data["typed"][file_type][i]["bytesOp"]
            g = data["generic"][file_type][i]["bytesOp"]
            print(
                s.format(
                    ft,
                    i,
                    noZedZed(t),
                    noZedZed(g),
                    noZedZed(g - t),
                    noZedZed(-(((t - g) / t) * 100)),
                )
            )

    printFileType("pkg")
    printFileType("bin")


parser = argparse.ArgumentParser(
    description="b2md", formatter_class=argparse.ArgumentDefaultsHelpFormatter
)


parser.add_argument(
    "-t",
    "--type",
    dest="type",
    nargs=1,
    default="boxing",
    choices=["boxing", "buildtime", "filesize"],
    help="type of input",
)

parser.add_argument(
    "--no-echo",
    dest="no_echo",
    action="store_true",
    help="do not echo stdin to stdout",
)

parser.add_argument(
    "file",
    metavar="FILE",
    nargs="?",
    help="path to an input file",
)

args = parser.parse_args()
if not args:
    parser.print_help()
    sys.exit(1)


if args.file:
    echo = False
    f = open(args.file, "r")
else:
    echo = not args.no_echo
    f = sys.stdin


rx = re.compile(
    r"(Benchmark[^\s]+-\d+)\s+(\d+)\s+([\d\.]+) ns/op\s+([\d\.]+) (?:B|filesize)/op(?:\s+([\d\.]+) allocs/op)?"
)

if args.type == "boxing":
    """
    data = {
        "boxed": [
            {
                "ops": 0,
                "nsOp": 0,
                "bytesOp": 0,
                "allocsOp": 0,
            },
        ],

        # same as boxed
        "generic": [],

        # same as boxed
        "typed": [],
    """
    data = {}
else:
    """
    data = {
        "boxed": {
            "pkg": {
                "0": {
                    "ops": 0,
                    "nsOp": 0,
                    "bytesOp": 0,
                },
                # repeat for 1-5
            },
            # same as pkg
            "bin": {},
        },

        # same as boxed
        "generic": {},

        # same as boxed
        "typed": {},
    """
    data = {}


if args.type == ["boxing"]:
    args.type = "boxing"
elif args.type == ["buildtime"]:
    args.type = "buildtime"
elif args.type == ["filesize"]:
    args.type = "filesize"


for line in f:
    if echo:
        print(line, end="")

    m = rx.match(line)
    if not m:
        continue

    name = m.group(1)
    ops = intOrFloat(m.group(2))
    nsOp = intOrFloat(m.group(3))
    bytesOp = intOrFloat(m.group(4))

    if args.type == "boxing":
        allocsOp = intOrFloat(m.group(5))

    if "/boxed" in name:
        list_type = "boxed"
    elif "/generic" in name:
        list_type = "generic"
    elif "/typed" in name:
        list_type = "typed"
    else:
        sys.exit("unrecognized list type")

    if args.type == "boxing":
        if list_type not in data:
            data[list_type] = []

        data[list_type].append(
            {
                "ops": ops,
                "nsOp": nsOp,
                "bytesOp": bytesOp,
                "allocsOp": allocsOp,
            }
        )
    else:
        if list_type not in data:
            data[list_type] = {}

        if "pkg/" in name:
            file_type = "pkg"
        elif "bin/" in name:
            file_type = "bin"
        else:
            sys.exit("unrecognized file type")

        if file_type not in data[list_type]:
            data[list_type][file_type] = {}

        if "/empty_interface" in name:
            num_types = -1
        elif "/0-types" in name:
            num_types = 0
        elif "/1-types" in name:
            num_types = 1
        elif "/2-types" in name:
            num_types = 2
        elif "/3-types" in name:
            num_types = 3
        elif "/4-types" in name:
            num_types = 4
        elif "/5-types" in name:
            num_types = 5
        else:
            sys.exit("unrecognized number of types")

        if num_types not in data[list_type][file_type]:
            data[list_type][file_type][num_types] = {}

        data[list_type][file_type][num_types]["ops"] = noZedZed(ops)
        data[list_type][file_type][num_types]["nsOp"] = noZedZed(nsOp)
        data[list_type][file_type][num_types]["bytesOp"] = noZedZed(bytesOp)

# Go ahead and close the file.
f.close()

if args.type == "boxing":
    avgBoxed = avgBoxingData("boxed", data)
    avgGeneric = avgBoxingData("generic", data)
    avgTyped = avgBoxingData("typed", data)

    print("| List type | Number of types | Operations | ns/op | Bytes/op | Allocs/op |")
    print("|:---------:|:---------------:|:----------:|:-----:|:--------:|:---------:|")
    printBoxingMarkdown("Boxed", avgBoxed)
    printBoxingMarkdown("Generic", avgGeneric)
    printBoxingMarkdown("Typed", avgTyped)

elif args.type == "buildtime":
    print(
        "| Artifact type | Number of types | Ops - typed | Ops - generic | Increase (ops) | Increase (%) | ns/op - typed | ns/op - generic | Increase (ns/op) | Increase (%) |"
    )
    print(
        "|:-------------:|:---------------:|:-----------:|:-------------:|:--------------:|:------------:|:-------------:|:---------------:|:----------------:|:------------:|"
    )
    printBuildTimeMarkdown(data)
elif args.type == "filesize":
    print(
        "| Artifact type | Number of types | File size (bytes) - typed | File size (bytes) - generic | Increase (bytes) | Increase (%) |"
    )
    print(
        "|:-------------:|:---------------:|:-------------------------:|:---------------------------:|:----------------:|:------------:|"
    )
    printFileSizeMarkdown(data)
