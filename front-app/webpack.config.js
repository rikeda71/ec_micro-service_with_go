const debug = process.env.NODE_DEV !== 'production';
const webpack = require('webpack');
const path = require('path');

module.exports = {
  context: path.join(__dirname, '/src'),
  entry: ['babel-polyfill', './index.tsx'],
  module: {
    rules: [
      {
        test: /\.(j|t)sx?$/,
        exclude: /(node_modules|bower_components)/,
        use: [
          {
            loader: 'babel-loader',
            options: {
              presets: ['@babel/preset-react', '@babel/preset-typescript'],
            },
          },
        ],
      },
    ],
  },
  devServer: {
    contentBase: path.join(__dirname, 'dist'),
    compress: true,
    inline: true,
    host: '0.0.0.0',
    port: 8080,
    historyApiFallback: true,
  },
  output: {
    path: __dirname + '/src/',
    filename: 'index.min.js',
  },
  plugins: debug
    ? []
    : [
        new webpack.optimize.OccurrenceOrderPlugin(),
        new webpack.optimize.UglifyJsPlugin({
          mangle: false,
          sourceMap: false,
        }),
      ],
  resolve: {
    extensions: ['.js', '.jsx', '.ts', '.tsx'],
  },
};
