import express from 'express';
import path from 'path';
import bodyParser from "body-parser";
import cookieParser from 'cookie-parser';
import axiosBase from "axios";
import {createServer} from "http";
import {Server} from "socket.io";

let cookie = require("cookie");

const session = require('express-session')
const redis = require('redis')
let RedisStore = require('connect-redis')(session)

const jsonParser = bodyParser.json();
const urlencodedParser = bodyParser.urlencoded({extended: false})

let redisClient = redis.createClient({
    host: 'session-store',
    port: 6379,
    // password: 'my secret',
    db: 1,
})
redisClient.unref()
redisClient.on('error', console.log)

let redisStore = new RedisStore({
    client: redisClient
})


const app = express();
app.use(urlencodedParser)
app.use(jsonParser)
app.use(cookieParser());
app.use(express.static(path.resolve('./', 'dist')));
app.use(session({
    store: redisStore,
    secret: ['keyboard cat', 'hey'],
    resave: false,
    saveUninitialized: true,
    cookie: {}
}))


const httpServer = createServer(app);
const io = new Server(httpServer, {path: '/api/web/', transports: ['websocket', 'polling']});


const axios = axiosBase.create({
    baseURL: 'http://backend:8080',
    headers: {
        'Content-Type': 'application/json',
        'X-Requested-With': 'XMLHttpRequest'
    },
    responseType: 'json'
});


const api = express.Router()
api.post('/auth', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    axios.post('/verify', {'account_id': ''}, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });
});

api.post('/signUp', (req, res) => {
    let data = {"account_id": req.body.account_id, "password": req.body.password}
    axios.post('/signUp', data)
        .then(function (response) {
            req.session.JWT = response.data.token
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });
});


api.post('/signIn', (req, res) => {
    let data = {"account_id": req.body.account_id, "password": req.body.password}
    console.log("sign", req.session)
    axios.post('/signIn', data)
        .then(function (response) {
            req.session.JWT = response.data.token
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });
});


api.post('/signOut', (req, res) => {
    req.session.destroy(err => {
        if (err) {
            res.send(err)
        } else {
            res.redirect('/')
        }
    })

});


const chat = express.Router()
chat.use(function (req, res, next) {
    axios.post('/').catch(() => {
        res.sendFile(path.resolve('./', 'dist', 'index.html'))
    })
})


app.post('/api/findMyRooms', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    axios.post('/api/findMyRoom', {}, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });
});


app.post('/api/findPublicRooms', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    axios.post('/api/findPublicRoom', {}, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });

});


app.post('/api/joinRoom', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    let data = {"room_id": req.body.room_id}
    axios.post('/api/joinRoom', data, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });

});


app.post('/api/createRoom', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    let data = {"room_name": req.body.room_name, "info": req.body.info, "private": req.body.private}
    axios.post('/api/createRoom', data, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });

});


app.post('/api/getMessage', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    let data = {"room_id": req.body.room_id, "message": req.body.message}
    axios.post('/api/chat/getMessage', data, config)
        .then(function (response) {
            res.send(response.data.chat_messages);
        })
        .catch(function (error) {
            res.status(401).send();
        });

});

app.post('/api/getRoomAccounts', (req, res) => {
    let token = req.session.JWT === undefined ? "" : req.session.JWT;
    let config = {headers: {Authorization: token}}
    let data = {"room_id": req.body.room_id}
    axios.post('/api/findRoomAccounts', data, config)
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.status(401).send();
        });
});


app.get('/*', function (req, res) {
    res.sendFile(path.resolve('./', 'dist', 'index.html'))
})


io.on('connection', function (socket) {
    console.log(socket.id + ' is connected.');
    socket.on("disconnect", () => {
        console.log(socket.id + ' is disconnected.');
    })

    // socket.emit()

    socket.on("join", (room_id) => {
        socket.rooms.forEach(v => {
            socket.leave(v)
        })
        socket.join(room_id)
    })


    socket.on("send", (e) => {
        let sid = cookie.parse(socket.handshake.headers.cookie)['connect.sid'];
        let sessionID = sid.split(".")[0].split(":")[1];
        redisStore.get(sessionID, (err, session) => {
            if (err == null) {
                let token = session.JWT
                let config = {headers: {Authorization: token}}
                let data = {"room_id": e['room'], "message": e['message']}
                axios.post('/api/chat/sendMessage', data, config)
                    .then(function (response) {
                        socket.to(e["room"]).emit("get", response.data.chat_message)
                    })
                    .catch(function (error) {
                        socket.on("send", (e) => {
                            socket.to(e["room"]).emit("get", {"error": 'send failed.'})
                        })
                    });
            }
        })
    })
})


app.use('/api', api)
app.use('/api/chat', chat)
httpServer.listen(3000, () => {
    console.log('server running');
})