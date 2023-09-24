import json

def fibonacci(n:int) -> list:
    res = [0,1]
    for i in range(n):
        res.append(res[-1] + res[-2])
    return res

def palind(data: str) -> str:
    rev = ""
    for x in range(len(data)):
        rev+= data[len(data)-1-x]

    # rev = data[::-1]
    if data == rev:
        return "ist Palindrome!"
    return "its not palindrome"

def sort(data:list):
    n = len(data)
    for d in range(n):
        for x in range(n- 1 - d):
            if data[x] < data[x + 1]:
                data[x],data[x+1] = data[x+1], data[x]
    return data

def odd_even(n:int):
    return [f"{x} - even " if x %2 ==0 else f"{x} - odd " for x in range(n+1)]

def fizzbuzz(n:int):
    return [f"{x} - fizzbuzz" if x%3==0 and x%5 == 0 else f"{x} - fizz" if x%3 == 0 else f"{x} - buzz" if x%5 == 0 else x for x in range(n+1)]
n = 20
res_fibo = fibonacci(n)
print(f"Fibonacci {n}\n{res_fibo}")

data = "kasur rusak"
palin_check = palind(data)
print(f"palin check {data}\n{palin_check}")

data=[0,2,1,23,1,7,2,2]
sort = sort(data)
print(sort)

print(odd_even(10))
print(fizzbuzz(20))

country_json = '{ "Country Name": "USA","States": {"Ohio":52343, "Texas":44554,"California":21233},"Crops":{"Crop 1":["Maize","tobacco","wheat"],"Crop 2":["Paddy","Chillies"]}}'

with open("data.json", "r") as f:
    # f.write(country_json)
    country_dict = json.load(f)
    print(country_dict['States']['Texas'])
    print(type(country_dict))