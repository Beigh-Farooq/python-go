'''def mergesort(A, left, right):
    if left < right:
        mid = (left + right) // 2
        mergesort(A, left, mid)
        mergesort(A, mid+1, right)
        merge(A, left, mid, right)

def merge(A, left, mid, right):
    i = left
    j = mid+1
    k = left
    B = [0] * (right+1)
    while i <= mid and j <= right:
        if A[i] < A[j]:
            B[k] = A[i]
            i = i + 1
        else:
            B[k] = A[j]
            j = j + 1
        k = k + 1

    while i <= mid:
        B[k] = A[i]
        i = i + 1
        k = k + 1

    while j <= right:
        B[k] = A[j]
        j = j + 1
        k = k + 1
    for x in range(left,right+1):
        A[x] = B[x]


A = [3, 5, 8, 9, 6, 2]
print('Original Array:',A)
mergesort(A,0,len(A)-1)
print('Sorted Array:',A)'''
def merge_sort(arr):
    if len(arr)<=1:
        return 
    mid=len(arr)//2
    left=arr[:mid]
    right=arr[mid:]

    merge_sort(left)
    merge_sort(right)

    merge_two_sorted_lists(left,right,arr)

def merge_two_sorted_lists(a,b,arr):
    len_a=len(a)
    len_b=len(b)
    i=j=k=0
    while i < len_a and j < len_b:
        if a[i] <= b[j]:
            arr[k]=a[i]
            i+=1
        else:
            arr[k]=b[j]
            j+=1
        k+=1
    while i< len_a:
        arr[k]=a[i]
        i+=1
        k+=1
    while j < len_b:
        arr[k]=b[j]
        j+=1
        k+=1


if __name__=='__main__':
    #arr=[10,3,15,7,8,23,98,29]
    test_cases=[
        [10,3,15,7,23,95,29],
        [],
        [3],
        [9,8,7,2],
        [1,2,3,4,5]
    ]
    for arr in test_cases:
        merge_sort(arr)
        print(arr)
