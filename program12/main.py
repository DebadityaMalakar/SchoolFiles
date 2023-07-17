def convert_currency(amt_in_dls,con_rate):
    con_amt=amt_in_dls*con_rate
    print(con_amt)
    return con_amt
amt=int(input("Enter amount in $: "))
cr=int(input("Enter conversion rate: "))
cnvrted=convert_currency(amt,cr)
print("₹",cnvrted)
