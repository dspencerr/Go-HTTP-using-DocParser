// Declare app level module which depends on filters, and services
angular.module('BiofireComp', [
    'ngResource',
    'ngRoute',
    'ui.bootstrap',
    'ui.date',
    'angularGrid'
])
  .config(['$routeProvider',
    function ($routeProvider) {
        $routeProvider
          .when('/', {
            templateUrl: 'views/home/home.html',
            controller: 'HomeController'})
          .when('/doc-search', {
            templateUrl: 'views/docsearch/docsearch.html',
            controller: 'DocsearchController'})
          .otherwise({redirectTo: '/'});
  }]);
