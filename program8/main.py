vowels=['a','e','i','o','u']
s=input("Enter string: ")
x=""
for i in s:
    if i.lower() in vowels:
        pass
    else:
        x += i
print(x)
