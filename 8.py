import numpy as np

def main():
    fo = open("input/8.txt")
    lines = fo.read().split("\n")

    width, height = 50, 6
    grid = np.zeros((height, width), dtype=bool)
    #print(grid)

    for line in lines:
        tokens = line.split(' ')
        action = tokens[0]
        if action == "rect":
            w, h = map(int, tokens[1].split("x"))
            grid[:h, :w] = True
        elif action == "rotate":
            idx = int(tokens[2][2:])
            by = int(tokens[-1])
            if tokens[1] == "row":
                grid[idx] = np.roll(grid[idx], by)
            elif tokens[1] == "column":
                grid[:,idx] = np.roll(grid[:,idx], by)

    print(sum(sum(grid)))
    for i in range(height):
        for j in range(width):
            print('*' if grid[i][j] else ' ', end="")
        print()
    # parr(grid)
def parr(arr):
    for l in arr:
        print(l)

main()
