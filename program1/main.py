n=float(input("Enter number: "))
if((n>0 and n <10) or (n<0 and n>-10)):
    print(f"{n} is a 1-digit number")
elif((n>=10 and n <=99) or (n<=-10 and n>=-99)):
    print(f"{n} is a 2-digit number")
elif((n>99 and n<=999) or (n<-99 and n>= -999)):
    print(f"{n} is a 3-digit number")
else:
    print("Invalid Case")
