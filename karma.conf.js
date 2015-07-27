// Karma configuration
// Generated on Thu Jul 23 2015 08:54:08 GMT-0600 (Mountain Daylight Time)

module.exports = function(config) {
  config.set({

    basePath: 'public/',

    files: [
        'lib/jquery/dist/jquery.js',
        'lib/jquery-ui/jquery-ui.js',
        'lib/lodash/dist/lodash.js',
        'lib/angular/angular.js',
        'lib/angular-resource/angular-resource.js',
        'lib/angular-route/angular-route.js',
        'lib/angular-bootstrap/ui-bootstrap-tpls.js',
        'lib/angular-route/angular-route.js',
        'lib/angular-ui-date/src/date.js',
        'lib/angular-mocks/angular-mocks.js',
        'js/app.js',
        'js/**/*.js',
        'js/tests/**/*.spec.js'
    ],

    autoWatch : true,

    frameworks: ['jasmine'],

    exclude: [

    ],

    reporters: ['progress'],

    port: 9876,

    colors: true,

    logLevel: config.LOG_INFO,

    browsers: ['Chrome'],

    captureTimeout: 60000,

    singleRun: false
  });
};
