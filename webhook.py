import requests
import json

msg = {
        'username': 'RegisteringsBOT',
        'content':'''
```css
Hello World! og velkommen til Discord kanalen for {Indledende programmering}. For at få adgang til alle kanalerne her på serveren, skal du bare skrive dit studienummer til mig her. 
[Eksempel: s123456]
```
Vi ses på den anden side. :partying_face:
''',
        }
r = requests.post('https://discordapp.com/api/webhooks/749107561529081867/Lw-lRyFVmc59L9K6h2yua21oL4hCoJnFv9CZByAFG13g_5UNHbwRA-6vsS1qddzj-Tfq', json=msg)


print(r.text)

