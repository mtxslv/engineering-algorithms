import copy
import json
import random


def generate_random_list(n):
    a_list = [i for i in range(0,n)] # this way we guarantee no duplicate
    random.shuffle(a_list)
    assert len(a_list) == n
    return a_list

def generate_ordered_sequence(n, ord):
    if ord == 'best':
        # increasing is best case: [0,1,2,3,4,5,6,7,8]
        seq = [i for i in range(n)]  
    else:
        # decreasing is worst case: [8,7,6,5,4,3,2,1,0]
        seq = [i for i in range(n-1, -1, -1)]  
    return seq
    
def worker_assignment(task_list, qtd_workers = 4):
    k = qtd_workers
    tasks_per_worker = len(task_list) // k
    remainder_tasks = len(task_list) % k
    # Split tasks among workers
    load_distribution = []
    start = 0
    for i in range(k):
        end = start + tasks_per_worker + (1 if i < remainder_tasks else 0)
        load_distribution.append(
            sum(task_list[start:end])
        )
        # result.append(task_list[start:end])
        start = end
    
    return load_distribution    

def main():
    pass

if __name__ == "__main__":
    loads_ = dict()
    for n in range(2,100+1):
        # Generate task list
        task_list = [i for i in range(1,4*n+1)]

        # Shuffle task list
        shuffled_task_list = copy.deepcopy(task_list)
        random.shuffle(shuffled_task_list)
        
        loads_[n] = {
            'shuffled': worker_assignment(shuffled_task_list,4),
            'original': worker_assignment(task_list, 4)
        }
    with open('load-balancer-problem/loads.json', 'w+') as file:
        json.dump(loads_, file, indent=4)