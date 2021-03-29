const express = require('express');
const axiosBase = require('axios');
const port = '3000';
const host = '0.0.0.0';
const app = express();
const axios = axiosBase.create({
    baseURL: 'http://backend:8080', // バックエンドB のURL:port を指定する
    headers: {
        'Content-Type': 'application/json',
        'X-Requested-With': 'XMLHttpRequest'
    },
    responseType: 'json'
});

app.get('/', (req, res) => {
    axios.get('/ping')
        .then(function (response) {
            res.send(response.data);
        })
        .catch(function (error) {
            res.send('backend error');
        });
});

app.listen(port, host);