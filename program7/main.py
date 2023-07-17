email_id=input("Enter email id: ")
p1,p2,p3=email_id.partition("@")
if not (p1) and not(p2) and not(p3):
    print("No data or @ does not exist")
elif p1 and not(p2) and p3:
    print("Invalid")
elif p1 and not(p3=="edupillar.com"):
    print("Domain does not match")
elif not(p1) and (p3=="edupillar.com"):
    print("Only Domain")
elif p1 and (p3=="edupillar.com"):
    print("Domain Matches")
elif not(p1) and p2 and not(p3):
    print("Invalid")
