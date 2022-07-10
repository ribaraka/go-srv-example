const path = require('path');
const HtmlWebpackPlugin = require('html-webpack-plugin');
const {CleanWebpackPlugin} = require('clean-webpack-plugin');
let mode = "development";
if (process.env.NODE_ENV === "production") {
  mode = "production";
}

module.exports = {
  context: path.resolve(__dirname, 'src'),
  mode: mode,
  devtool: 'inline-source-map',
  devServer: {
    publicPath: '/',
    port: 9000,
    contentBase: path.join(__dirname, 'form'),
    host: '127.0.0.1',
    hot: true,
  },
  entry: './index.js',
  module: {
    rules: [
      {
        test: /\.js$/,
        exclude: /node_modules/,
        use: {
          loader: 'babel-loader',
        }
      },
      {
        test: /\.(png|jpg|svg|gif)$/,
        use: ['file-loader']
      },
      {
        test: /\.css$/i,
        use: ['style-loader', 'css-loader'],
      },
      {
        test: /\.s[ac]ss$/i,
        use: [
          // Creates `style` nodes from JS strings
          "style-loader",
          // Translates CSS into CommonJS
          "css-loader",
          // Compiles Sass to CSS
          "sass-loader",
        ],
      },
    ],
  },
  output: {
    path: path.resolve(__dirname, '../../cmd/form'),
    filename: '[name].[hash].bundle.js',
  },
  plugins: [
    new HtmlWebpackPlugin({
      filename: "index.html",
      template: './index.html',
    }),
    new HtmlWebpackPlugin({
      filename: "result.html",
      template: './result.html'
    }),
    new CleanWebpackPlugin(),
  ],
};
