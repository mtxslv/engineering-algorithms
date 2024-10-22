import copy
import json
import time

MIN_LENGTH = 10
MAX_LENGTH = 1000


def insertion_sort(A, n):
    for i in range(1,n): # 1 instead of 2 because arrays start at 0
        key = A[i]
        # Insert A[i] into the sorted subarray A[1:i-1]
        j = i - 1
        while j >= 0 and A[j] > key: # j >= 0 instead of j > 0 because starts at 0
            A[j+1] = A[j]
            j = j - 1
            # print(f'\t{A} | key = {key}')
        A[j+1] = key
    return A

def merge_sort(arr):
    # Base case: if the array has 1 or 0 elements, it is already sorted
    if len(arr) <= 1:
        return arr

    # Recursive case: divide the array into two halves
    mid = len(arr) // 2
    left_half = merge_sort(arr[:mid])
    right_half = merge_sort(arr[mid:])

    # Merge the two sorted halves
    return merge(left_half, right_half)

def merge(left, right):
    merged = [0] * (len(left) + len(right))  # Preallocate array
    left_index = right_index = merged_index = 0

    # Compare elements from left and right arrays, and merge them in sorted order
    while left_index < len(left) and right_index < len(right):
        if left[left_index] < right[right_index]:
            merged[merged_index] = left[left_index]
            left_index += 1
        else:
            merged[merged_index] = right[right_index]
            right_index += 1
        merged_index += 1

    # Copy any remaining elements from the left array
    while left_index < len(left):
        merged[merged_index] = left[left_index]
        left_index += 1
        merged_index += 1

    # Copy any remaining elements from the right array
    while right_index < len(right):
        merged[merged_index] = right[right_index]
        right_index += 1
        merged_index += 1

    return merged

def generate_ordered_sequence(n, ord):
    if ord == 'best':
        seq = [i for i in range(n)]  # increasing is best case: [0,1,2,3,4,5,6,7,8]
    elif ord == 'mid-best':  # [0,1,2,3,4,8,7,6,5]
        mid = n // 2
        seq = [i for i in range(mid)] + [i for i in range(n-1, mid-1, -1)]
    elif ord == 'mid-worst':  # [8,7,6,5,4,0,1,2,3]
        mid = n // 2
        seq = [i for i in range(n-1, mid-1, -1)] + [i for i in range(mid)]
    else:
        seq = [i for i in range(n-1, -1, -1)]  # decreasing is worst case: [8,7,6,5,4,3,2,1,0]
    
    return seq

if __name__ == "__main__":
    tempos_ = {
        'best':{
            'n': [],
            'insertion': [],
            'merge': []
        },
        'mid-best':{
            'n': [],
            'insertion': [],
            'merge': []
        },
        'mid-worst':{
            'n': [],
            'insertion': [],
            'merge': []
        },
        'worst':{
            'n': [],
            'insertion': [],
            'merge': []
        },
    }
    for tipo_execucao in tempos_.keys():
        print(f'Rodando para {tipo_execucao}')
        for n in range(MIN_LENGTH, MAX_LENGTH+1):
            # BEST CASE

            # Sequence (copy to prevent in-place sorting)
            arr = generate_ordered_sequence(n, tipo_execucao)
            arr_insertion = copy.deepcopy(arr)
            arr_merge = copy.deepcopy(arr)

            # Insertion sort
            start_insertion = time.time_ns()
            insertion_sorted = insertion_sort(arr_insertion, n)
            finish_insertion = time.time_ns()

            # Merge sort
            start_merge = time.time_ns()
            merge_sorted = merge_sort(arr_merge)
            finish_merge = time.time_ns()

            tempos_[tipo_execucao]['n'].append(n)
            tempos_[tipo_execucao]['insertion'].append(finish_insertion - start_insertion)
            tempos_[tipo_execucao]['merge'].append(finish_merge - start_merge)

    with open('comparison.json', 'w+') as file:
        json.dump(tempos_, file, indent=4)

    # print("Best case:      ", generate_ordered_sequence(n, 'best'))
    # print("Mid-best case:  ", generate_ordered_sequence(n, 'mid-best'))
    # print("Mid-worst case: ", generate_ordered_sequence(n, 'mid-worst'))
    # print("Worst case:     ", generate_ordered_sequence(n, 'worst'))
