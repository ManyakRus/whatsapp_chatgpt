Autoresponder for WhatsApp and Telegram messenger, with artificial intelligence ChatGPT
Made by Aleksander Nikitin

Can:
- automatically respond to user messages,
using artificial intelligence ChatGPT
- understands Russian (any) language
- answer in Russian (any) language
- respond on behalf of a man (woman)
- does not respond in group chats

To get started:
1. Register on the website www.openai.com
and get API_KEY there (paid),
or on the website www.proxyapi.ru (paid, for Russian users)
2. Fill in the settings in the settings.txt (or .env) file:
```
WHATSAPP_ENABLED=true
WHATSAPP_PHONE_FROM= your phone number in WhatsApp messenger,
in the format "7ХХХХХХХХХХ" for Russia
CHATGPT_API_KEY= API_KEY from point (1)
CHATGPT_NAME= name to display in chat = "Answering machine:"
CHATGPT_START_TEXT= request start text for ChatGPT
CHATGPT_END_TEXT= request end text for ChatGPT
 =" Answer in Russian on behalf of the man."
```
You can connect through a proxy proxyapi.ru (for Russian users),
if you fill in the parameters:
```
CHATGPT_PROXY_API_URL="https://api.proxyapi.ru/openai/v1"
CHATGPT_PROXY_API_KEY= API_KEY from the website proxyapi.ru
```

For use in the Telegam messenger:
```
TELEGRAM_ENABLED=true
TELEGRAM_APP_ID=
TELEGRAM_APP_HASH=
TELEGRAM_PHONE_FROM=
```
You can find your APP_ID and APP_HASH from website:
https://my.telegram.org/apps

3. Compile the code and run the file
for linux:
>make build

>./bin/whatsapp_chatgpt
4. After launching, you need to scan the QR code
from the monitor screen in WhatsApp on your phone.
Telegram will receive a message from the user "Telegram" with a code,
which must be entered into the console to confirm the connection.
5. Write a message to yourself on WhatsApp (and Telegram):
On = enable answering machine
Off = turn off answering machine


Tested on Linux Ubuntu 20.04.5, Windows 10
Readme version from: 26.06.2024

The source code is open:
https://github.com/ManyakRus/whatsapp_chatgpt