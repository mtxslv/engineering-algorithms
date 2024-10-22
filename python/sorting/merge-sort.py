import random

def generate_random_list(n):
    a_list = [i for i in range(0,n)] # this way we guarantee no duplicate
    random.shuffle(a_list)
    assert len(a_list) == n
    return a_list

def merge_sort(arr):
    print(f'ARRIVING ARRAY:\n{arr}')
    # Base case: if the array has 1 or 0 elements, it is already sorted
    if len(arr) <= 1:
        return arr

    # Recursive case: divide the array into two halves
    mid = len(arr) // 2
    left_half = merge_sort(arr[:mid])
    right_half = merge_sort(arr[mid:])

    print('passing to merge')

    # Merge the two sorted halves
    return merge(left_half, right_half)

def merge(left, right):
    print(f'LEFT {left} | RIGHT {right}')
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

    print(f'MERGED: {merged}')
    return merged

if __name__ == "__main__":
    # Example usage:
    arr = [9,8,7,6,5,4,3,2,1]
    sorted_arr = merge_sort(arr)
    print(sorted_arr)

