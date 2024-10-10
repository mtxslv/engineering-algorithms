import copy
from textwrap import indent
import time
import json
import random

from matplotlib.font_manager import json_dump

HOW_MANY_TRIALS = 30
MIN_LENGTH = 10
MAX_LENGTH = 1500

def generate_random_list(n):
    a_list = [i for i in range(0,n)] # this way we guarantee no duplicate
    random.shuffle(a_list)
    assert len(a_list) == n
    return a_list


def insertion_sort(arr, n):
    ''' 
    ALGORITHM DETAILS: 
    - for stars at 1 (instead of 2) because python arrays start at 0
    - on the while loop, j >= 0 instead of j > 0 because starts at 0
    '''
    A = copy.deepcopy(arr)
    num_of_operations = 0
    time_when_start = time.time_ns()
    for i in range(1,n):  # TWO OPERATIONS ON LOOP: ADD STEP (IN THIS CASE ONE) AND COMPARE IF LIMIT WAS REACHED
        key = A[i] # TWO OPERATIONS HERE: RETRIEVE AND ATTRIBUTION
        j = i - 1 # TWO OPERATIONS HERE: SUBTRACTION AND ATTRIBUTION
        while j >= 0 and A[j] > key: # THREE OPERATIONS: THE TWO COMPARISONS AND THE LOGICAL CONJUNCTION
            A[j+1] = A[j] # FOUR OPERATIONS: SUM, THE ATTRIBUTION AND THE INDEXATION
            j = j - 1 # TWO OPERATIONS: SUBTRACTION AND ATTRIBUTION
            num_of_operations += 11 # previous three lines and this line cost 
        A[j+1] = key # TWO OPERATIONS: ADDITION, INDEXING AND ATTRIBUTION
    time_when_finish = time.time_ns()
    time_it_took = time_when_finish - time_when_start
    num_of_operations += n*8
    return A, time_it_took, num_of_operations

if __name__ == "__main__":

    times_ = dict()
    operations_ = dict()
    for list_length in range(MIN_LENGTH, MAX_LENGTH):
        print(f'List length: {list_length}')
        times_of_execution = []
        quantity_of_ops = []
        for trial in range(HOW_MANY_TRIALS):

            list_to_be_sorted = generate_random_list(list_length) #[9,8,7,6,5,4,3,2,1]
            sorted_list, time_to_run, num_of_operations = insertion_sort(list_to_be_sorted, list_length)
            
            
            times_of_execution.append(time_to_run)
            quantity_of_ops.append(num_of_operations)

        times_[list_length] = times_of_execution
        operations_[list_length] = quantity_of_ops
    with open('times.json', 'w+') as file:
        json.dump( times_, file, indent=4)
    with open('ops.json', 'w+') as file:
        json.dump( operations_, file, indent=4)


        # print(sorted_list)