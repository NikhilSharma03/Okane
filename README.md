# Okane

<div align="center">  
<br>

<img width=100% src="https://github.com/NikhilSharma03/Okane/blob/main/assets/demo.gif"></p>

</div>

<div align="center">  
<br>
 
[![Open Source Love](https://badges.frapsoft.com/os/v1/open-source.svg?v=95)](https://github.com/NikhilSharma03/Okane) 
[![Build by Nikhil](https://img.shields.io/badge/Built%20by-Nikhil-Green)](https://github.com/NikhilSharma03/Okane)

</div>

> Okane is an CLI 💻 application build with Cobra Go🚀 to help managing your expenses 💵

## Technology Stack

<div align="center">  

<img alt="Go" src="https://img.shields.io/badge/go%20-%231572B6.svg?&style=for-the-badge&logo=go&logoColor=white"/> <img alt="gRPC" 
src="https://img.shields.io/badge/grpc%20-%231572B6.svg?&style=for-the-badge"/> <img alt="Redis" 
src="https://img.shields.io/badge/redis%20-%231572B6.svg?&color=red&style=for-the-badge&logo=redis&logoColor=white"/>
 
</div>

<br>

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