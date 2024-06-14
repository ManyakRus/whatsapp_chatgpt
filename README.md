Autoresponder for WhatsApp messenger, with artificial intelligence ChatGPT
Made by Aleksandr Nikitin
Skype login: Travianbot

Can:
- automatically reply to user messages,
using artificial intelligence ChatGPT
- understands Russian (any) language
- answer in Russian (any) language
- answer on behalf of a man (woman)
- does not respond in group chats

To get started:
1. Register at www.openai.com
and get API_KEY there (free)
2. Fill in the settings in the .env file:
WHATSAPP_PHONE_FROM - your phone number in whatsapp messenger,
in the format "7ХХХХХХХХХ" for Russia
CHATGPT_API_KEY - API_KEY from item (1)
CHATGPT_NAME - name to display in chat = "Autoresponder:"
CHATGPT_START_TEXT - ChatGPT start request text
CHATGPT_END_TEXT - ChatGPT request end text
  =" Answer in Russian on behalf of the man."

You can connect through a proxy proxyapi.ru (for Russian users),
if you fill in the parameters:
CHATGPT_PROXY_API_URL="https://api.proxyapi.ru/openai/v1"
CHATGPT_PROXY_API_KEY=


3. Compile the code and run the file
whatsapp_chatgpt
4. After starting, you need to scan the QR code
from the monitor screen in WhatsApp on the phone
5. Write a WhatsApp message to yourself:
On = turn on auto answer
Off = disable auto answer

Tested on Linux Ubuntu 20.04.5
Readme version from: 03/06/2022

Source code open:
https://github.com/ManyakRus/whatsapp_chatgpt