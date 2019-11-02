import path from 'path'
import HtmlWebpackPlugin from 'html-webpack-plugin'

const src = path.resolve(__dirname, 'src');
const dist = path.resolve(__dirname, 'dist');

export default {
    mode: 'development',
    entry: src + '/index.jsx',

    output: {
        path: dist,
        filename: 'bundle.js'
    },

    module: {
        rules: [
            {
                test: /\.jsx$/,
                use: 'babel-loader'
            },
            {
                test: /\.css$/i,
                use: ['style-loader', 'css-loader']
            },
            {
                test: /\.(woff|woff2|eot|ttf|svg)$/,
                loader: 'file-loader'
            }
        ],
    },

    resolve: {
        extensions: ['.js', '.jsx']
    },

    plugins: [
        new HtmlWebpackPlugin({
            template: src + '/index.html',
            filename: 'index.html'
        })
    ]
}