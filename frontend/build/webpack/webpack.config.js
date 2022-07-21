const path = require('path');
const webpack = require('webpack');
const glob = require('glob');

const relatedResolve = (paths) => path.resolve(__dirname, paths);

// Get entry files in src/packs
const webpackEntries = () => {
  const entries = {};
  const packsFolder = relatedResolve('../../src/packs');

  glob.sync(`${packsFolder}/**/*.{js,ts}`).map((file) => {
    const filePath = path.parse(file);
    entries[filePath.name] = relatedResolve(file);
  });

  return entries;
};

module.exports = {
  mode: 'development',
  entry: webpackEntries(),
  resolve: {
    extensions: ['.ts', '.tsx', '.js', '.jsx'],
    fallback: {
      crypto: require.resolve('crypto-browserify'),
      stream: require.resolve('stream-browserify'),
      assert: require.resolve('assert'),
      http: require.resolve('stream-http'),
      https: require.resolve('https-browserify'),
      os: require.resolve('os-browserify'),
      url: require.resolve('url'),
    },
  },
  stats: { children: false },
  target: 'web',
  output: {
    filename: '[name].js',
    path: relatedResolve('dist'),
    library: {
      type: 'module',
    },
  },
  module: {
    rules: [
      {
        test: /\.(js|jsx|mjs|ts|tsx)$/,
        exclude: /\.test.tsx?$/,
        loader: 'babel-loader',
        options: {
          presets: [
            '@babel/preset-env',
            ['@babel/preset-typescript', { allExtensions: true, isTSX: true }],
          ],
        },
      },
    ],
  },
  experiments: {
    outputModule: true,
  },
  plugins: [
    new webpack.ProvidePlugin({
      process: 'process/browser',
      Buffer: ['buffer', 'Buffer'],
    }),
    new webpack.ProgressPlugin({
      activeModules: false,
      entries: true,
      modules: true,
      modulesCount: 5000,
      profile: false,
      dependencies: true,
      dependenciesCount: 10000,
      percentBy: null,
    }),
  ],
};
