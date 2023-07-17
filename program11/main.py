L=[]
u=[]
dname=[]
n=int(input("Enter no of students: "))
for i in range(n):
    x=input(f"Enter email id of student {i+1}: ")
    while "@" not in x:
        print("Invalid email id")
        x=input("Please enter valid email id: ")
    L.append(x)
L=tuple(L)
for j in L:
    U=j.partition("@")[0]
    Dn=j.partition("@")[-1]
    u.append(U)
    dname.append(Dn)
u=tuple(u)
dname=tuple(dname)
print(u)
print(dname)
