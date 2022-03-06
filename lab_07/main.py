import csv
from json import load
import matplotlib.pyplot as plt

N = 2000


class Segment:
    def __init__(self, letter, start_index, end_index=-1, freq=1):
        self.start_index = start_index
        self.letter = letter
        self.end_index = end_index
        self.freq = freq


def name_fletter_compare(rec, key):
    if rec[0] > key:
        return 1
    elif rec[0] < key:
        return -1
    else:
        return 0


def seg_comp(segment, key):
    if segment.letter > key:
        return 1
    elif segment.letter < key:
        return -1
    else:
        return 0


def brute_search(dictionary, key, comp_func):
    for i in range(len(dictionary)):
        if comp_func(dictionary[i], key) == 0:
            return dictionary[i], i + 1
    return "not found", len(dictionary)


def bin_search(dictionary, key, comp_func, left, right, is_sorted=False):
    if not is_sorted:
        dictionary.sort(key=lambda x: x[0])
    cmp = 0
    while left < right:
        mid = (left + right) // 2
        midval = dictionary[mid]
        if comp_func(midval, key) < 0:
            left = mid + 1
            cmp += 1
        elif comp_func(midval, key) > 0:
            right = mid
            cmp += 1
        else:
            cmp += 1
            return dictionary[mid], cmp
    return "not found", cmp


def make_segments(dictionary):
    segments = []
    first = '\n'
    j = -1
    #dictionary.sort(key=lambda x: x[0])
    for i in range(len(dictionary)):
        if dictionary[i][0][0] != first:
            if j != -1:
                segments[j].end_index = i - 1

            segments.append(Segment(dictionary[i][0][0], i))
            j += 1
            first = dictionary[i][0][0]
        else:
            segments[j].freq += 1
    segments.sort(key=lambda s: s.freq, reverse=True)
    return segments


def seg_search(dictionary, segments, key, comp_func):
    seg = brute_search(segments, key[0], seg_comp)
    res = bin_search(dictionary, key, comp_func, seg[0].start_index, seg[0].end_index + 1, is_sorted=True)
    return res[0], seg[1] + res[1]


def get_data(size):
    data = []
    with open('MTV.csv') as mtv:
        reader = csv.DictReader(mtv)
        cnt = 0
        for row in reader:
            name = row.get('name')
            if name != '' and name[1].isalpha():
                data.append((row.get('name')[1:-1], row.get('website')))
                cnt += 1
            if cnt == size:
                break

    return data


def test(dictionary):
    #key = "Dion DiMucci"
    key = "Aerosmith"
    res = brute_search(dictionary, key, name_fletter_compare)
    print(f"linear: res - {res[0][1]} compares - {res[1]}")

    res = bin_search(dictionary, key, name_fletter_compare, 0, len(dictionary))
    print(f"binary: res - {res[0][1]} compares - {res[1]}")

    segments = make_segments(dictionary)
    res = seg_search(dictionary, segments, key, name_fletter_compare)
    print(f"segment: res - {res[0][1]} compares - {res[1]}")


def statistic(dictionary):
    x = range(0, N)

    lin_cmps = []
    length = len(dictionary)
    for i in range(length):
        lin_cmps.append(brute_search(dictionary, dictionary[i][0], name_fletter_compare)[1])
    lin_cmps.sort()
    plt.figure(0)
    plt.title("brute search")
    plt.xlabel("dictionary length")
    plt.ylabel("compares")
    plt.hist(x, length, weights=lin_cmps, color="black")
    plt.colormaps()

    bin_cmps = []
    for i in range(length):
        bin_cmps.append(bin_search(dictionary, dictionary[i][0], name_fletter_compare, 0, length)[1])
    bin_cmps.sort()
    plt.figure(1)
    plt.title("binary search")
    plt.xlabel("dictionary length")
    plt.ylabel("compares")
    plt.hist(x, length, weights=bin_cmps, color="black")

    seg_cmps = []
    segs = make_segments(dictionary)
    for i in range(length):
        seg_cmps.append(seg_search(dictionary, segs, dictionary[i][0], name_fletter_compare)[1])
    seg_cmps.sort()
    plt.figure(2)
    plt.title("segment search")
    plt.xlabel("dictionary length")
    plt.ylabel("compares")
    plt.hist(x, length, weights=seg_cmps, color="black")
    plt.show()


if __name__ == "__main__":
    dictionary = get_data(N)
    if 0:
        #test(dictionary)
        statistic(dictionary)
    else:
        key = input("Key: ")
        res = brute_search(dictionary, key, name_fletter_compare)
        print(f"Linear: {res[0][1]} compares - {res[1]}")
        res = bin_search(dictionary, key, name_fletter_compare, 0, len(dictionary))
        print(f"Binary: {res[0][1]} compares - {res[1]}")
        segments = make_segments(dictionary)
        res = seg_search(dictionary, segments, key, name_fletter_compare)
        print(f"Segment: {res[0][1]} compares - {res[1]}")
        #statistic(dictionary)
