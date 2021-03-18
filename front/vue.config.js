module.exports = {
  pages: {
    'editor': {
      entry: './src/pages/editor/main.js',
      template: 'public/index.html',
      title: 'Editor',
      chunks: [ 'chunk-vendors', 'chunk-common', 'editor' ]
    },
    'print': {
      entry: './src/pages/print/main.js',
      template: 'public/index.html',
      title: 'Print',
      chunks: [ 'chunk-vendors', 'chunk-common', 'print' ]
    }
  }
}