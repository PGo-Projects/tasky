const VuetifyLoaderPlugin = require('vuetify-loader/lib/plugin');

module.exports = {
  configureWebpack: {
    plugins: [
      new VuetifyLoaderPlugin()
    ]
  },
  pages: {
    index: {
      entry: './src/pages/Index/main.js',
      template: 'public/index.html',
      title: 'home',
      chunks: ['chunk-vendors', 'chunk-common', 'index'],
    },
    tasky: {
      entry: './src/pages/Tasky/main.js',
      template: 'public/tasky.html',
      title: 'tasky',
      chunks: ['chunk-vendors', 'chunk-common', 'tasky'],
    },
  },
};
