import sys


for line in sys.stdin:
    line = line.lstrip("data:")
    line = line[1:]
    line = line.rstrip("\n")
    # line = line.strip(" ")
    sys.stdout.write(line)
    sys.stdout.flush()
    # print(line, flush=True)
    # pr