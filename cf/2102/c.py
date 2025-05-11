debug = False

def dprint(*args, **kwargs):
    if debug:
        print(*args, **kwargs)

def fill(a, lowi, lowj, highi, highj, n, value):
    dprint(f"!fill(lowi={lowi}, lowj={lowj}, highi={highi}, highj={highj}, n={n}, value={value}")

    if lowi > highi:
        dprint("Early break")
        return 

    workers_1 = [
        [(lowi, lowj), lambda i, j: (i+1, j)],
        [(highi, lowj), lambda i, j: (i-1, j)],
    ]

    workers_2 = [
        [(lowi, highj), lambda i, j: (i, j-1)],
        [(lowi, lowj+1), lambda i, j: (i, j+1)],
    ]

    workers_3 = [
        [(highi, highj), lambda i, j: (i, j-1)],
        [(highi, lowj+1), lambda i, j: (i, j+1)],
    ]

    workers_4 = [
        [(lowi+1, highj), lambda i, j: (i+1, j)],
        [(highi-1, highj), lambda i, j: (i-1,j)],
    ]

    workers = [workers_1, workers_2, workers_3, workers_4]

    def handle_workers(worker, worker_num):
        w_i = -1
        while True:
            w_i = (w_i + 1) % len(worker)
            dprint(f"Worker ({w_i})[Set {worker_num}] is working!")
            dprint("Info:")
            w = worker[w_i]
            i, j = w[0]
            dprint(f"\ti={i}, j={j}")
            dprint(f"a[{i}][{j}]=a[i][j]")
            if a[i][j] >= 0:
                dprint("Assigned Before")
                dprint(f"i,j updaed to {w[1](i, j)}")
                w[0] = w[1](i, j)
                if w_i == 0:
                    dprint("Terminated because it's the first worker")
                    break
                else:
                    dprint("Skipped")
                    continue
            dprint(f"Assigned {value[0]}")
            a[i][j] = value[0]
            dprint(f"Value decreased")
            value[0] -= 1
            dprint(f"i,j updaed to {w[1](i, j)}")
            w[0] = w[1](i, j)

    for w_num in range(len(workers)):
        handle_workers(workers[w_num], w_num)

    dprint(f"Recursive call: lowi={lowi+1}, lowj={lowj+1}, highi={highi-1}, highj={highj-1}, n={n-2}, value={value}")
    fill(a, lowi+1, lowj+1, highi-1, highj-1, n-2, value)


def solve():
    n = int(input())
    a = [[-1 for j in range(n)] for i in range(n)]

    value = [n**2-1]
    fill(a, 0, 0, n-1, n-1, n, value)

    for row in a:
        for v in row:
            print(v, end=' ')
        print()


for _ in range(int(input())):
    solve()
