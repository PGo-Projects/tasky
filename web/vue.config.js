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
  },
};
