const { merge } = require('webpack-merge');
const webpackConfig = require('./webpack.config.js');

module.exports = merge(webpackConfig, {
    mode: 'development',
    devServer: {
        historyApiFallback: true,
        inline: true,
        open: true,
        host: '0.0.0.0',
        port: 8000,
        proxy: {
            '/api/*': {
                target: 'http://localhost:3000',
                secure: false,
                logLevel: 'debug'
            }
        },
    }
})