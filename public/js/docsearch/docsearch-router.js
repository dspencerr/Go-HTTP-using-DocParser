'use strict';

angular.module('BiofireComp')
  .config(['$routeProvider', function ($routeProvider) {
    $routeProvider
      .when('/docsearches', {
        templateUrl: 'views/docsearch/docsearches.html',
        controller: 'DocsearchController',
        resolve:{
          resolvedDocsearch: ['Docsearch', function (Docsearch) {
            return Docsearch.query();
          }]
        }
      })
    }]);
