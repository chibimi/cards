module.exports = {
  publicPath: 'dist',
  pages: {
    'editor': {
      entry: './src/pages/editor/main.js',
      template: 'public/index.html',
      title: 'Editor',
      filename: 'editor/index.html',
      chunks: [ 'chunk-vendors', 'chunk-common', 'editor' ]
    },
    'print': {
      entry: './src/pages/print/main.js',
      template: 'public/index.html',
      title: 'Print',
      filename: 'print/index.html',
      chunks: [ 'chunk-vendors', 'chunk-common', 'print' ]
    }
  }
}