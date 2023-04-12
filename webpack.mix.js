const mix = require('laravel-mix');

mix.options({ processCssUrls: false })
  .webpackConfig({
    resolve: {
      extensions: ['*', '.js', '.jsx', '.vue', '.json'],
    }
  })
  .js('resources/js/app.js', 'static/js')
  .react()
  .version()
  .sourceMaps()
  .setPublicPath('static');

