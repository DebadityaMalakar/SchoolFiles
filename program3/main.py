num=int(input("Enter a number: "))
if num>1:
    i=2
    while i<=(num**0.5):
        if num%i==0:
            print(f"{num} is not a prime number")
            break
        i+=1
    else:
        print(f"{num} is a prime number")
else:
    print(f"{num} is not a prime number")
