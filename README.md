## P2PSWAP

Built with Go + Redis + Telegram for alerts

A minimalist p2p trade platform of various crypto assets. People have a crypto wallet for the various assets. They can buy/sell p2p to other users.

The system holds funds in escrow until confirmed. (They must both use their kyc verified accounts).

## Redis

### Redis Easy spin with Docker

```
docker run --name redis6888 -p 6888:6379 -d redis
```

### Redis Deamon

```bash
redis-server --daemonize yes --port 6888

redis-cli -p 6888
```
