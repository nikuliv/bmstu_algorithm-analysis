class SparseMatrix:
    an = []
    nr = []
    nc = []
    jr = []
    jc = []

    def __init__(self, *args):
        if len(args) == 7:
            self.an = args[0]
            self.nr = args[1]
            self.nc = args[2]
            self.jr = args[3]
            self.jc = args[4]
            self.n = args[5]
            self.m = args[6]
        if len(args) == 2:
            self.an = []
            self.nr = []
            self.nc = []
            self.n = args[0]
            self.m = args[1]
            self.jr = [0 for _ in range(self.n)]
            self.jc = [0 for _ in range(self.m)]

    def sparse_print(self):
        spr = ''
        spr += "AN " + str(self.an) + '\n'
        spr += "NR " + str(self.nr) + '\n'
        spr += "NC " + str(self.nc) + '\n'
        spr += "JR " + str(self.jr) + '\n'
        spr += "JC " + str(self.jc) + '\n'
        return spr

    def get_element(self, i, j):
        si, sj = set(), set()
        if self.jr[i] != 0:
            if i >= 1:
                start = i
                while start - 1 >= 0:
                    if self.jr[start - 1] != 0:
                        start -= 1
                        break
                    else:
                        start -= 1
                if self.jr[start] != 0:
                    cur_el = self.nr.index(self.jr[start]) + 2
                else:
                    cur_el = 1
            else:
                cur_el = 1
            while self.jr[i] != self.nr[cur_el - 1]:
                si.add(cur_el)
                cur_el += 1
            si.add(cur_el)

        if self.jc[j] != 0:
            cur_el = self.jc[j]
            while self.jc[j] != self.nc[cur_el - 1]:
                sj.add(cur_el)
                cur_el = self.nc[cur_el - 1]
            sj.add(cur_el)
        res = si.intersection(sj)
        if len(res) != 0:
            return res.pop()
        else:
            return 0

    def matrix_print(self):
        mtr = ''
        for i in range(self.n):
            for j in range(self.m):
                ind = self.get_element(i, j)
                if ind != 0:
                    mtr += '{:^2}'.format(self.an[ind - 1]) + ' '
                else:
                    mtr += '{:^2}'.format(0) + ' '
            mtr += '\n'
        return mtr


def transform_matrix_to_rm(mtx, mtx_n, mtx_m):
    an = []
    nr = []
    nc = []
    jr = [0 for _ in range(mtx_n)]
    jc = [0 for _ in range(mtx_n)]
    cnt = 0

    for i in range(mtx_n):
        for j in range(mtx_m):
            if mtx[i][j] != 0:
                an.append(mtx[i][j])
                if jr[i] == 0:
                    jr[i] = cnt + 1
                    nr.append(cnt + 1)
                else:
                    nr.append(jr[i])
                    nr[nr.index(jr[i])] = cnt + 1
                if jc[j] == 0:
                    jc[j] = cnt + 1
                    nc.append(cnt + 1)
                else:
                    nc.append(jc[j])
                    nc[nc.index(jc[j])] = cnt + 1
                cnt += 1

    result = SparseMatrix(an, nr, nc, jr, jc, mtx_n, mtx_m)
    return result


def matrix_addition(a, b):
    n = a.n
    m = a.m
    c = [[0] * m for i in range(n)]
    for i in range(n):
        for j in range(m):
            c[i][j] = a[i][j] + b[i][j]


def sum_sparse(a: SparseMatrix, b: SparseMatrix):
    n = a.n
    m = a.m
    c = SparseMatrix(n, m)
    cnt = 0
    for i in range(n):
        for j in range(m):
            a_ind = a.get_element(i, j)
            b_ind = b.get_element(i, j)
            if a_ind != 0 and b_ind != 0:
                c.an.append(a.an[a_ind - 1] + b.an[b_ind - 1])
            elif a_ind == 0 and b_ind != 0:
                c.an.append(b.an[b_ind - 1])
            elif b_ind == 0 and a_ind != 0:
                c.an.append(a.an[a_ind - 1])
            else:
                continue
            if c.jr[i] == 0:
                c.jr[i] = cnt + 1
                c.nr.append(cnt + 1)
            else:
                c.nr.append(c.jr[i])
                c.nr[c.nr.index(c.jr[i])] = cnt + 1
            if c.jc[j] == 0:
                c.jc[j] = cnt + 1
                c.nc.append(cnt + 1)
            else:
                c.nc.append(c.jc[j])
                c.nc[c.nc.index(c.jc[j])] = cnt + 1
            cnt += 1

    return c
