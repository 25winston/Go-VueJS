let mix = require('laravel-mix')

// Load Plugins
const VueLoaderPlugin = require('vue-loader/lib/plugin')

const SWPrecacheWebpackPlugin = require('sw-precache-webpack-plugin');
const BundleAnalyzerPlugin = require('webpack-bundle-analyzer').BundleAnalyzerPlugin;

const package = require('./package.json');
const dependencies = Object.keys(package.dependencies);

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for your application, as well as bundling up your JS files.
 |
 */



// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// //
// //
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// mix.options({ 
//   imgLoaderOptions: { enabled: false }, 
//   setResourceRoot: path.normalize('./'),
//   publicPath: path.normalize('./public')
// })

// // .copyDirectory('./app/resources/assets/img', './public/assets/img')
// .copyDirectory('./node_modules/bootstrap/dist/fonts', './public/assets/fonts')
// .copy('./app/resources/static/index.html', './public/index.html')
// .sass('./app/resources/assets/sass/app.scss', '../app/resources/assets/css/')  // target path: `publicPath`
// .styles([
//   './app/resources/assets/css/main.css', 
//   './app/resources/assets/css/app.css', 
//   './node_modules/bootstrap/dist/css/bootstrap.css',
//   './node_modules/bootstrap/dist/css/bootstrap-theme.css',
// ], './public/assets/css/all.css')
// // .react('./app/resources/reactjs/app.js', './public/assets/js')
// .js('./app/resources/vuejs/app.js', './public/assets/js')
// // .minify()
// // .version()
// .webpackConfig({
//   plugins: [
//     new VueLoaderPlugin()
//   ]
// })




// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// //
// //
// ////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
mix
// .extract(dependencies)
// .autoload({
//   jquery: ['$', 'window.jQuery', 'jQuery'],
//   lodash: ['_', 'window._'],
//   vue: ['Vue', 'window.Vue'],
//   axios: ['axios', 'window.axios'],
//   "js-cookie": ['Cookies', 'window.Cookies'],
// })
.js('./app/resources/vuejs/app.js', './public/assets/js')
.sass('./app/resources/assets/sass/app.scss', 'assets/css/')  // target path: `publicPath`
.copyDirectory('./app/resources/assets/images', './public/assets/images')
.copyDirectory('./app/resources/assets/fonts', './public/assets/fonts')
// .copyDirectory('./node_modules/slick-carousel/slick/fonts', './public/assets/fonts')
.copyDirectory('./node_modules/@fortawesome/fontawesome-free/sprites', './public/assets/sprites')
.copyDirectory('./node_modules/@fortawesome/fontawesome-free/svgs', './public/assets/svgs')
.copyDirectory('./node_modules/@fortawesome/fontawesome-free/webfonts', './public/assets/webfonts')
// .copy('./node_modules/slick-carousel/slick/ajax-loader.gif', './public/assets/css/ajax-loader.gif')
.copy('./app/resources/static/index.html', './public/index.html')
.styles([

  // './app/resources/assets/css/bootstrap.min.css', 
  // './app/resources/assets/css/bootstrap-reset.css', 
  // './app/resources/assets/css/font-awesome.css', 
  // './app/resources/assets/css/style.css', 
  // './app/resources/assets/css/style-responsive.css', 

  './public/assets/css/app.css', 
], './public/assets/css/all.css')
// .react('./app/resources/reactjs/app.js', './public/assets/js')
.options({
  imgLoaderOptions: { enabled: false },
  setResourceRoot: path.normalize('./'),
  publicPath: path.normalize('./public'),
  // extractVueStyles: true,
  // processCssUrls: true,
  // uglify: {},
  // purifyCss: false,
  // //purifyCss: {},
  // postCss: [require('autoprefixer')],
  // clearConsole: false
})
.webpackConfig({
  resolve: {
    alias: {
      '@': path.resolve(__dirname, 'app/resources/vuejs'),
      '@JS': path.resolve(__dirname, 'app/resources/assets/js') 
    }
  },
  plugins: [
    new VueLoaderPlugin(),
    // new SWPrecacheWebpackPlugin({
    //   cacheId: 'pwa',
    //   filename: 'assets/js/service-worker.js',
    //   staticFileGlobs: ['public/**/*.{css,eot,svg,ttf,woff,woff2,js,html}'],
    //   minify: true,
    //   stripPrefix: 'public/',
    //   handleFetch: true,
    //   dynamicUrlToDependencies: {
    //     // '/': ['resources/views/welcome.blade.php'],
    //   },
    //   staticFileGlobsIgnorePatterns: [/\.map$/, /mix-manifest\.json$/, /manifest\.json$/, /service-worker\.js$/],
    //   runtimeCaching: [{
    //     urlPattern: /^https:\/\/fonts\.googleapis\.com\//,
    //     handler: 'cacheFirst'
    //   }],
    //   // importScripts: ['./js/push_message.js']
    // }),
    new BundleAnalyzerPlugin({
      analyzerMode: 'static',
      reportFilename: path.join(`${__dirname}/public`, 'webpack-report.html'),
      openAnalyzer: false,
      logLevel: 'silent'
    }),
  ]
})
.sourceMaps(!mix.inProduction())
.disableNotifications()
// .minify()
// .version()


// Full API
// mix.js(src, output);
// mix.react(src, output); <-- Identical to mix.js(), but registers React Babel compilation.
// mix.preact(src, output); <-- Identical to mix.js(), but registers Preact compilation.
// mix.coffee(src, output); <-- Identical to mix.js(), but registers CoffeeScript compilation.
// mix.ts(src, output); <-- TypeScript support. Requires tsconfig.json to exist in the same folder as webpack.mix.js
// mix.extract(vendorLibs);
// mix.sass(src, output);
// mix.standaloneSass('src', output); <-- Faster, but isolated from Webpack.
// mix.fastSass('src', output); <-- Alias for mix.standaloneSass().
// mix.less(src, output);
// mix.stylus(src, output);
// mix.postCss(src, output, [require('postcss-some-plugin')()]);
// mix.browserSync('my-site.test');
// mix.combine(files, destination);
// mix.babel(files, destination); <-- Identical to mix.combine(), but also includes Babel compilation.
// mix.copy(from, to);
// mix.copyDirectory(fromDir, toDir);
// mix.minify(file);
// mix.sourceMaps(); // Enable sourcemaps
// mix.version(); // Enable versioning.
// mix.disableNotifications();
// mix.setPublicPath('path/to/public');
// mix.setResourceRoot('prefix/for/resource/locators');
// mix.autoload({}); <-- Will be passed to Webpack's ProvidePlugin.
// mix.webpackConfig({}); <-- Override webpack.config.js, without editing the file directly.
// mix.babelConfig({}); <-- Merge extra Babel configuration (plugins, etc.) with Mix's default.
// mix.then(function () {}) <-- Will be triggered each time Webpack finishes building.
// mix.extend(name, handler) <-- Extend Mix's API with your own components.
// mix.options({
//   extractVueStyles: false, // Extract .vue component styling to file, rather than inline.
//   globalVueStyles: file, // Variables file to be imported in every component.
//   processCssUrls: true, // Process/optimize relative stylesheet url()'s. Set to false, if you don't want them touched.
//   purifyCss: false, // Remove unused CSS selectors.
//   uglify: {}, // Uglify-specific options. https://webpack.github.io/docs/list-of-plugins.html#uglifyjsplugin
//   postCss: [] // Post-CSS options: https://github.com/postcss/postcss/blob/master/docs/plugins.md
// });
