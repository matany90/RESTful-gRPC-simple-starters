const nodeExternals = require("webpack-node-externals")

// base webpack configurations
let baseConfig = {
  // entry file
  entry: [__dirname, "./src/index.js"],

  // output
  output: {
    path: __dirname + "/dist",
    filename: "index.js",
    libraryTarget: "commonjs2"
  },

  // target build
  target: "node",

  // externals
  externals: [nodeExternals()],
}

if (process.env.NODE_ENV === "production") {
  // add production optimizations
  baseConfig = Object.assign(baseConfig, {
    // set mode
    mode: "production",

    // modules
    module: {
      rules: [
        {
          test: /\.js$/,
          use: "babel-loader",
          exclude: /node_modules/
        }
      ]
    }
  })
} else {
  // add development tools
  baseConfig = Object.assign(baseConfig, {
    // set mode
    mode: "development",

    // add source map
    devtool: "source-map",

    // modules
    module: {
      rules: [
        {
          test: /\.js$/,
          use: ["babel-loader", "eslint-loader"],
          exclude: /node_modules/
        }
      ]
    }
  })
}

module.exports = baseConfig