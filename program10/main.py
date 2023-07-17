L=[]
n=int(input("Enter no of elements: "))
for i in range(n):
    x=int(input(f"Enter {i+1} element of the tuple: "))
    L.append(x)
L=tuple(L)
L11=L[-1:0:-3]
L12=L[2:10:2]
print(L11)
print(L12)
