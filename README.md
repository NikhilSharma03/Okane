# Okane

> Okane is an CLI 💻 application build with Cobra Go🚀 to help managing your expenses 💵

## Setup and Installation

Add `Environment Variables` by creating a new `.env` file in `root` folder and add the contents following `.env.example`

Once you have added correct credentials, run the server using 

```
make run-server
```

now the server is running at `localhost:8000`

Now, Lets build the okane cli using

```
make build-cli
```

Now naviagate to `cli` folder and you will find the `okane` file which we can use as

```
./okane
```

<br>

If you want to run the server using docker, just add `redis` before port in `REDIS_ADDRESS` in `.env` file and run

```
docker-compose up --build
```

<br>

# License

<div align="center">  
<br>

<img width=35% src="https://media0.giphy.com/media/3ornjXbo3cjqh2BIyY/200.gif"></p>

<br>
</div>