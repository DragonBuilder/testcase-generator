import sys


for line in sys.stdin:
    line = line.lstrip("data: ")
    line = line.rstrip("\n\n")
    line = line.lstrip("[")
    line = line.rstrip("]")
    # line = line.strip(" ")
    # sys.stdout.write(line)
    # sys.stdout.flush()
    print(line, end="", flush=True)
    # print(line, flush=True)
    # pr