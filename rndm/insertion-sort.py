import copy
def insertion_sort(arr, n):
    A = copy.deepcopy(arr)
    for i in range(1,n): # 1 instead of 2 because arrays start at 0
        key = A[i]
        # Insert A[i] into the sorted subarray A[1:i-1]
        j = i - 1
        while j >= 0 and A[j] > key: # j >= 0 instead of j > 0 because starts at 0
            A[j+1] = A[j]
            j = j - 1
            print(f'\t{A} | key = {key}')
        A[j+1] = key
        print(A)
    return A

if __name__ == "__main__":
    list_to_be_sorted = [9,8,7,6,5,4,3,2,1]
    # list_to_be_sorted = [5,9,7,6,3,1,2,8,4]
    list_length = len(list_to_be_sorted)
    sorted_list = insertion_sort(list_to_be_sorted,
                                 list_length)
    print(sorted_list)