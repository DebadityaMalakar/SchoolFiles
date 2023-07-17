l1=[]
n=int(input("How many elements do you want? "))
for i in range(n):
    x=int(input(f"Enter element {i+1} of the list: "))
    l1.append(x)
maxe=max(l1)
imax=l1.index(maxe)
if imax > len(l1)//2:
    print("Maximum element lies in second half")
elif imax < len(l1)//2:
    print("Maximum element lies in first half")
