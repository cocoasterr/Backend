import websockets
import asyncio

async def listen():
    url = "ws://127.0.0.1:7890"
    async with websockets.connect(url) as ws:
        while True:
            message = input("insert your bc msg: ")
            await ws.send(f"Admin: {message}")

asyncio.get_event_loop().run_until_complete(listen())

